package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"gmanager/app/constants"
	"gmanager/app/dao"
	"gmanager/app/model"
	"gmanager/app/service/log"
	"gmanager/library"
	"gmanager/library/base"
)

// 文章管理
var User = userSvc{}

type userSvc struct{}

// 请求参数
type UserReq struct {
	model.User
	UserId int `form:"userId" json:"userId"`
}

// 通过id获取实体
func (s *userSvc) GetById(ctx context.Context, id int64) (*model.User, error) {
	output := &model.User{}
	if id <= 0 {
		glog.Error(" get id error")
		return new(model.User), errors.New("参数不合法")
	}

	err := dao.User.Ctx(ctx).WherePri(id).Scan(&output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// 根据条件获取实体
func (s *userSvc) GetOne(ctx context.Context, form *base.BaseForm) (*model.User, error) {
	output := &model.User{}

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

	err := dao.User.Ctx(ctx).Where(where, params).Scan(&output)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return output, nil
}

// 删除实体
func (s *userSvc) Delete(ctx context.Context, id int64, userId int) error {
	if id <= 0 {
		glog.Error("delete id error")
		return errors.New("参数不合法")
	}

	_, err := dao.User.Ctx(ctx).WherePri(gconv.Int(id)).Delete()
	if err != nil {
		return err
	}

	// 获取删除对象
	entity := model.User{
		Id:         gconv.Int(id),
		UpdateId:   userId,
		UpdateTime: library.GetNow(),
	}
	log.SaveLog(entity, constants.DELETE)
	return nil
}

// 保存实体
func (s *userSvc) Save(ctx context.Context, request *UserReq) (int64, error) {
	entity := (*model.User)(nil)
	err := gconv.Struct(request.User, &entity)
	if err != nil {
		return 0, errors.New("数据错误")
	}

	entity.UpdateId = request.UserId
	entity.UpdateTime = library.GetNow()

	// 判断新增还是修改
	if entity.Id <= 0 {
		entity.CreateId = request.UserId
		entity.CreateTime = library.GetNow()

		r, err := dao.User.Ctx(ctx).Insert(entity)
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
		r, err := dao.User.Ctx(ctx).OmitEmpty().Where(" id = ?", entity.Id).Update(entity)
		if err != nil {
			return 0, err
		}

		log.SaveLog(entity, constants.UPDATE)
		return r.RowsAffected()
	}
}

// 列表数据查询
func (s *userSvc) List(ctx context.Context, form *base.BaseForm) (list []*model.User, err error) {
	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	err = dao.User.Ctx(ctx).Order(form.OrderBy).Where(where, params).Scan(&list)
	return
}

// 分页查询
func (s *userSvc) Page(ctx context.Context, form *base.BaseForm) (list []*model.User, err error) {
	if form.Page <= 0 || form.Rows <= 0 {
		glog.Error("page param error", form.Page, form.Rows)
		err = errors.New("page param error")
		return
	}

	where := " 1 = 1 "
	var params []interface{}
	if form.Params != nil && form.Params["name"] != "" {
		where += " and t.name like ? "
		params = append(params, "%"+form.Params["name"]+"%")
	}

	num, err := dao.User.Ctx(ctx).As("t").FindCount(where, params)
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
		return
	}

	dbModel := dao.User.Ctx(ctx).As("t").Fields(s.Columns() + ",su1.real_name as updateName,su2.real_name as createName")
	dbModel = dbModel.LeftJoin("sys_user su1", " t.update_id = su1.id ")
	dbModel = dbModel.LeftJoin("sys_user su2", " t.update_id = su2.id ")

	err = dbModel.Order(form.OrderBy).Where(where, params).Page(form.Page, form.Rows).Scan(&list)
	return
}

func (s *userSvc) Columns() string {
	sqlColumns := "t.id,t.parent_id as parentId,t.name,t.code,t.sort,t.linkman,t.linkman_no as linkmanNo,t.remark,t.enable,t.update_time as updateTime,t.update_id as updateId,t.create_time as createTime,t.create_id as createId"
	return sqlColumns
}
