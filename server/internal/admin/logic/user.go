package logic

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/guid"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/admin/consts"
	dao2 "gmanager/internal/admin/dao"
	"gmanager/internal/admin/model/do"
	entity2 "gmanager/internal/admin/model/entity"
	input2 "gmanager/internal/admin/model/input"
	"gmanager/internal/library/gftoken"
)

// User 用户服务
var User = new(user)

type user struct{}

// List 获取用户列表
func (s *user) List(ctx context.Context, in *v1.UserListReq) (res *v1.UserListRes, err error) {
	if in == nil {
		return
	}
	m := dao2.User.Ctx(ctx)
	columns := dao2.User.Columns()
	res = &v1.UserListRes{}

	// where条件
	if in.Keywords != "" {
		m = m.Where(m.Builder().WhereLike(columns.UserName, "%"+in.Keywords+"%").
			WhereOrLike(columns.NickName, "%"+in.Keywords+"%").
			WhereOrLike(columns.Mobile, "%"+in.Keywords+"%"))
	}
	if in.DeptId > 0 {
		m = m.Where(columns.DeptId, in.DeptId)
	}
	if in.Status > 0 {
		m = m.Where(columns.Status, in.Status)
	}
	if in.Enable > 0 {
		m = m.Where(columns.Enable, in.Enable)
	}
	if len(in.CreateAt) > 0 && in.CreateAt[0] != "" {
		m = m.WhereBetween(columns.CreateAt, in.CreateAt[0]+" 00:00:00", in.CreateAt[1]+" 23:59:59")
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

	if in.OrderBy != "" {
		m = m.Order(in.OrderBy)
	} else {
		m = m.Order("id desc")
	}
	var pageList []*input2.User
	if err = m.Page(in.PageNum, in.PageSize).Scan(&pageList); err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
	}
	// 设置部门信息
	depts, err := Dept.DeptMap(ctx)
	if err != nil {
		return nil, err
	}
	for _, userInfo := range pageList {
		if deptInfo, ok := depts[userInfo.DeptId]; ok {
			userInfo.DeptName = deptInfo.Name
		}
	}
	res.List = pageList
	return
}

// Get 获取用户详情
func (s *user) Get(ctx context.Context, id int64) (res *v1.UserGetRes, err error) {
	err = dao2.User.Ctx(ctx).Where(dao2.User.Columns().Id, id).Scan(&res)
	if err != nil {
		return
	}
	values, err := dao2.UserRole.Ctx(ctx).Fields(dao2.UserRole.Columns().RoleId).Where(dao2.UserRole.Columns().UserId, id).Array()
	if err != nil {
		return
	}
	res.RoleIds = gconv.SliceInt64(values)
	return
}

// Save 保存用户
func (s *user) Save(ctx context.Context, in *v1.UserSaveReq) error {
	var model do.User
	err := gconv.Struct(in, &model)
	if err != nil {
		return gerror.Wrap(err, "数据转换错误")
	}

	m := dao2.User.Ctx(ctx)
	columns := dao2.User.Columns()

	// 用户名唯一性校验
	nameCount, err := m.Where(columns.UserName, model.UserName).
		WhereNot(columns.Id, in.Id).Count()
	if err != nil {
		return gerror.Wrap(err, "检查用户名唯一性失败")
	}
	if nameCount > 0 {
		return gerror.New("用户名已存在")
	}

	userId := gftoken.GetSessionUser(ctx).Id
	model.UpdateId = userId
	model.UpdateAt = gtime.Now()
	if in.Id > 0 {
		_, err = m.Where(columns.Id, model.Id).Update(model)
		if err != nil {
			return err
		}
		_ = Log.Save(ctx, model, consts.UPDATE)
		// 删除历史角色
		_, err = dao2.UserRole.Ctx(ctx).Where(dao2.UserRole.Columns().UserId, model.Id).Delete()
		if err != nil {
			return err
		}
	} else {
		model.CreateId = userId
		model.CreateAt = gtime.Now()
		model.UserType = consts.UserTypeNormal
		model.Avatar = consts.DefaultAvatar

		model.Uuid = guid.S()
		tmpStr, err := gmd5.Encrypt(consts.DefaultPassword)
		if err != nil {
			return err
		}
		salt := grand.Letters(6)
		model.Salt = salt
		model.Password, err = gmd5.Encrypt(tmpStr + salt)
		if err != nil {
			return err
		}
		modelId, err := m.InsertAndGetId(model)
		if err != nil {
			return err
		}
		model.Id = modelId
		_ = Log.SaveLog(ctx, &input2.LogData{
			Model:      model,
			OperType:   consts.INSERT,
			OperRemark: "角色ID：" + gconv.String(in.RoleIds),
		})
	}

	// 插入新角色
	userRoleList := g.List{}
	for _, roleId := range in.RoleIds {
		userRoleList = append(userRoleList,
			g.Map{dao2.UserRole.Columns().UserId: model.Id,
				dao2.UserRole.Columns().RoleId: roleId,
			})
	}
	_, err = dao2.UserRole.Ctx(ctx).Insert(userRoleList)
	return err
}

// Delete 删除用户
func (s *user) Delete(ctx context.Context, ids []int) error {
	// 删除用户角色关联
	_, err := dao2.UserRole.Ctx(ctx).WhereIn(dao2.UserRole.Columns().UserId, ids).Delete()
	if err != nil {
		return err
	}

	// 删除用户
	_, err = dao2.User.Ctx(ctx).WhereIn(dao2.User.Columns().Id, ids).Delete()
	if err != nil {
		return err
	}

	// 删除日志
	for _, id := range ids {
		_ = Log.Save(ctx, do.User{
			Id: gconv.Int64(id),
		}, consts.DELETE)
	}
	return nil
}

