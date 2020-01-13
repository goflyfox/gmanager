package log

import (
	"errors"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"gmanager/app/model/log"
	"gmanager/module/component/started"
	"gmanager/utils/base"
	"reflect"
)

// 请求参数
type Request struct {
	LogType    int    `orm:"log_type"    json:"log_type"`    // 类型
	OperObject string `orm:"oper_object" json:"oper_object"` // 操作对象
	OperTable  string `orm:"oper_table"  json:"oper_table"`  // 操作表
	OperId     int    `orm:"oper_id"     json:"oper_id"`     // 操作主键
	OperType   string `orm:"oper_type"   json:"oper_type"`   // 操作类型
	OperRemark string `orm:"oper_remark" json:"oper_remark"` // 操作备注
}

func GetById(id int64) (*log.Entity, error) {
	if id <= 0 {
		glog.Error(" get id error")
		return new(log.Entity), errors.New("参数不合法")
	}

	return log.Model.FindOne(" id = ?", id)
}

func GetOne(form *base.BaseForm) (*log.Entity, error) {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["id"] != "" {
		where += " and id = ? "
		params = append(params, gconv.Int(form.Params["id"]))
	}

	return log.Model.FindOne(where, params)
}

func List(form *base.BaseForm) ([]*log.Entity, error) {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	return log.Model.Order(form.OrderBy).FindAll(where, params)
}

func Delete(id int64) (int64, error) {
	if id <= 0 {
		glog.Error("delete id error")
		return 0, errors.New("参数不合法")
	}

	r, err := log.Model.Delete(" id = ?", id)
	if err != nil {
		return 0, err
	}

	return r.RowsAffected()
}

func Update(entity *log.Entity) (int64, error) {
	if entity.Id <= 0 {
		glog.Error("update id error")
		return 0, errors.New("参数不合法")
	}

	r, err := log.Model.Where(" id = ?", entity.Id).Update(entity)
	if err != nil {
		return 0, err
	}

	return r.RowsAffected()
}

func Insert(entity *log.Entity) (int64, error) {
	if entity.Id > 0 {
		glog.Error("insert id error")
		return 0, errors.New("参数不合法")
	}

	r, err := log.Model.Insert(entity)
	if err != nil {
		return 0, err
	}

	return r.RowsAffected()
}

// 分页查询
func Page(form *base.BaseForm) ([]log.Entity, error) {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error("page param error", form.Page, form.Rows)
		return []log.Entity{}, nil
	}

	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["operObject"] != "" {
		where += " and oper_object like ? "
		params = append(params, "%"+form.Params["operObject"]+"%")
	}
	if form.Params != nil && form.Params["operTable"] != "" {
		where += " and oper_table like ? "
		params = append(params, "%"+form.Params["operTable"]+"%")
	}
	if form.Params != nil && gconv.Int(form.Params["logType"]) > 0 {
		where += " and log_type = ? "
		params = append(params, gconv.Int(form.Params["logType"]))
	}
	if form.Params != nil && gconv.Int(form.Params["operType"]) > 0 {
		where += " and oper_type = ? "
		params = append(params, gconv.Int(form.Params["operType"]))
	}

	num, err := log.Model.FindCount(where, params)
	form.TotalSize = num
	form.TotalPage = num / form.Rows

	if err != nil {
		glog.Error("page count error", err)
		return []log.Entity{}, err
	}

	// 没有数据直接返回
	if num == 0 {
		form.TotalPage = 0
		form.TotalSize = 0
		return []log.Entity{}, err
	}

	var resData []log.Entity
	pageNum, pageSize := (form.Page-1)*form.Rows, form.Rows
	dbModel := log.Model.Fields(log.Model.Columns() + ",su1.real_name as updateName,su2.real_name as createName")
	dbModel = dbModel.LeftJoin("sys_user su1", " t.update_id = su1.id ")
	dbModel = dbModel.LeftJoin("sys_user su2", " t.update_id = su2.id ")
	err = dbModel.Where(where, params).Order(form.OrderBy).Page(pageNum, pageSize).M.Structs(&resData)
	if err != nil {
		glog.Error("page list error", err)
		return []log.Entity{}, err
	}

	return resData, nil
}

// 日志保存
func SaveLog(model interface{}, operType string) (int64, error) {
	iModel, ok := model.(base.IModel)
	if !ok {
		glog.Error("transfer error", model)
		return 0, errors.New("transfer error")
	}

	var updateId int
	var updateTime string
	baseModel := reflect.ValueOf(model)
	if kind := baseModel.Kind(); kind == reflect.Ptr {
		updateId = gconv.Int(baseModel.Elem().FieldByName("UpdateId").Interface())
		updateTime = baseModel.Elem().FieldByName("UpdateTime").String()
	} else {
		updateId = gconv.Int(baseModel.FieldByName("UpdateId").Interface())
		updateTime = baseModel.FieldByName("UpdateTime").String()
	}

	logType := log.TypeEdit
	// SELECT table_name,table_comment FROM information_schema.TABLES where table_SCHEMA='gmanager'
	operRemark := ""
	operObject := started.TableInfo[iModel.TableName()]
	if operType == log.LOGIN || operType == log.LOGOUT {
		logType = log.TypeSystem
	} else {
		operRemark = gconv.String(model)
	}

	entity := log.Entity{
		LogType:    logType,
		OperType:   operType,
		OperId:     iModel.PkVal(),
		OperTable:  iModel.TableName(),
		OperObject: operObject,
		OperRemark: operRemark,
		UpdateId:   updateId,
		UpdateTime: updateTime,
		CreateId:   updateId,
		CreateTime: updateTime,
	}
	result, err := log.Insert(entity)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}