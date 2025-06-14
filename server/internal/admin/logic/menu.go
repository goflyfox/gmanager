package logic

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/admin/consts"
	"gmanager/internal/admin/dao"
	"gmanager/internal/admin/model/do"
	"gmanager/internal/admin/model/entity"
	input2 "gmanager/internal/admin/model/input"
	"gmanager/internal/library/gftoken"
)

// Menu 菜单
var Menu = new(menu)

type menu struct{}

// List 获取菜单列表
func (s *menu) List(ctx context.Context, in *v1.MenuListReq) (list *v1.MenuListRes, err error) {
	if in == nil {
		return
	}
	m := dao.Menu.Ctx(ctx)
	columns := dao.Menu.Columns()
	list = &v1.MenuListRes{}

	if in.Name != "" {
		m = m.WhereLike(columns.Name, "%"+in.Name+"%")
	}
	if in.Enable > 0 {
		m = m.Where(columns.Enable, in.Enable)
	}
	// 只查询顶级菜单
	m = m.Where(columns.ParentId, 0)

	list.Total, err = m.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据数量失败！")
		return
	}
	list.CurrentPage = in.PageNum
	if list.Total == 0 {
		return
	}

	if in.OrderBy != "" {
		m = m.Order(in.OrderBy)
	} else {
		m = m.Order("sort asc,id desc")
	}
	var pageList []*entity.Menu
	if err = m.Page(in.PageNum, in.PageSize).Scan(&pageList); err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
	}

	// 获取child列表
	treeData, err := s.allTree(ctx, in)
	if err != nil {
		return
	}
	for _, v := range pageList {
		for _, selectV := range treeData {
			if v.Id == selectV.Id && selectV.ParentId == 0 {
				list.List = append(list.List, selectV)
			}
		}
	}

	return
}

func (s *menu) allTree(ctx context.Context, in *v1.MenuListReq) (list []*input2.MenuTreeRes, err error) {
	if in == nil {
		return
	}
	m := dao.Menu.Ctx(ctx)
	columns := dao.Menu.Columns()
	if in.Enable > 0 {
		m = m.Where(columns.Enable, in.Enable)
	}
	m = m.Order("sort asc,id desc")
	var menus []*entity.Menu
	if err = m.Scan(&menus); err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
	}
	if len(menus) > 0 {
		list = s.GetTree(0, menus)
	}

	return
}

// GetTree 菜单树形菜单
func (s *menu) GetTree(pid int64, list []*entity.Menu) (tree []*input2.MenuTreeRes) {
	tree = make([]*input2.MenuTreeRes, 0, len(list))
	for _, v := range list {
		if v.ParentId == pid {
			t := &input2.MenuTreeRes{
				Menu: v,
			}
			child := s.GetTree(v.Id, list)
			if len(child) > 0 {
				t.Children = child
			}
			tree = append(tree, t)
		}
	}
	return
}

func (s *menu) Options(ctx context.Context, in *v1.MenuOptionsReq) (res *v1.MenuOptionsRes, err error) {
	if in == nil {
		return
	}

	m := dao.Menu.Ctx(ctx)
	columns := dao.Menu.Columns()
	m = m.Where(columns.Enable, consts.EnableYes)
	if in.OnlyParent {
		m = m.WhereIn(columns.Type, []int{consts.MenuTypeMenu, consts.MenuTypeCatalog})
	}
	m = m.Order("sort asc,id desc")
	var menus []*entity.Menu
	if err = m.Scan(&menus); err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
	}
	if len(menus) > 0 {
		tree := s.OptionsTree(0, menus)
		res = &tree
	}
	return
}

// OptionsTree 部门下拉框
func (s *menu) OptionsTree(pid int64, list []*entity.Menu) (tree []*input2.OptionVal) {
	tree = make([]*input2.OptionVal, 0, len(list))
	for _, v := range list {
		if v.ParentId == pid {
			t := &input2.OptionVal{
				Value: v.Id,
				Label: v.Name,
			}
			child := s.OptionsTree(v.Id, list)
			if len(child) > 0 {
				t.Children = child
			}
			tree = append(tree, t)
		}
	}
	return
}

