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

// Dept 部门服务
var Dept = new(dept)

type dept struct{}

// List 获取部门列表
func (s *dept) List(ctx context.Context, in *v1.DeptListReq) (list *v1.DeptListRes, err error) {
	if in == nil {
		return
	}
	m := dao.Dept.Ctx(ctx)
	columns := dao.Dept.Columns()
	list = &v1.DeptListRes{}

	if in.Keywords != "" {
		m = m.WhereLike(columns.Name, "%"+in.Keywords+"%")
	}
	if in.Code != "" {
		m = m.WhereLike(columns.Code, "%"+in.Code+"%")
	}
	if in.Enable > 0 {
		m = m.Where(columns.Enable, in.Enable)
	}
	// 只查询顶级部门
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

	if in.NeedOrderBy() {
		m = m.Order(in.OrderBy)
	} else {
		m = m.Order("sort asc,id desc")
	}
	var pageList []*entity.Dept
	if err = m.Page(in.PageNum, in.PageSize).Scan(&pageList); err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
	}

	// 获取child列表
	selectList, err := s.SelectALlList(ctx, in)
	if err != nil {
		return
	}
	for _, v := range pageList {
		for _, selectV := range selectList {
			if v.Id == selectV.Id && selectV.ParentId == 0 {
				list.List = append(list.List, selectV)
			}
		}
	}

	return
}

// SelectALlList 获取所有部门
func (s *dept) SelectALlList(ctx context.Context, in *v1.DeptListReq) (list []*input2.DeptTreeRes, err error) {
	if in == nil {
		return
	}
	m := dao.Dept.Ctx(ctx)
	columns := dao.Dept.Columns()
	list = make([]*input2.DeptTreeRes, 0)
	if in.Enable > 0 {
		m = m.Where(columns.Enable, in.Enable)
	}
	m = m.Order("sort asc,id desc")
	var depts []*entity.Dept
	if err = m.Scan(&depts); err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
	}
	if len(depts) > 0 {
		list = s.GetTree(0, depts)
	}

	return
}

// DeptMap 获取所有部门Map
func (s *dept) DeptMap(ctx context.Context) (deptMap map[int64]*entity.Dept, err error) {
	m := dao.Dept.Ctx(ctx)
	columns := dao.Dept.Columns()
	m = m.Where(columns.Enable, consts.EnableYes)
	var depts []*entity.Dept
	if err = m.Scan(&depts); err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	deptMap = make(map[int64]*entity.Dept, len(depts))
	for _, e := range depts {
		deptMap[e.Id] = e
	}
	return
}

// DeptNameMap 获取所有部门Map
func (s *dept) DeptNameMap(ctx context.Context) (deptMap map[string]*entity.Dept, err error) {
	m := dao.Dept.Ctx(ctx)
	columns := dao.Dept.Columns()
	m = m.Where(columns.Enable, consts.EnableYes)
	var depts []*entity.Dept
	if err = m.Scan(&depts); err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	deptMap = make(map[string]*entity.Dept, len(depts))
	for _, e := range depts {
		deptMap[e.Name] = e
	}
	return
}

// GetTree 部门树形菜单
func (s *dept) GetTree(pid int64, list []*entity.Dept) (tree []*input2.DeptTreeRes) {
	tree = make([]*input2.DeptTreeRes, 0, len(list))
	for _, v := range list {
		if v.ParentId == pid {
			t := &input2.DeptTreeRes{
				Dept: v,
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

func (s *dept) Options(ctx context.Context, in *v1.DeptOptionsReq) (res *v1.DeptOptionsRes, err error) {
	if in == nil {
		return
	}

	m := dao.Dept.Ctx(ctx)
	columns := dao.Dept.Columns()
	m = m.Where(columns.Enable, consts.EnableYes)
	var depts []*entity.Dept
	if err = m.Scan(&depts); err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
	}
	if len(depts) > 0 {
		tree := s.OptionsTree(0, depts)
		res = &tree
	}
	return
}

// OptionsTree 部门下拉框
func (s *dept) OptionsTree(pid int64, list []*entity.Dept) (tree []*input2.OptionVal) {
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

// Get 获取部门详情
func (s *dept) Get(ctx context.Context, id int64) (model *v1.DeptGetRes, err error) {
	err = dao.Dept.Ctx(ctx).Where(dao.Dept.Columns().Id, id).Scan(&model)
	return
}

// Save 保存部门
func (s *dept) Save(ctx context.Context, in *v1.DeptSaveReq) error {
	var model do.Dept
	err := gconv.Struct(in, &model)
	if err != nil {
		return errors.New("数据转换错误")
	}

	m := dao.Dept.Ctx(ctx)
	columns := dao.Dept.Columns()

	// 唯一性校验：部门名称
	nameCount, err := m.Where(columns.Name, model.Name).WhereNot(columns.Id, in.Id).Count()
	if err != nil {
		return gerror.Wrap(err, "检查部门名称唯一性失败")
	}
	if nameCount > 0 {
		return gerror.New("部门名称已存在")
	}

	// 唯一性校验：部门编号
	codeCount, err := m.Where(columns.Code, model.Code).WhereNot(columns.Id, in.Id).Count()
	if err != nil {
		return gerror.Wrap(err, "检查部门编号唯一性失败")
	}
	if codeCount > 0 {
		return gerror.New("部门编号已存在")
	}

	userId := gftoken.GetSessionUser(ctx).Id
	model.UpdateId = userId
	model.UpdateAt = gtime.Now()
	if in.Id > 0 {
		model.CreateId = userId
		model.CreateAt = gtime.Now()
		_, err = m.Where(columns.Id, model.Id).Update(model)
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

// Delete 删除部门
func (s *dept) Delete(ctx context.Context, ids []int) error {
	_, err := dao.Dept.Ctx(ctx).WhereIn(dao.Dept.Columns().Id, ids).Delete()
	if err != nil {
		return err
	}
	// 删除日志
	for _, id := range ids {
		_ = Log.Save(ctx, do.Dept{
			Id: gconv.Int64(id),
		}, consts.DELETE)
	}
	return nil
}
