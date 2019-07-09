package system

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gconv"
	"gmanager/utils/base"
)

type SysConfig struct {
	// columns START
	Id           int    `json:"id" gconv:"id,omitempty"`                      // 主键
	Name         string `json:"name" gconv:"name,omitempty"`                  // 名称
	Key          string `json:"key" gconv:"key,omitempty"`                    // 键
	Value        string `json:"value" gconv:"value,omitempty"`                // 值
	Code         string `json:"code" gconv:"code,omitempty"`                  // 编码
	DataType     int    `json:"dataType" gconv:"data_type,omitempty"`         // 数据类型//radio/1,KV,2,字典,3,数组
	ParentId     int    `json:"parentId" gconv:"parent_id,omitempty"`         // 类型
	ParentKey    string `json:"parentKey" gconv:"parent_key,omitempty"`       //
	Sort         int    `json:"sort" gconv:"sort,omitempty"`                  // 排序号
	ProjectId    int    `json:"projectId" gconv:"project_id,omitempty"`       // 项目ID
	CopyStatus   string `json:"copyStatus" gconv:"copy_status,omitempty"`     // 拷贝状态 1 拷贝  2  不拷贝
	ChangeStatus string `json:"changeStatus" gconv:"change_status,omitempty"` // 1 不可更改 2 可以更改
	// columns END

	base.BaseModel
}

func (model SysConfig) Get() SysConfig {
	if model.Id <= 0 {
		glog.Error(model.TableName() + " get id error")
		return SysConfig{}
	}

	var resData SysConfig
	err := model.dbModel("t").Where(" id = ?", model.Id).Fields(model.columns()).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return SysConfig{}
	}

	return resData
}

func (model SysConfig) GetOne(form *base.BaseForm) SysConfig {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["id"] != "" {
		where += " and id = ? "
		params = append(params, gconv.Int(form.Params["id"]))
	}

	var resData SysConfig
	err := model.dbModel("t").Where(where, params).Fields(model.columns()).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return SysConfig{}
	}

	return resData
}

func (model SysConfig) List(form *base.BaseForm) []SysConfig {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}
	if form.Params != nil && form.Params["parentId"] != "" {
		where += " and parent_id = ? "
		params = append(params, gconv.Int(form.Params["parentId"]))
	}

	var resData []SysConfig
	err := model.dbModel("t").Fields(
		model.columns()).Where(where, params...).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" list error", err)
		return []SysConfig{}
	}

	return resData
}

func (model SysConfig) Page(form *base.BaseForm) []SysConfig {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error(model.TableName()+" page param error", form.Page, form.Rows)
		return []SysConfig{}
	}

	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil {
		if form.Params["name"] != "" {
			where += " and name like ? "
			params = append(params, "%"+form.Params["name"]+"%")
		}
		if form.Params["key"] != "" {
			where += " and t.key like ? "
			params = append(params, "%"+form.Params["key"]+"%")
		}
		if gconv.Int(form.Params["parentId"]) > 0 {
			where += " and t.parent_id = ? "
			params = append(params, gconv.Int(form.Params["parentId"]))
		}
	}

	num, err := model.dbModel("t").Where(where, params...).Count()
	form.TotalSize = num
	form.TotalPage = num / form.Rows

	// 没有数据直接返回
	if num == 0 {
		form.TotalPage = 0
		form.TotalSize = 0
		return []SysConfig{}
	}

	var resData []SysConfig
	pageNum, pageSize := (form.Page-1)*form.Rows, form.Rows
	dbModel := model.dbModel("t").Fields(model.columns() + ",su1.real_name as updateName,su2.real_name as createName")
	dbModel = dbModel.LeftJoin("sys_user su1", " t.update_id = su1.id ")
	dbModel = dbModel.LeftJoin("sys_user su2", " t.update_id = su2.id ")
	err = dbModel.Where(where, params...).Limit(pageNum, pageSize).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" page list error", err)
		return []SysConfig{}
	}

	return resData
}

func (model SysConfig) Delete() int64 {
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

	LogSave(model, DELETE)
	return res
}

func (model SysConfig) Update() int64 {
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
	LogSave(model, UPDATE)

	return res
}

func (model *SysConfig) Insert() int64 {
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

	LogSave(model, INSERT)
	return res
}

func (model SysConfig) dbModel(alias ...string) *gdb.Model {
	var tmpAlias string
	if len(alias) > 0 {
		tmpAlias = " " + alias[0]
	}
	tableModel := g.DB().Table(model.TableName() + tmpAlias).Safe()
	return tableModel
}

func (model SysConfig) PkVal() int {
	return model.Id
}

func (model SysConfig) TableName() string {
	return "sys_config"
}

func (model SysConfig) columns() string {
	sqlColumns := "t.id,t.name,t.key,t.value,t.code,t.data_type as dataType,t.parent_id as parentId,t.parent_key as parentKey,t.sort,t.project_id as projectId,t.copy_status as copyStatus,t.change_status as changeStatus,t.enable,t.update_time as updateTime,t.update_id as updateId,t.create_time as createTime,t.create_id as createId"
	return sqlColumns
}