// Get 获取菜单详情
func (s *menu) Get(ctx context.Context, id int) (res *v1.MenuGetRes, err error) {
	err = dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().Id, id).Scan(&res)
	if err != nil {
		return
	}
	if res.Params != "" {
		var paramList []*input2.KeyValue
		err = gconv.Struct(res.Params, &paramList)
		res.ParamList = paramList
	}
	if res.AlwaysShow == 0 {
		res.AlwaysShow = consts.EnableNo
	}
	return
}

// Save 保存菜单
func (s *menu) Save(ctx context.Context, in *v1.MenuSaveReq) error {
	var model do.Menu
	if len(in.ParamList) > 0 {
		in.Params = gconv.String(in.ParamList)
	} else {
		in.Params = "[]"
	}
	err := gconv.Struct(in, &model)
	if err != nil {
		return errors.New("数据转换错误")
	}

	// 菜单名称唯一性校验
	nameCount, err := dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().Name, model.Name).WhereNot(dao.Menu.Columns().Id, in.Id).Count()
	if err != nil {
		return gerror.Wrap(err, "检查菜单名称唯一性失败")
	}
	if nameCount > 0 {
		return gerror.New("菜单名称已存在")
	}

	// 路由名称唯一性校验
	if model.Type != consts.MenuTypeButton && model.RouteName != "" {
		routeNameCount, err := dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().RouteName, model.RouteName).WhereNot(dao.Menu.Columns().Id, in.Id).Count()
		if err != nil {
			return gerror.Wrap(err, "检查路由名称唯一性失败")
		}
		if routeNameCount > 0 {
			return gerror.New("路由名称已存在")
		}
	}
	// 路由路径唯一性校验
	if model.Type != consts.MenuTypeButton && model.RoutePath != "" {
		routePathCount, err := dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().RoutePath, model.RoutePath).WhereNot(dao.Menu.Columns().Id, in.Id).Count()
		if err != nil {
			return gerror.Wrap(err, "检查路由路径唯一性失败")
		}
		if routePathCount > 0 {
			return gerror.New("路由路径已存在")
		}
	}
	// 组件路径唯一性校验
	if model.Type != consts.MenuTypeButton && model.Component != "" {
		componentCount, err := dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().Component, model.Component).WhereNot(dao.Menu.Columns().Id, in.Id).Count()
		if err != nil {
			return gerror.Wrap(err, "检查组件路径唯一性失败")
		}
		if componentCount > 0 {
			return gerror.New("组件路径已存在")
		}
	}

	// 按钮权限唯一性校验
	if model.Type == consts.MenuTypeButton && model.Perm != "" {
		permCount, err := dao.Menu.Ctx(ctx).Where(dao.Menu.Columns().Perm, model.Perm).WhereNot(dao.Menu.Columns().Id, in.Id).Count()
		if err != nil {
			return gerror.Wrap(err, "检查权限标识唯一性失败")
		}
		if permCount > 0 {
			return gerror.New("权限标识已存在")
		}
	}

	m := dao.Menu.Ctx(ctx)
	columns := dao.Menu.Columns()

	userId := gftoken.GetSessionUser(ctx).Id
	model.UpdateId = userId
	model.UpdateAt = gtime.Now()
	if in.Id > 0 {
		model.CreateId = userId
		model.CreateAt = gtime.Now()
		_, err := m.Where(columns.Id, model.Id).Update(model)
		if err != nil {
			return err
		}
		_ = Log.Save(ctx, model, consts.UPDATE)
	} else {
		model.CreateId = userId
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

// Delete 删除菜单
func (s *menu) Delete(ctx context.Context, ids []int) error {
	_, err := dao.Menu.Ctx(ctx).WhereIn(dao.Menu.Columns().Id, ids).Delete()
	if err != nil {
		return err
	}
	// 删除日志
	for _, id := range ids {
		_ = Log.Save(ctx, do.Menu{
			Id: gconv.Int64(id),
		}, consts.DELETE)
	}
	return nil
}
