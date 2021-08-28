package user

import (
	"errors"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"gmanager/app/constants"
	"gmanager/app/model/user"
	"gmanager/app/model/user_role"
	"gmanager/app/service/log"
	"gmanager/library"
	"gmanager/library/base"
)

// 请求参数
type Request struct {
	user.Entity
	UserId int `form:"userId" json:"userId"`
}

// 通过id获取实体
func GetById(id int64) (*user.Entity, error) {
	if id <= 0 {
		glog.Error(" get id error")
		return new(user.Entity), errors.New("参数不合法")
	}

	return user.Model.FindOne(" id = ?", id)
}

// 根据条件获取实体
func GetOne(form *base.BaseForm) (*user.Entity, error) {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["id"] != "" {
		where += " and id = ? "
		params = append(params, gconv.Int(form.Params["id"]))
	}

	return user.Model.FindOne(where, params)
}

// 根据用户名获取实体
func GetByUsername(username string) (*user.Entity, error) {
	if username == "" {
		glog.Error(" getByUsername username error")
		return new(user.Entity), nil
	}

	return user.Model.FindOne("username = ?", username)
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

	r, err1 := user.Model.Delete(" id = ?", id)
	if err1 != nil {
		return 0, err1
	}

	log.SaveLog(entity, constants.DELETE)
	return r.RowsAffected()
}

// 保存实体
func Save(request *Request) (int64, error) {
	entity := (*user.Entity)(nil)
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

		r, err := user.Model.Insert(entity)
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
		r, err := user.Model.OmitEmpty().Where(" id = ?", entity.Id).Update(entity)
		if err != nil {
			return 0, err
		}

		log.SaveLog(entity, constants.UPDATE)
		return r.RowsAffected()
	}
}

// 列表数据查询
func List(form *base.BaseForm) ([]*user.Entity, error) {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	return user.Model.Order(form.OrderBy).FindAll(where, params)
}

// 用户角色列表数据查询
func ListUserRole(form *base.BaseForm) ([]*user_role.Entity, error) {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && gconv.Int(form.Params["userId"]) > 0 {
		where += " and user_id = ? "
		params = append(params, gconv.Int(form.Params["userId"]))
	}

	return user_role.Model.Order(form.OrderBy).FindAll(where, params)
}

// 分页查询
func Page(form *base.BaseForm) ([]user.Entity, error) {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error("page param error", form.Page, form.Rows)
		return []user.Entity{}, nil
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

	num, err := user.Model.As("t").FindCount(where, params)
	form.TotalSize = num
	form.TotalPage = num / form.Rows

	if err != nil {
		glog.Error("page count error", err)
		return []user.Entity{}, err
	}

	// 没有数据直接返回
	if num == 0 {
		form.TotalPage = 0
		form.TotalSize = 0
		return []user.Entity{}, err
	}

	var resData []user.Entity
	dbModel := user.Model.As("t").Fields(user.Model.Columns() +
		",depart.name as departName,su1.real_name as updateName,su2.real_name as createName")
	dbModel = dbModel.LeftJoin("sys_department depart", "t.depart_id = depart.id ")
	dbModel = dbModel.LeftJoin("sys_user su1", " t.update_id = su1.id ")
	dbModel = dbModel.LeftJoin("sys_user su2", " t.update_id = su2.id ")
	err = dbModel.Where(where, params).Order(form.OrderBy).Page(form.Page, form.Rows).M.Structs(&resData)
	if err != nil {
		glog.Error("page list error", err)
		return []user.Entity{}, err
	}

	return resData, nil
}

// 保存用户和角色
func SaveUserRole(userId int, roleIds string) error {
	if userId <= 0 {
		return errors.New("参数错误")
	}

	_, err := user_role.Model.Delete(" user_id = ?", userId)
	if err != nil {
		return err
	}

	if roleIds == "" {
		return nil
	}
	roleIdArray := gstr.Split(roleIds, ",")
	for _, roleId := range roleIdArray {
		userRole := user_role.Entity{UserId: userId, RoleId: gconv.Int(roleId)}
		_, err := userRole.Insert()
		if err != nil {
			return err
		}
	}

	return nil
}
