package role

import (
	"errors"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"gmanager/app/constants"
	"gmanager/app/model/role"
	"gmanager/app/model/role_menu"
	"gmanager/app/service/log"
	"gmanager/library"
	"gmanager/library/base"
)

// 请求参数
type Request struct {
	role.Entity
	UserId int `form:"userId" json:"userId"`
}

// 通过id获取实体
func GetById(id int64) (*role.Entity, error) {
	if id <= 0 {
		glog.Error(" get id error")
		return new(role.Entity), errors.New("参数不合法")
	}

	return role.Model.FindOne(" id = ?", id)
}

// 根据条件获取实体
func GetOne(form *base.BaseForm) (*role.Entity, error) {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["id"] != "" {
		where += " and id = ? "
		params = append(params, gconv.Int(form.Params["id"]))
	}

	return role.Model.FindOne(where, params)
}

// 删除实体
func Delete(id int64, userId int) (int64, error) {
	if id <= 0 {
		glog.Error("delete id error")
		return 0, errors.New("参数不合法")
	}

	// 获取删除对象
	entity, err := GetById(id)
	if err != nil {
		return 0, err
	}
	entity.UpdateId = userId
	entity.UpdateTime = library.GetNow()

	r, err1 := role.Model.Delete(" id = ?", id)
	if err1 != nil {
		return 0, err1
	}

	log.SaveLog(entity, constants.DELETE)
	return r.RowsAffected()
}

// 保存实体
func Save(request *Request) (int64, error) {
	entity := (*role.Entity)(nil)
	err := gconv.Struct(request.Entity, &entity)
	if err != nil {
		return 0, errors.New("数据错误")
	}

	entity.UpdateId = request.UserId
	entity.UpdateTime = library.GetNow()

	// 判断新增还是修改
	if entity.Id <= 0 {
		entity.CreateId = request.UserId
		entity.CreateTime = library.GetNow()

		r, err := role.Model.Insert(entity)
		if err != nil {
			return 0, err
		}
		// 回写主键
		lastId, err := r.LastInsertId()
		if err != nil {
			return 0, err
		}
		entity.Id = gconv.Int(lastId)
		request.Id = gconv.Int(lastId)

		log.SaveLog(entity, constants.INSERT)
		return r.RowsAffected()
	} else {
		r, err := role.Model.OmitEmpty().Where(" id = ?", entity.Id).Update(entity)
		if err != nil {
			return 0, err
		}

		log.SaveLog(entity, constants.UPDATE)
		return r.RowsAffected()
	}
}

// 列表数据查询
func List(form *base.BaseForm) ([]*role.Entity, error) {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	return role.Model.Order(form.OrderBy).FindAll(where, params)
}

// 角色菜单列表数据查询
func ListRoleMenu(form *base.BaseForm) ([]*role_menu.Entity, error) {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && gconv.Int(form.Params["roleId"]) > 0 {
		where += " and role_id = ? "
		params = append(params, gconv.Int(form.Params["roleId"]))
	}

	return role_menu.Model.Order(form.OrderBy).FindAll(where, params)
}

// 分页查询
func Page(form *base.BaseForm) ([]*role.Entity, error) {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error("page param error", form.Page, form.Rows)
		return []*role.Entity{}, nil
	}

	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	num, err := role.Model.As("t").FindCount(where, params)
	form.TotalSize = num
	form.TotalPage = num / form.Rows

	if err != nil {
		glog.Error("page count error", err)
		return []*role.Entity{}, err
	}

	// 没有数据直接返回
	if num == 0 {
		form.TotalPage = 0
		form.TotalSize = 0
		return []*role.Entity{}, err
	}

	dbModel := role.Model.As("t").Fields(role.Model.Columns() + ",su1.real_name as updateName,su2.real_name as createName")
	dbModel = dbModel.LeftJoin("sys_user su1", " t.update_id = su1.id ")
	dbModel = dbModel.LeftJoin("sys_user su2", " t.update_id = su2.id ")
	return dbModel.Where(where, params).Order(form.OrderBy).Page(form.Page, form.Rows).FindAll(where, params)
}

// 保存角色和菜单
func SaveRoleMenu(roleId int, menuIds string) error {
	if roleId <= 0 {
		return errors.New("参数错误")
	}

	_, err := role_menu.Model.Delete(" role_id = ?", roleId)
	if err != nil {
		return err
	}

	if menuIds == "" {
		return nil
	}
	menuIdArray := gstr.Split(menuIds, ",")
	for _, menuId := range menuIdArray {
		roleMenu := role_menu.Entity{RoleId: roleId, MenuId: gconv.Int(menuId)}
		_, err := roleMenu.Insert()
		if err != nil {
			return err
		}
	}

	return nil
}
