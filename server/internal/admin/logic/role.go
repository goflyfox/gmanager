package logic

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/admin/consts"
	dao2 "gmanager/internal/admin/dao"
	"gmanager/internal/admin/model/do"
	"gmanager/internal/admin/model/entity"
	input2 "gmanager/internal/admin/model/input"
	"gmanager/internal/library/gftoken"
)

// Role 角色服务
var Role = new(role)

type role struct{}

// List 获取角色列表
func (s *role) List(ctx context.Context, in *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	if in == nil {
		return
	}
	m := dao2.Role.Ctx(ctx)
	columns := dao2.Role.Columns()
	res = &v1.RoleListRes{}

	// where条件
	if in.Keywords != "" {
		m = m.Where(m.Builder().WhereLike(columns.Name, "%"+in.Keywords+"%").
			WhereOrLike(columns.Code, "%"+in.Keywords+"%"))
	}
	if in.Name != "" {
		m = m.Where(columns.Name, in.Name)
	}
	if in.Enable > 0 {
		m = m.Where(columns.Enable, in.Enable)
	}

	res.Total, err = m.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据数量失败！")
		return
	}
	res.CurrentPage = in.PageNum
	if res.Total == 0 {
		return
	}

	if in.NeedOrderBy() {
		m = m.Order(in.OrderBy)
	} else {
		m = m.Order("sort asc,id desc")
	}
	var pageList []*entity.Role
	if err = m.Page(in.PageNum, in.PageSize).Scan(&pageList); err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
	}
	res.List = pageList

	return
}

func (s *role) Options(ctx context.Context, in *v1.RoleOptionsReq) (res *v1.RoleOptionsRes, err error) {
	if in == nil {
		return
	}

	m := dao2.Role.Ctx(ctx)
	columns := dao2.Role.Columns()
	m = m.Where(columns.Enable, consts.EnableYes).Order("sort asc,id desc")
	var roles []*entity.Role
	if err = m.Scan(&roles); err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
	}

	options := make([]*input2.OptionVal, 0, len(roles))
	for _, v := range roles {
		t := &input2.OptionVal{
			Value: v.Id,
			Label: v.Name,
		}
		options = append(options, t)
	}
	res = &options
	return
}

// Get 获取角色详情
func (s *role) Get(ctx context.Context, id int64) (res *v1.RoleGetRes, err error) {
	err = dao2.Role.Ctx(ctx).Where(dao2.Role.Columns().Id, id).Scan(&res)
	return
}

// Save 保存角色
func (s *role) Save(ctx context.Context, in *v1.RoleSaveReq) error {
	var model do.Role
	err := gconv.Struct(in, &model)
	if err != nil {
		return errors.New("数据转换错误")
	}

	m := dao2.Role.Ctx(ctx)
	columns := dao2.Role.Columns()

	roleId := gftoken.GetSessionUser(ctx).Id
	model.UpdateId = roleId
	model.UpdateAt = gtime.Now()
	if in.Id > 0 {
		model.CreateId = roleId
		model.CreateAt = gtime.Now()
		_, err := m.Where(columns.Id, model.Id).Update(model)
		if err != nil {
			return err
		}
		_ = Log.Save(ctx, model, consts.UPDATE)
	} else {
		model.CreateId = roleId
		model.CreateAt = gtime.Now()
		modelId, err := m.InsertAndGetId(model)
		if err != nil {
			return err
		}
		model.Id = modelId
		_ = Log.Save(ctx, model, consts.INSERT)
	}
	return nil
}

// Delete 删除角色
func (s *role) Delete(ctx context.Context, ids []int) error {
	// 检查是否有用户使用该角色
	userRoleCount, err := dao2.UserRole.Ctx(ctx).WhereIn(dao2.UserRole.Columns().RoleId, ids).Count()
	if err != nil {
		return err
	}
	if userRoleCount > 0 {
		return gerror.New("该角色已被用户使用，无法删除")
	}

	// 删除角色菜单关联
	_, err = dao2.RoleMenu.Ctx(ctx).WhereIn(dao2.RoleMenu.Columns().RoleId, ids).Delete()
	if err != nil {
		return err
	}

	// 删除角色
	_, err = dao2.Role.Ctx(ctx).WhereIn(dao2.Role.Columns().Id, ids).Delete()
	if err != nil {
		return err
	}

	// 删除日志
	for _, id := range ids {
		_ = Log.Save(ctx, do.Role{
			Id: gconv.Int64(id),
		}, consts.DELETE)
	}
	return nil
}

func (s *role) MenuIds(ctx context.Context, id int64) (res *v1.RoleMenuIdsRes, err error) {
	values, err := dao2.RoleMenu.Ctx(ctx).Fields(dao2.RoleMenu.Columns().MenuId).Where(dao2.RoleMenu.Columns().RoleId, id).Array()
	if err != nil {
		return
	}
	arr := gconv.SliceInt64(values)
	res = &arr
	return
}

func (s *role) AddMenuIds(ctx context.Context, in *v1.RoleAddMenuIdsReq) error {
	// 删除历史菜单
	_, err := dao2.RoleMenu.Ctx(ctx).Where(dao2.RoleMenu.Columns().RoleId, in.Id).Delete()
	if err != nil {
		return err
	}
	// 插入新菜单
	roleMenuList := g.List{}
	for _, menuId := range in.MenuIds {
		roleMenuList = append(roleMenuList,
			g.Map{dao2.RoleMenu.Columns().RoleId: in.Id,
				dao2.RoleMenu.Columns().MenuId: menuId,
			})
	}
	_, err = dao2.RoleMenu.Ctx(ctx).Insert(roleMenuList)

	model := do.Role{Id: in.Id}
	userId := gftoken.GetSessionUser(ctx).Id
	model.UpdateId = userId
	model.UpdateAt = gtime.Now()
	_ = Log.SaveLog(ctx, &input2.LogData{
		Model:      model,
		OperType:   consts.INSERT,
		OperRemark: "菜单ID：" + gconv.String(in.MenuIds),
	})
	return err
}
