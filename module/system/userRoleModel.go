package system

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gconv"
	"gmanager/utils/base"
)

type SysUserRole struct {
	// columns START
	Id     int `json:"id" gconv:"id,omitempty"`          // 主键
	UserId int `json:"userId" gconv:"user_id,omitempty"` // 用户id
	RoleId int `json:"roleId" gconv:"role_id,omitempty"` // 角色id
	// columns END

	base.BaseModel
}

func (model SysUserRole) Get() SysUserRole {
	if model.Id <= 0 {
		glog.Error(model.TableName() + " get id error")
		return SysUserRole{}
	}

	var resData SysUserRole
	err := model.dbModel("t").Where(" id = ?", model.Id).Fields(model.columns()).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return SysUserRole{}
	}

	return resData
}

func (model SysUserRole) List(form *base.BaseForm) []SysUserRole {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}
	if form.Params != nil && gconv.Int(form.Params["userId"]) > 0 {
		where += " and user_id = ? "
		params = append(params, gconv.Int(form.Params["userId"]))
	}

	var resData []SysUserRole
	err := model.dbModel("t").Fields(model.columns()).Where(where, params).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" list error", err)
		return []SysUserRole{}
	}

	return resData
}

func (model SysUserRole) Page(form *base.BaseForm) []SysUserRole {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error(model.TableName()+" page param error", form.Page, form.Rows)
		return []SysUserRole{}
	}

	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	num, err := model.dbModel("t").Count()
	form.TotalSize = num
	form.TotalPage = num / form.Rows

	// 没有数据直接返回
	if num == 0 {
		form.TotalPage = 0
		form.TotalSize = 0
		return []SysUserRole{}
	}

	var resData []SysUserRole
	pageNum, pageSize := (form.Page-1)*form.Rows, form.Rows
	err = model.dbModel("t").Fields(
		model.columns()).Where(where, params).Limit(pageNum, pageSize).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" page list error", err)
		return []SysUserRole{}
	}

	return resData
}

func (model SysUserRole) DeleteByUserId() int64 {
	if model.UserId <= 0 {
		glog.Error(model.TableName() + " delete by user id error")
		return 0
	}

	r, err := model.dbModel().Where(" user_id = ?", model.UserId).Delete()
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

func (model SysUserRole) Delete() int64 {
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

func (model SysUserRole) Update() int64 {
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

func (model SysUserRole) Insert() int64 {
	r, err := model.dbModel().Data(model).Insert()
	if err != nil {
		glog.Error(model.TableName()+" insert error", err)
		return 0
	}

	res, err2 := r.RowsAffected()
	if err2 != nil {
		glog.Error(model.TableName()+" insert res error", err2)
		return 0
	}

	return res
}

func (model SysUserRole) dbModel(alias ...string) *gdb.Model {
	var tmpAlias string
	if len(alias) > 0 {
		tmpAlias = " " + alias[0]
	}
	tableModel := g.DB().Table(model.TableName() + tmpAlias).Safe()
	return tableModel
}

func (model SysUserRole) PkVal() int {
	return model.Id
}

func (model SysUserRole) TableName() string {
	return "sys_user_role"
}

func (model SysUserRole) columns() string {
	sqlColumns := "t.id,t.user_id as userId,t.role_id as roleId"
	return sqlColumns
}