// PasswordReset 重置密码
func (s *user) PasswordReset(ctx context.Context, in *v1.UserPasswordResetReq) error {
	res, err := s.Get(ctx, in.Id)
	if err != nil {
		return err
	}
	password, err := gmd5.Encrypt(in.Password + res.Salt)
	if err != nil {
		return err
	}
	userId := gftoken.GetSessionUser(ctx).Id
	columns := dao2.User.Columns()
	_, err = dao2.User.Ctx(ctx).Where(columns.Id, in.Id).Update(do.User{
		UpdateId: userId,
		UpdateAt: gtime.Now(),
		Password: password,
	})
	return nil
}

// UserInfo 获取用户信息接口
func (s *login) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	var model *entity2.User
	err = dao2.User.Ctx(ctx).Where(dao2.User.Columns().UserName, gftoken.GetUserKey(ctx)).Scan(&model)
	if err != nil {
		return
	}
	if model == nil || model.Id <= 0 {
		err = gerror.NewCode(gcode.CodeValidationFailed, "用户名信息获取失败!")
		return
	}
	perms, err := getUserPerms(ctx, model.Id, model.UserType)
	if err != nil {
		return
	}
	roleNames, err := getUserRoleNames(ctx, model.Id, model.UserType)
	if err != nil {
		return
	}
	res = &v1.UserInfoRes{
		Avatar:   model.Avatar,
		Nickname: model.NickName,
		Roles:    roleNames,
		Perms:    perms,
		UserID:   gconv.String(model.Id),
		Username: model.UserName,
	}

	return
}

// UserMenus 获取用户菜单接口
func (s *login) UserMenus(ctx context.Context, in *v1.UserMenusReq) (res *v1.UserMenusRes, err error) {
	var (
		model   *entity2.User
		menus   []*entity2.Menu
		columns = dao2.Menu.Columns()
	)
	res = &v1.UserMenusRes{}
	m := dao2.Menu.Ctx(ctx).Where(columns.Enable, consts.EnableYes).WhereNot(columns.Type, consts.MenuTypeButton)
	// 获取当前用户
	err = dao2.User.Ctx(ctx).Where(dao2.User.Columns().UserName, gftoken.GetUserKey(ctx)).Scan(&model)
	if err != nil {
		return
	}
	if model.UserType == consts.UserTypeAdmin {
		m = m.OrderAsc(columns.Sort)
		if err = m.Scan(&menus); err != nil {
			err = gerror.Wrap(err, "获取数据失败！")
		}
	} else {
		roleIds, err2 := dao2.UserRole.Ctx(ctx).Fields(dao2.UserRole.Columns().RoleId).
			Where(dao2.UserRole.Columns().UserId, model.Id).Array()
		if err2 != nil {
			return nil, err2
		}
		if len(roleIds) == 0 {
			return
		}
		menuIds, err2 := dao2.RoleMenu.Ctx(ctx).Fields(dao2.RoleMenu.Columns().MenuId).
			WhereIn(dao2.RoleMenu.Columns().RoleId, gconv.SliceInt64(roleIds)).Array()
		if err2 != nil {
			return nil, err2
		}
		m = m.WhereIn(columns.Id, gconv.SliceInt64(menuIds))
		m = m.OrderAsc(columns.Sort)
		if err2 = m.Scan(&menus); err2 != nil {
			return nil, err2
		}
	}

	if len(menus) > 0 {
		tree := s.GetTree(0, menus)
		res = &tree
	}

	return
}

// getUserRoleNames 获取用户对应角色名称
func getUserRoleNames(ctx context.Context, userId int64, userType int) (res []string, err error) {
	if userType == consts.UserTypeAdmin {
		res = append(res, consts.RoleAdmin)
		return
	}

	values, err := dao2.UserRole.Ctx(ctx).Fields(dao2.UserRole.Columns().RoleId).Where(dao2.UserRole.Columns().UserId, userId).Array()
	if err != nil {
		return
	}
	var roles []*entity2.Role
	err = dao2.Role.Ctx(ctx).WhereIn(dao2.Role.Columns().Id, gconv.SliceInt64(values)).Scan(&roles)
	for _, e := range roles {
		res = append(res, e.Code)
	}
	return
}

// getUserPerms 获取用户对应按钮权限
func getUserPerms(ctx context.Context, userId int64, userType int) (res []string, err error) {
	// 管理员权限
	var menus []*entity2.Menu
	columns := dao2.Menu.Columns()
	m := dao2.Menu.Ctx(ctx).Where(columns.Enable, consts.EnableYes).Where(columns.Type, consts.MenuTypeButton)
	if userType == consts.UserTypeAdmin {
		// 管理员获取所有按钮权限
		if err = m.Scan(&menus); err != nil {
			return nil, err
		}
	} else {
		roleIds, err2 := dao2.UserRole.Ctx(ctx).Fields(dao2.UserRole.Columns().RoleId).
			Where(dao2.UserRole.Columns().UserId, userId).Array()
		if err2 != nil {
			return nil, err2
		}
		if len(roleIds) == 0 {
			return
		}
		menuIds, err2 := dao2.RoleMenu.Ctx(ctx).Fields(dao2.RoleMenu.Columns().MenuId).
			WhereIn(dao2.RoleMenu.Columns().RoleId, gconv.SliceInt64(roleIds)).Array()
		if err2 != nil {
			return nil, err2
		}
		if len(menuIds) == 0 {
			return
		}
		m = m.WhereIn(columns.Id, gconv.SliceInt64(menuIds))
		if err2 = m.Scan(&menus); err2 != nil {
			return nil, err2
		}
	}
	if len(menus) > 0 {
		for _, v := range menus {
			if v.Perm != "" {
				res = append(res, v.Perm)
			}
		}
	}
	return
}
