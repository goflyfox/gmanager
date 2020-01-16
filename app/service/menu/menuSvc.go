package menu

import (
	"errors"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"gmanager/app/model/menu"
	"gmanager/utils/base"
)

// 请求参数
type Request struct {
	menu.Entity
}

// 通过id获取实体
func GetById(id int64) (*menu.Entity, error) {
	if id <= 0 {
		glog.Error(" get id error")
		return new(menu.Entity), errors.New("参数不合法")
	}

	return menu.Model.FindOne(" id = ?", id)
}

// 根据条件获取实体
func GetOne(form *base.BaseForm) (*menu.Entity, error) {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["id"] != "" {
		where += " and id = ? "
		params = append(params, gconv.Int(form.Params["id"]))
	}
	if form.Params != nil && form.Params["parentId"] != "" {
		where += " and parent_id = ? "
		params = append(params, gconv.Int(form.Params["parentId"]))
	}

	return menu.Model.FindOne(where, params)
}

// 删除实体
func Delete(id int64) (int64, error) {
	if id <= 0 {
		glog.Error("delete id error")
		return 0, errors.New("参数不合法")
	}

	r, err := menu.Model.Delete(" id = ?", id)
	if err != nil {
		return 0, err
	}

	return r.RowsAffected()
}

// 更新实体
func Update(request *Request) (int64, error) {
	entity := (*menu.Entity)(nil)
	err := gconv.StructDeep(request.Entity, &entity)
	if err != nil {
		return 0, errors.New("数据错误")
	}

	if entity.Id <= 0 {
		glog.Error("update id error")
		return 0, errors.New("参数不合法")
	}

	r, err := menu.Model.Where(" id = ?", entity.Id).Update(entity)
	if err != nil {
		return 0, err
	}

	return r.RowsAffected()
}

// 插入实体
func Insert(request *Request) (int64, error) {
	entity := (*menu.Entity)(nil)
	err := gconv.StructDeep(request.Entity, &entity)
	if err != nil {
		return 0, errors.New("数据错误")
	}

	if entity.Id > 0 {
		glog.Error("insert id error")
		return 0, errors.New("参数不合法")
	}

	r, err := menu.Model.Insert(entity)
	if err != nil {
		return 0, err
	}

	return r.RowsAffected()
}

func ListUser(userId int, userType int) ([]*menu.Entity, error) {
	return menu.Model.Fields(menu.Model.Columns()).LeftJoin(
		"sys_role_menu rm", "rm.menu_id = t.id ").LeftJoin(
		"sys_user_role ur", "ur.role_id = rm.role_id ").Where(
		"ur.user_id = ? ", userId).FindAll()
}

// 列表数据查询
func List(form *base.BaseForm) ([]*menu.Entity, error) {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}
	if form.Params != nil && form.Params["level"] != "" {
		where += " and level in (?) "
		params = append(params, gstr.Split(form.Params["level"], ","))
	}
	if gstr.Trim(form.OrderBy) == "" {
		form.OrderBy = " sort,id desc"
	}

	return menu.Model.Order(form.OrderBy).FindAll(where, params)
}

// 分页查询
func Page(form *base.BaseForm) ([]menu.Entity, error) {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error("page param error", form.Page, form.Rows)
		return []menu.Entity{}, nil
	}

	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	num, err := menu.Model.As("t").FindCount(where, params)
	form.TotalSize = num
	form.TotalPage = num / form.Rows

	if err != nil {
		glog.Error("page count error", err)
		return []menu.Entity{}, err
	}

	// 没有数据直接返回
	if num == 0 {
		form.TotalPage = 0
		form.TotalSize = 0
		return []menu.Entity{}, err
	}

	var resData []menu.Entity
	dbModel := menu.Model.As("t").Fields(menu.Model.Columns() + ",su1.real_name as updateName,su2.real_name as createName")
	dbModel = dbModel.LeftJoin("sys_user su1", " t.update_id = su1.id ")
	dbModel = dbModel.LeftJoin("sys_user su2", " t.update_id = su2.id ")
	err = dbModel.Where(where, params).Order(form.OrderBy).Page(form.Page, form.Rows).M.Structs(&resData)
	if err != nil {
		glog.Error("page list error", err)
		return []menu.Entity{}, err
	}

	return resData, nil
}
