package system

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gconv"
	"gmanager/utils/base"
)

type SysUser struct {
	// columns START
	Id           int    `json:"id" gconv:"id,omitempty"`                       // 主键
	Uuid         string `json:"uuid" gconv:"uuid,omitempty"`                   // UUID
	Username     string `json:"username" gconv:"username,omitempty"`           // 登录名/11111
	Password     string `json:"password" gconv:"password,omitempty"`           // 密码
	Salt         string `json:"salt" gconv:"salt,omitempty"`                   // 密码盐
	RealName     string `json:"realName" gconv:"real_name,omitempty"`          // 真实姓名
	DepartId     int    `json:"departId" gconv:"depart_id,omitempty"`          // 部门/11111/dict
	UserType     int    `json:"userType" gconv:"user_type,omitempty"`          // 类型//select/1,管理员,2,普通用户,3,前台用户,4,第三方用户,5,API用户
	Status       int    `json:"status" gconv:"status,omitempty"`               // 状态
	Thirdid      string `json:"thirdid" gconv:"thirdid,omitempty"`             // 第三方ID
	Endtime      string `json:"endtime" gconv:"endtime,omitempty"`             // 结束时间
	Email        string `json:"email" gconv:"email,omitempty"`                 // email
	Tel          string `json:"tel" gconv:"tel,omitempty"`                     // 手机号
	Address      string `json:"address" gconv:"address,omitempty"`             // 地址
	TitleUrl     string `json:"titleUrl" gconv:"title_url,omitempty"`          // 头像地址
	Remark       string `json:"remark" gconv:"remark,omitempty"`               // 说明
	Theme        string `json:"theme" gconv:"theme,omitempty"`                 // 主题
	BackSiteId   int    `json:"backSiteId" gconv:"back_site_id,omitempty"`     // 后台选择站点ID
	CreateSiteId int    `json:"createSiteId" gconv:"create_site_id,omitempty"` // 创建站点ID
	ProjectId    int    `json:"projectId" gconv:"project_id,omitempty"`        // 项目ID
	ProjectName  string `json:"projectName" gconv:"project_name,omitempty"`    // 项目名称
	// columns END
	DepartName string `json:"departName" gconv:"departName,omitempty"` // 项目名称

	base.BaseModel
}

func (model SysUser) Get() SysUser {
	if model.Id <= 0 {
		glog.Error(model.TableName() + " get id error")
		return SysUser{}
	}

	var resData SysUser
	err := model.dbModel("t").Where(" id = ?", model.Id).Fields(model.columns()).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get one error", err)
		return SysUser{}
	}

	return resData
}

func (model SysUser) GetByUsername() SysUser {
	if model.Username == "" {
		glog.Error(model.TableName() + " get username error")
		return SysUser{}
	}

	var resData SysUser
	err := model.dbModel("t").Where("username = ?", model.Username).Fields(model.columns()).Struct(&resData)
	if err != nil {
		glog.Error(model.TableName()+" get username one error", err)
		return SysUser{}
	}

	return resData
}

func (model SysUser) List(form *base.BaseForm) []SysUser {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	var resData []SysUser
	err := model.dbModel("t").Fields(
		model.columns()).Where(where, params...).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" list error", err)
		return []SysUser{}
	}

	return resData
}

func (model SysUser) Page(form *base.BaseForm) []SysUser {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error(model.TableName()+" page param error", form.Page, form.Rows)
		return []SysUser{}
	}

	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil {
		if form.Params["username"] != "" {
			where += " and t.username like ? "
			params = append(params, "%"+form.Params["username"]+"%")
		}
		if form.Params["realName"] != "" {
			where += " and t.real_name like ? "
			params = append(params, "%"+form.Params["realName"]+"%")
		}
		if gconv.Int(form.Params["userType"]) > 0 {
			where += " and t.user_type = ? "
			params = append(params, gconv.Int(form.Params["userType"]))
		}
		if gconv.Int(form.Params["departId"]) > 0 {
			where += " and t.depart_id = ? "
			params = append(params, gconv.Int(form.Params["departId"]))
		}
	}

	num, err := model.dbModel("t").Where(where, params...).Count()
	form.TotalSize = num
	form.TotalPage = num / form.Rows

	// 没有数据直接返回
	if num == 0 {
		form.TotalPage = 0
		form.TotalSize = 0
		return []SysUser{}
	}

	var resData []SysUser
	pageNum, pageSize := (form.Page-1)*form.Rows, form.Rows
	dbModel := model.dbModel("t").Fields(model.columns() +
		",depart.name as departName,su1.real_name as updateName,su2.real_name as createName")
	dbModel = dbModel.LeftJoin("sys_department depart", "t.depart_id = depart.id ")
	dbModel = dbModel.LeftJoin("sys_user su1", " t.update_id = su1.id ")
	dbModel = dbModel.LeftJoin("sys_user su2", " t.update_id = su2.id ")
	dbModel = dbModel.Where(where, params...)
	err = dbModel.Limit(pageNum, pageSize).OrderBy(form.OrderBy).Structs(&resData)
	if err != nil {
		glog.Error(model.TableName()+" page list error", err)
		return []SysUser{}
	}

	return resData
}

func (model SysUser) Delete() int64 {
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

func (model SysUser) Update() int64 {
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

func (model *SysUser) Insert() int64 {
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

func (model SysUser) dbModel(alias ...string) *gdb.Model {
	var tmpAlias string
	if len(alias) > 0 {
		tmpAlias = " " + alias[0]
	}
	tableModel := g.DB().Table(model.TableName() + tmpAlias).Safe()
	return tableModel
}

func (model SysUser) PkVal() int {
	return model.Id
}

func (model SysUser) TableName() string {
	return "sys_user"
}

func (model SysUser) columns() string {
	sqlColumns := "t.id,t.uuid,t.username,t.password,t.salt,t.real_name as realName,t.depart_id as departId,t.user_type as userType,t.status,t.thirdid,t.endtime,t.email,t.tel,t.address,t.title_url as titleUrl,t.remark,t.theme,t.back_site_id as backSiteId,t.create_site_id as createSiteId,t.project_id as projectId,t.project_name as projectName,t.enable,t.update_time as updateTime,t.update_id as updateId,t.create_time as createTime,t.create_id as createId"
	return sqlColumns
}
