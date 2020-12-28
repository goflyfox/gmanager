package menu

import (
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"gmanager/app/constants"
	"gmanager/app/model/menu"
	"gmanager/app/service/log"
	"gmanager/library"
	"gmanager/library/base"
)

// 请求参数
type Request struct {
	menu.Entity
	UserId int `form:"userId" json:"userId"`
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

	r, err1 := menu.Model.Delete(" id = ?", id)
	if err1 != nil {
		return 0, err1
	}

	log.SaveLog(entity, constants.DELETE)
	return r.RowsAffected()
}

// 保存实体
func Save(request *Request) (int64, error) {
	entity := (*menu.Entity)(nil)
	err := gconv.Struct(request.Entity, &entity)
	if err != nil {
		return 0, errors.New("数据错误")
	}

	entity.UpdateId = request.UserId
	entity.UpdateTime = library.GetNow()

	// 根目录级别为1，其他为父节点 + 1
	parentId := entity.ParentId
	if parentId == 0 {
		entity.Level = 1
	} else {
		form := base.NewForm(g.Map{"id": parentId})
		parentModel, err := GetOne(&form)
		if err != nil {
			return 0, err
		}
		entity.Level = parentModel.Level + 1
	}

	// 判断新增还是修改
	if entity.Id <= 0 {
		entity.CreateId = request.UserId
		entity.CreateTime = library.GetNow()

		r, err := menu.Model.Insert(entity)
		if err != nil {
			return 0, err
		}
		// 回写主键
		lastId, err := r.LastInsertId()
		if err != nil {
			return 0, err
		}
		entity.Id = gconv.Int(lastId)

		log.SaveLog(entity, constants.INSERT)
		return r.RowsAffected()
	} else {
		r, err := menu.Model.OmitEmpty().Where(" id = ?", entity.Id).Update(entity)
		if err != nil {
			return 0, err
		}

		log.SaveLog(entity, constants.UPDATE)
		return r.RowsAffected()
	}
}

func ListUser(userId int, userType int) ([]*menu.Entity, error) {
	if userType == constants.UserTypeAdmin {
		return List(&base.BaseForm{})
	}

	return menu.Model.As("t").Fields(menu.Model.Columns()).LeftJoin(
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
