package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmeta"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/consts"
	"gmanager/internal/dao"
	"gmanager/internal/library/gftoken"
	"gmanager/internal/model/do"
	"gmanager/internal/model/entity"
	"gmanager/internal/model/input"
)

// Log 用户服务
var Log = new(log)

type log struct{}

// List 获取用户列表
func (s *log) List(ctx context.Context, in *v1.LogListReq) (res *v1.LogListRes, err error) {
	if in == nil {
		return
	}
	m := dao.Log.Ctx(ctx)
	columns := dao.Log.Columns()
	res = &v1.LogListRes{}

	// where条件
	if in.Keywords != "" {
		m = m.Where(m.Builder().WhereLike(columns.OperObject, "%"+in.Keywords+"%").
			WhereOrLike(columns.OperTable, "%"+in.Keywords+"%").
			WhereOrLike(columns.OperRemark, "%"+in.Keywords+"%"))
	}
	if in.LogType > 0 {
		m = m.Where(columns.LogType, in.LogType)
	}
	if in.Enable > 0 {
		m = m.Where(columns.Enable, in.Enable)
	}
	if len(in.CreateAt) > 0 && in.CreateAt[0] != "" {
		m = m.WhereBetween(columns.CreateAt, in.CreateAt[0]+" 00:00:00", in.CreateAt[1]+" 23:59:59")
	}

	res.Total, err = m.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据数量失败！")
		return
	}
	res.CurrentPage = in.PageNum
	if res.Total == 0 {
		return
	}

	if in.OrderBy != "" {
		m = m.Order(in.OrderBy)
	} else {
		m = m.Order("id desc")
	}
	var pageList []*entity.Log
	if err = m.Page(in.PageNum, in.PageSize).Scan(&pageList); err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
	}
	res.List = pageList

	return
}

// Get 获取用户详情
func (s *log) Get(ctx context.Context, id int64) (res *v1.LogGetRes, err error) {
	err = dao.Log.Ctx(ctx).Where(dao.Log.Columns().Id, id).Scan(&res)
	return
}

// Save 日志保存
func (s *log) Save(ctx context.Context, model any, operType string) error {
	return s.SaveLog(ctx, &input.LogData{
		Model:    model,
		OperType: operType,
	})
}

// SaveLog 日志保存
func (s *log) SaveLog(ctx context.Context, input *input.LogData) error {
	if input.Model == nil || input.OperType == "" {
		return errors.New("参数错误")
	}

	var (
		logMeta  *entity.Log
		request  = ghttp.RequestFromCtx(ctx)
		clientIp = getClientIp(request)
		operator = input.Operator
	)
	err := gconv.Struct(input.Model, &logMeta)
	if err != nil {
		return err
	}
	session := gftoken.GetSessionUser(ctx)
	if logMeta.UpdateId == 0 {
		logMeta.UpdateId = session.Id
		logMeta.UpdateAt = gtime.Now()
	}
	if session != nil {
		operator = session.Username
	}
	if input.PkVal > 0 {
		logMeta.Id = input.PkVal
	}
	tableName := getTableName(input.Model)
	if input.TableName != "" {
		tableName = input.TableName
	}
	logType := consts.TypeEdit
	if input.OperType == consts.LOGIN || input.OperType == consts.LOGOUT {
		logType = consts.TypeSystem
	}
	logMeta.OperType = input.OperType
	if input.OperRemark != "" {
		logMeta.OperRemark = input.OperRemark
	} else {
		logMeta.OperRemark = input.OperType + ":" + gconv.String(logMeta.Id)
	}
	if input.OperObject != "" {
		logMeta.OperObject = input.OperObject
	} else {
		logMeta.OperObject = gconv.String(input.Model)
	}

	executeTime := gtime.Now().Sub(request.EnterTime).Milliseconds()
	model := do.Log{
		LogType:       logType,
		OperType:      logMeta.OperType,
		OperId:        logMeta.Id,
		OperTable:     tableName,
		OperObject:    logMeta.OperObject,
		OperRemark:    logMeta.OperRemark,
		Url:           request.URL.Path,
		Method:        request.Method,
		Ip:            clientIp,
		UserAgent:     request.Header.Get("User-Agent"),
		ExecutionTime: executeTime,
		Operator:      operator,
		Enable:        consts.EnableYes,
		UpdateId:      logMeta.UpdateId,
		UpdateAt:      logMeta.UpdateAt,
		CreateId:      logMeta.UpdateId,
		CreateAt:      logMeta.UpdateAt,
	}
	_, err = dao.Log.Ctx(ctx).Insert(model)
	return err
}

func getTableName(object any) string {
	if ormTag := gmeta.Get(object, gdb.OrmTagForStruct); !ormTag.IsEmpty() {
		match, _ := gregex.MatchString(
			fmt.Sprintf(`%s\s*:\s*([^,]+)`, gdb.OrmTagForTable),
			ormTag.String(),
		)
		if len(match) > 1 {
			return match[1]
		}
	}
	return ""
}

// getClientIp 获取客户端IP
func getClientIp(r *ghttp.Request) string {
	if r == nil {
		return ""
	}
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.GetClientIp()
	}

	// 兼容部分云厂商CDN，如果存在多个，默认取第一个
	if gstr.Contains(ip, ",") {
		ip = gstr.StrTillEx(ip, ",")
	}

	if gstr.Contains(ip, ", ") {
		ip = gstr.StrTillEx(ip, ", ")
	}
	return ip
}

// Delete 删除用户
func (s *log) Delete(ctx context.Context, ids []int) error {
	_, err := dao.Log.Ctx(ctx).WhereIn(dao.Log.Columns().Id, ids).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (s *log) VisitTrend(ctx context.Context, req *v1.LogVisitTrendReq) (*v1.LogVisitTrendRes, error) {
	var dates []string
	startTime := gtime.NewFromStrFormat(req.StartDate, "Y-m-d")
	endTime := gtime.NewFromStrFormat(req.EndDate, "Y-m-d")
	for !startTime.After(endTime) {
		dates = append(dates, startTime.Format("Y-m-d"))
		startTime = startTime.AddDate(0, 0, 1)
	}

	type result struct {
		PVCount int    `json:"pvCount"`
		IPCount int    `json:"ipCount"`
		Date    string `json:"date"`
	}
	resultToMap := func(res []*result) map[string]*result {
		resMap := make(map[string]*result)
		for _, v := range res {
			resMap[v.Date] = v
		}
		return resMap
	}
	var pvResults []*result
	err := dao.Log.Ctx(ctx).
		Fields("COUNT(1) AS pvCount,COUNT(DISTINCT ip) AS ipCount,DATE_FORMAT(create_at,'%Y-%m-%d') AS date").
		WhereBetween(dao.Log.Columns().CreateAt, req.StartDate+" 00:00:00", req.EndDate+" 23:59:59").
		Group("DATE_FORMAT(create_at, '%Y-%m-%d')").Scan(&pvResults)
	if err != nil {
		return nil, err
	}
	pvMap := resultToMap(pvResults)

	var pvList []int
	var ipList []int
	for _, date := range dates {
		if v, ok := pvMap[date]; ok {
			ipList = append(ipList, v.IPCount)
			pvList = append(pvList, v.PVCount)
		} else {
			ipList = append(ipList, 0)
			pvList = append(pvList, 0)
		}
	}

	return &v1.LogVisitTrendRes{
		Dates:  dates,
		PVList: pvList,
		IPList: ipList,
	}, nil
}
