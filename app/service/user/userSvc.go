package user

import (
	"errors"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"gmanager/app/constants"
	"gmanager/app/dao"
	"gmanager/app/model"
	"gmanager/app/service/log"
	"gmanager/library"
	"gmanager/library/base"
)

// 请求参数
type Request struct {
	model.User
	UserId int `form:"userId" json:"userId"`
}

// 通过id获取实体
func GetById(id int64) (user *model.User, err error) {
	if id <= 0 {
		glog.Error(" get id error")
		err = errors.New("参数不合法")
		return
	}

	err = dao.User.Scan(&user, " id = ?", id)
	return
}

// 根据条件获取实体
func GetOne(form *base.BaseForm) (user *model.User, err error) {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["id"] != "" {
		where += " and id = ? "
		params = append(params, gconv.Int(form.Params["id"]))
	}

	err = dao.User.Scan(&user, where, params)
	return
}

// 根据用户名获取实体
func GetByUsername(username string) (user *model.User, err error) {
	if username == "" {
		glog.Error(" getByUsername username error")
		user = new(model.User)
		return
	}

	err = dao.User.Scan(&user, "username = ?", username)
	return
}

// 删除实体
func Delete(id int64, userId int) (int64, error) {
	if id <= 0 {
		glog.Error("delete id error")
		return 0, errors.New("参数不合法")
	}

	// 获取删除对象
	model, err := GetById(id)
	if err != nil {
		return 0, err
	}
	model.UpdateId = userId
	model.UpdateTime = library.GetNow()

	r, err1 := dao.User.Delete(" id = ?", id)
	if err1 != nil {
		return 0, err1
	}

	log.SaveLog(model, constants.DELETE)
	return r.RowsAffected()
}

// 保存实体
func Save(request *Request) (int64, error) {
	model := (*model.User)(nil)
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

		r, err := dao.User.Insert(model)
		if err != nil {
			return 0, err
		}
		// 回写主键
		lastId, err := r.LastInsertId()
		if err != nil {
			return 0, err
		}
		model.Id = gconv.Int(lastId)

		log.SaveLog(model, constants.INSERT)
		return r.RowsAffected()
	} else {
		r, err := dao.User.OmitEmpty().Where(" id = ?", model.Id).Update(model)
		if err != nil {
			return 0, err
		}

		log.SaveLog(model, constants.UPDATE)
		return r.RowsAffected()
	}
}

// 列表数据查询
func List(form *base.BaseForm) (list []*model.User, err error) {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	err = dao.User.Order(form.OrderBy).Scan(&list, where, params)
	return
}

// 用户角色列表数据查询
func ListUserRole(form *base.BaseForm) (list []*model.UserRole, err error) {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && gconv.Int(form.Params["userId"]) > 0 {
		where += " and user_id = ? "
		params = append(params, gconv.Int(form.Params["userId"]))
	}

	err = dao.UserRole.Order(form.OrderBy).Scan(&list, where, params)
	return
}

// 分页查询
func Page(form *base.BaseForm) (list []*model.User, err error) {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error("page param error", form.Page, form.Rows)
		list = []*model.User{}
		return
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

	num, err := dao.User.As("t").FindCount(where, params)
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
		list = []*model.User{}
		return
	}

	dbModel := dao.User.As("t").Fields((model.User{}).Columns() +
		",depart.name as departName,su1.real_name as updateName,su2.real_name as createName")
	dbModel = dbModel.LeftJoin("sys_department depart", "t.depart_id = depart.id ")
	dbModel = dbModel.LeftJoin("sys_user su1", " t.update_id = su1.id ")
	dbModel = dbModel.LeftJoin("sys_user su2", " t.update_id = su2.id ")
	err = dbModel.Where(where, params).Order(form.OrderBy).Page(form.Page, form.Rows).Scan(&list)
	return
}

// 保存用户和角色
func SaveUserRole(userId int, roleIds string) error {
	if userId <= 0 {
		return errors.New("参数错误")
	}

	_, err := dao.UserRole.Delete(" user_id = ?", userId)
	if err != nil {
		return err
	}

	if roleIds == "" {
		return nil
	}
	roleIdArray := gstr.Split(roleIds, ",")
	for _, roleId := range roleIdArray {
		userRole := model.UserRole{UserId: userId, RoleId: gconv.Int(roleId)}
		_, err := dao.UserRole.Insert(userRole)
		if err != nil {
			return err
		}
	}

	return nil
}
