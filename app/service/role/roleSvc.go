package role

import (
	"errors"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"gmanager/app/constants"
	"gmanager/app/dao"
	"gmanager/app/model"
	"gmanager/app/model/role"
	"gmanager/app/service/log"
	"gmanager/library"
	"gmanager/library/base"
)

// 请求参数
type Request struct {
	*model.Role
	UserId int `form:"userId" json:"userId"`
}

// 通过id获取实体
func GetById(id int64) (role *model.Role, err error) {
	if id <= 0 {
		glog.Error(" get id error")
		err = errors.New("参数不合法")
		return
	}
	err = dao.Role.Scan(role, " id = ?", id)
	return
}

// 根据条件获取实体
func GetOne(form *base.BaseForm) (role *model.Role, err error) {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["id"] != "" {
		where += " and id = ? "
		params = append(params, gconv.Int(form.Params["id"]))
	}

	err = dao.Role.Scan(role, where)
	return
}

// 删除实体
func Delete(id int64, userId int) error {
	if id <= 0 {
		glog.Error("delete id error")
		return errors.New("参数不合法")
	}

	// 获取删除对象
	entity, err := GetById(id)
	if err != nil {
		return err
	}
	entity.UpdateId = userId
	entity.UpdateTime = library.GetNow()

	_, err1 := dao.Role.Delete(" id = ?", id)
	log.SaveLog(entity, constants.DELETE)
	return err1
}

// 保存实体
func Save(request *Request) (int64, error) {
	model := (*model.Role)(nil)
	err := gconv.Struct(request, &model)
	if err != nil {
		return 0, errors.New("数据错误")
	}

	model.UpdateId = request.UserId
	model.UpdateTime = library.GetNow()

	// 判断新增还是修改
	if model.Id <= 0 {
		model.CreateId = request.UserId
		model.CreateTime = library.GetNow()

		r, err := dao.Role.Insert(model)
		if err != nil {
			return 0, err
		}
		// 回写主键
		lastId, err := r.LastInsertId()
		if err != nil {
			return 0, err
		}
		model.Id = gconv.Int(lastId)
		request.Id = gconv.Int(lastId)

		log.SaveLog(model, constants.INSERT)
		return r.RowsAffected()
	} else {
		r, err := dao.Role.OmitEmpty().Where(" id = ?", model.Id).Update(model)
		if err != nil {
			return 0, err
		}

		log.SaveLog(model, constants.UPDATE)
		return r.RowsAffected()
	}
}

// 列表数据查询
func List(form *base.BaseForm) (list []*model.Role, err error) {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	err = dao.Role.Order(form.OrderBy).Scan(list, where, params)
	return
}

// 角色菜单列表数据查询
func ListRoleMenu(form *base.BaseForm) (list []*model.RoleMenu, err error) {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && gconv.Int(form.Params["roleId"]) > 0 {
		where += " and role_id = ? "
		params = append(params, gconv.Int(form.Params["roleId"]))
	}
	err = dao.RoleMenu.Order(form.OrderBy).Scan(list, where, params)
	return
}

// 分页查询
func Page(form *base.BaseForm) (list []*model.Role, err error) {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error("page param error", form.Page, form.Rows)
		list = []*model.Role{}
		return
	}

	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	num, err := dao.Role.As("t").FindCount(where, params)
	form.TotalSize = num
	form.TotalPage = num / form.Rows

	if err != nil {
		glog.Error("page count error", err)
		return
	}

	// 没有数据直接返回
	if num == 0 {
		form.TotalPage = 0
		form.TotalSize = 0
		list = []*model.Role{}
		return
	}

	dbModel := dao.Role.As("t").Fields(role.Model.Columns() + ",su1.real_name as updateName,su2.real_name as createName")
	dbModel = dbModel.LeftJoin("sys_user su1", " t.update_id = su1.id ")
	dbModel = dbModel.LeftJoin("sys_user su2", " t.update_id = su2.id ")
	err = dbModel.Where(where, params).Order(form.OrderBy).Page(form.Page, form.Rows).Scan(list, where, params)
	return
}

// 保存角色和菜单
func SaveRoleMenu(roleId int, menuIds string) error {
	if roleId <= 0 {
		return errors.New("参数错误")
	}

	_, err := dao.RoleMenu.Delete(" role_id = ?", roleId)
	if err != nil {
		return err
	}

	if menuIds == "" {
		return nil
	}
	menuIdArray := gstr.Split(menuIds, ",")
	for _, menuId := range menuIdArray {
		roleMenu := model.RoleMenu{RoleId: roleId, MenuId: gconv.Int(menuId)}
		_, err = dao.RoleMenu.Insert(roleMenu)
		if err != nil {
			return err
		}
	}

	return nil
}
