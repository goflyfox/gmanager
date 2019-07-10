package system

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gconv"
	"gmanager/module/component/started"
	"gmanager/utils/base"
	"reflect"
)

type SysLog struct {
	// columns START
	Id         int    `json:"id" gconv:"id,omitempty"`                  // 主键
	LogType    int    `json:"logType" gconv:"log_type,omitempty"`       // 类型
	OperObject string `json:"operObject" gconv:"oper_object,omitempty"` // 操作对象
	OperTable  string `json:"operTable" gconv:"oper_table,omitempty"`   // 操作表
	OperId     int    `json:"operId" gconv:"oper_id,omitempty"`         // 操作主键
	OperType   string `json:"operType" gconv:"oper_type,omitempty"`     // 操作类型
	OperRemark string `json:"operRemark" gconv:"oper_remark,omitempty"` // 操作备注
	// columns END

	base.BaseModel
}

const (
	LOGIN  = "登录"
	LOGOUT = "登出"
	INSERT = "插入"
	UPDATE = "更新"
	DELETE = "删除"
)

func (model SysLog) Get() SysLog {
	if model.Id <= 0 {
		glog.Error(model.TableName() + " get id error")
		return SysLog{}
	}

	var resData SysLog
	err := model.dbModel("t").Where(" id = ?", model.Id).Fields(model.columns()).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return SysLog{}
	}

	return resData
}

func (model SysLog) GetOne(form *base.BaseForm) SysLog {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["id"] != "" {
		where += " and id = ? "
		params = append(params, gconv.Int(form.Params["id"]))
	}

	var resData SysLog
	err := model.dbModel("t").Where(where, params).Fields(model.columns()).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return SysLog{}
	}

	return resData
}

func (model SysLog) List(form *base.BaseForm) []SysLog {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	var resData []SysLog
	err := model.dbModel("t").Fields(
		model.columns()).Where(where, params...).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" list error", err)
		return []SysLog{}
	}

	return resData
}

func (model SysLog) Page(form *base.BaseForm) []SysLog {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error(model.TableName()+" page param error", form.Page, form.Rows)
		return []SysLog{}
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

	num, err := model.dbModel("t").Where(where, params...).Count()
	form.TotalSize = num
	form.TotalPage = num / form.Rows

	// 没有数据直接返回
	if num == 0 {
		form.TotalPage = 0
		form.TotalSize = 0
		return []SysLog{}
	}

	var resData []SysLog
	pageNum, pageSize := (form.Page-1)*form.Rows, form.Rows
	dbModel := model.dbModel("t").Fields(model.columns() + ",su1.real_name as updateName,su2.real_name as createName")
	dbModel = dbModel.LeftJoin("sys_user su1", " t.update_id = su1.id ")
	dbModel = dbModel.LeftJoin("sys_user su2", " t.update_id = su2.id ")
	err = dbModel.Where(where, params...).Limit(pageNum, pageSize).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" page list error", err)
		return []SysLog{}
	}

	return resData
}

func (model SysLog) Delete() int64 {
	if model.Id <= 0 {
		glog.Error(model.TableName() + " delete id error")
		return 0
	}

	r, err := model.dbModel().Where(" id = ?", model.Id).Delete()
	if err != nil {
		glog.Error(model.TableName()+" delete error", err)
		return 0
	}

	res, err2 := r.RowsAffected()
	if err2 != nil {
		glog.Error(model.TableName()+" delete res error", err2)
		return 0
	}

	return res
}

func (model SysLog) Update() int64 {
	r, err := model.dbModel().Data(model).Where(" id = ?", model.Id).Update()
	if err != nil {
		glog.Error(model.TableName()+" update error", err)
		return 0
	}

	res, err2 := r.RowsAffected()
	if err2 != nil {
		glog.Error(model.TableName()+" update res error", err2)
		return 0
	}

	return res
}

func (model *SysLog) Insert() int64 {
	r, err := model.dbModel().Data(model).Insert()
	if err != nil {
		glog.Error(model.TableName()+" insert error", err)
		return 0
	}

	res, err2 := r.RowsAffected()
	if err2 != nil {
		glog.Error(model.TableName()+" insert res error", err2)
		return 0
	} else if res > 0 {
		lastId, err2 := r.LastInsertId()
		if err2 != nil {
			glog.Error(model.TableName()+" LastInsertId res error", err2)
			return 0
		} else {
			model.Id = gconv.Int(lastId)
		}
	}

	return res
}

func (model SysLog) dbModel(alias ...string) *gdb.Model {
	var tmpAlias string
	if len(alias) > 0 {
		tmpAlias = " " + alias[0]
	}
	tableModel := g.DB().Table(model.TableName() + tmpAlias).Safe()
	return tableModel
}

func (model SysLog) PkVal() int {
	return model.Id
}

func (model SysLog) TableName() string {
	return "sys_log"
}

func (model SysLog) columns() string {
	sqlColumns := "t.id,t.log_type as logType,t.oper_object as operObject,t.oper_table as operTable,t.oper_id as operId,t.oper_type as operType,t.oper_remark as operRemark,t.enable,t.update_time as updateTime,t.update_id as updateId,t.create_time as createTime,t.create_id as createId"
	return sqlColumns
}

func LogSave(model interface{}, operType string) int64 {
	iModel, ok := model.(base.IModel)
	if !ok {
		glog.Error("transfer error", model)
		return 0
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

	logType := 2
	// SELECT table_name,table_comment FROM information_schema.TABLES where table_SCHEMA='gmanager'
	operRemark := ""
	operObject := started.TableInfo[iModel.TableName()]
	if operType == LOGIN || operType == LOGOUT {
		logType = 1
	} else {
		operRemark = gconv.String(model)
	}

	log := SysLog{
		LogType:    logType,
		OperType:   operType,
		OperId:     iModel.PkVal(),
		OperTable:  iModel.TableName(),
		OperObject: operObject,
		OperRemark: operRemark,
		BaseModel: base.BaseModel{
			UpdateId:   updateId,
			UpdateTime: updateTime,
			CreateId:   updateId,
			CreateTime: updateTime,
		},
	}
	return log.Insert()
}
