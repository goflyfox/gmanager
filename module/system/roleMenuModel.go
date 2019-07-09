package system

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gconv"
	"gmanager/utils/base"
)

type SysRoleMenu struct {
	// columns START
	Id     int `json:"id" gconv:"id,omitempty"`          // 主键
	RoleId int `json:"roleId" gconv:"role_id,omitempty"` // 角色id
	MenuId int `json:"menuId" gconv:"menu_id,omitempty"` // 菜单id
	// columns END

	base.BaseModel
}

func (model SysRoleMenu) Get() SysRoleMenu {
	if model.Id <= 0 {
		glog.Error(model.TableName() + " get id error")
		return SysRoleMenu{}
	}

	var resData SysRoleMenu
	err := model.dbModel("t").Where(" id = ?", model.Id).Fields(model.columns()).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return SysRoleMenu{}
	}

	return resData
}

func (model SysRoleMenu) List(form *base.BaseForm) []SysRoleMenu {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}
	if form.Params != nil && gconv.Int(form.Params["roleId"]) > 0 {
		where += " and role_id = ? "
		params = append(params, gconv.Int(form.Params["roleId"]))
	}

	var resData []SysRoleMenu
	err := model.dbModel("t").Fields(
		model.columns()).Where(where, params...).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" list error", err)
		return []SysRoleMenu{}
	}

	return resData
}

func (model SysRoleMenu) Page(form *base.BaseForm) []SysRoleMenu {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error(model.TableName()+" page param error", form.Page, form.Rows)
		return []SysRoleMenu{}
	}

	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	num, err := model.dbModel("t").Where(where, params...).Count()
	form.TotalSize = num
	form.TotalPage = num / form.Rows

	// 没有数据直接返回
	if num == 0 {
		form.TotalPage = 0
		form.TotalSize = 0
		return []SysRoleMenu{}
	}

	var resData []SysRoleMenu
	pageNum, pageSize := (form.Page-1)*form.Rows, form.Rows
	err = model.dbModel("t").Fields(
		model.columns()).Where(where, params...).Limit(pageNum, pageSize).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" page list error", err)
		return []SysRoleMenu{}
	}

	return resData
}

func (model SysRoleMenu) DeleteByRoleId() int64 {
	if model.RoleId <= 0 {
		glog.Error(model.TableName() + " delete by role id error")
		return 0
	}

	r, err := model.dbModel().Where(" role_id = ?", model.RoleId).Delete()
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

func (model SysRoleMenu) Delete() int64 {
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

func (model SysRoleMenu) Update() int64 {
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

func (model *SysRoleMenu) Insert() int64 {
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

func (model SysRoleMenu) dbModel(alias ...string) *gdb.Model {
	var tmpAlias string
	if len(alias) > 0 {
		tmpAlias = " " + alias[0]
	}
	tableModel := g.DB().Table(model.TableName() + tmpAlias).Safe()
	return tableModel
}

func (model SysRoleMenu) PkVal() int {
	return model.Id
}

func (model SysRoleMenu) TableName() string {
	return "sys_role_menu"
}

func (model SysRoleMenu) columns() string {
	sqlColumns := "t.id,t.role_id as roleId,t.menu_id as menuId"
	return sqlColumns
}
