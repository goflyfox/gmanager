package logic

import (
	"context"
	"fmt"
	"gmanager/api/admin/v1"
	"gmanager/internal/admin/consts"
	"gmanager/internal/admin/dao"
	"gmanager/internal/admin/model/do"
	"gmanager/internal/admin/model/entity"
	"gmanager/internal/admin/model/input"
	"gmanager/internal/library/cache"
	"gmanager/internal/library/gftoken"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/guid"
)

// User 用户服务
var User = new(user)

type user struct{}

// List 获取用户列表
func (s *user) List(ctx context.Context, in *v1.UserListReq) (res *v1.UserListRes, err error) {
	if in == nil {
		return
	}
	m := dao.User.Ctx(ctx)
	columns := dao.User.Columns()
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

	if in.NeedOrderBy() {
		m = m.Order(in.OrderBy)
	} else {
		m = m.Order("id desc")
	}
	var pageList []*input.User
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
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Scan(&res)
	if err != nil {
		return
	}
	values, err := dao.UserRole.Ctx(ctx).Fields(dao.UserRole.Columns().RoleId).Where(dao.UserRole.Columns().UserId, id).Array()
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

	m := dao.User.Ctx(ctx)
	columns := dao.User.Columns()

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
		_, err = dao.UserRole.Ctx(ctx).Where(dao.UserRole.Columns().UserId, model.Id).Delete()
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
		_ = Log.SaveLog(ctx, &input.LogData{
			Model:      model,
			OperType:   consts.INSERT,
			OperRemark: "角色ID：" + gconv.String(in.RoleIds),
		})
	}

	// 插入新角色
	userRoleList := g.List{}
	for _, roleId := range in.RoleIds {
		userRoleList = append(userRoleList,
			g.Map{dao.UserRole.Columns().UserId: model.Id,
				dao.UserRole.Columns().RoleId: roleId,
			})
	}
	_, err = dao.UserRole.Ctx(ctx).Insert(userRoleList)
	return err
}

// Delete 删除用户
func (s *user) Delete(ctx context.Context, ids []int) error {
	// 删除用户角色关联
	_, err := dao.UserRole.Ctx(ctx).WhereIn(dao.UserRole.Columns().UserId, ids).Delete()
	if err != nil {
		return err
	}

	// 删除用户
	_, err = dao.User.Ctx(ctx).WhereIn(dao.User.Columns().Id, ids).Delete()
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
	columns := dao.User.Columns()
	_, err = dao.User.Ctx(ctx).Where(columns.Id, in.Id).Update(do.User{
		UpdateId: userId,
		UpdateAt: gtime.Now(),
		Password: password,
	})
	return nil
}

// UserInfo 获取用户信息接口
func (s *user) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	var model *entity.User
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().UserName, gftoken.GetUserKey(ctx)).Scan(&model)
	if err != nil {
		return
	}
	if model == nil || model.Id <= 0 {
		err = gerror.NewCode(gcode.CodeValidationFailed, "用户名信息获取失败!")
		return
	}
	perms, err := getUserPerms(ctx, false)
	if err != nil {
		return
	}
	roleNames, err := getUserRoleNames(ctx)
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
func (s *user) UserMenus(ctx context.Context, in *v1.UserMenusReq) (res *v1.UserMenusRes, err error) {
	var (
		menus []*entity.Menu
	)
	menus, err = getUserMenus(ctx, false)
	if err != nil {
		return
	}
	if len(menus) > 0 {
		tree := s.GetTree(0, menus)
		res = &tree
	}
	return
}

/** getUserMenus 获取用户菜单
 * @param cacheFlag 是否缓存
 */
func getUserMenus(ctx context.Context, cacheFlag bool) (menus []*entity.Menu, err error) {
	var (
		columns = dao.Menu.Columns()
	)
	sessionUser := gftoken.GetSessionUser(ctx)
	if sessionUser == nil {
		return
	}
	if cacheFlag { // 使用缓存
		cacheMap, err2 := cache.Instance().Get(ctx, fmt.Sprintf(consts.CacheUserMenu, sessionUser.Id))
		if err2 != nil {
			return nil, err2
		}
		if len(cacheMap) > 0 {
			err = gconv.Struct(cacheMap[consts.CacheData], &menus)
			if err != nil {
				return nil, err
			}
			if len(menus) > 0 {
				return
			}
		}
	}

	m := dao.Menu.Ctx(ctx).Where(columns.Enable, consts.EnableYes).WhereNot(columns.Type, consts.MenuTypeButton)
	if sessionUser.UserType == consts.UserTypeAdmin {
		m = m.OrderAsc(columns.Sort)
		if err = m.Scan(&menus); err != nil {
			err = gerror.Wrap(err, "获取数据失败！")
		}
	} else {
		roleIds, err2 := dao.UserRole.Ctx(ctx).Fields(dao.UserRole.Columns().RoleId).
			Where(dao.UserRole.Columns().UserId, sessionUser.Id).Array()
		if err2 != nil {
			return nil, err2
		}
		if len(roleIds) == 0 {
			return
		}
		menuIds, err2 := dao.RoleMenu.Ctx(ctx).Fields(dao.RoleMenu.Columns().MenuId).
			WhereIn(dao.RoleMenu.Columns().RoleId, gconv.SliceInt64(roleIds)).Array()
		if err2 != nil {
			return nil, err2
		}
		m = m.WhereIn(columns.Id, gconv.SliceInt64(menuIds))
		m = m.OrderAsc(columns.Sort)
		if err2 = m.Scan(&menus); err2 != nil {
			return nil, err2
		}
	}
	// 设置缓存
	if len(menus) > 0 {
		err = cache.Instance().Set(ctx, fmt.Sprintf(consts.CacheUserMenu, sessionUser.Id), g.Map{consts.CacheData: menus})
		if err != nil {
			return
		}
	}
	return
}

// GetTree 菜单树形菜单
func (s *user) GetTree(pid int64, list []*entity.Menu) (tree []*input.UserMenu) {
	tree = make([]*input.UserMenu, 0, len(list))
	for _, v := range list {
		if v.ParentId == pid {
			name := v.RouteName
			if name == "" {
				name = v.RoutePath
			}
			t := &input.UserMenu{
				Id:        v.Id,
				Name:      name,
				Component: v.Component,
				Path:      v.RoutePath,
				Redirect:  v.Redirect,
				Meta: input.Meta{
					AlwaysShow: v.AlwaysShow == 1,
					Hidden:     v.Enable != 1,
					Icon:       v.Icon,
					KeepAlive:  v.KeepAlive == 1,
					Title:      v.Name,
				},
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

// getUserRoleNames 获取用户对应角色名称
func getUserRoleNames(ctx context.Context) (res []string, err error) {
	sessionUser := gftoken.GetSessionUser(ctx)
	if sessionUser == nil {
		return
	}
	if sessionUser.UserType == consts.UserTypeAdmin {
		res = append(res, consts.RoleAdmin)
		return
	}

	values, err := dao.UserRole.Ctx(ctx).Fields(dao.UserRole.Columns().RoleId).Where(dao.UserRole.Columns().UserId, sessionUser.Id).Array()
	if err != nil {
		return
	}
	var roles []*entity.Role
	err = dao.Role.Ctx(ctx).WhereIn(dao.Role.Columns().Id, gconv.SliceInt64(values)).Scan(&roles)
	for _, e := range roles {
		res = append(res, e.Code)
	}
	return
}

/** getUserPerms 获取用户对应按钮权限
 * @param cacheFlag 是否使用缓存
 */
func getUserPerms(ctx context.Context, cacheFlag bool) (res []string, err error) {
	sessionUser := gftoken.GetSessionUser(ctx)
	if sessionUser == nil {
		return
	}

	if cacheFlag { // 使用缓存
		cacheMap, err2 := cache.Instance().Get(ctx, fmt.Sprintf(consts.CacheUserPerm, sessionUser.Id))
		if err2 != nil {
			return nil, err2
		}
		if len(cacheMap) > 0 {
			res = gconv.SliceStr(cacheMap[consts.CacheData])
			if len(res) > 0 {
				return
			}
		}
	}

	// 管理员权限
	var menus []*entity.Menu
	columns := dao.Menu.Columns()
	m := dao.Menu.Ctx(ctx).Where(columns.Enable, consts.EnableYes).Where(columns.Type, consts.MenuTypeButton)
	if sessionUser.UserType == consts.UserTypeAdmin {
		// 管理员获取所有按钮权限
		if err = m.Scan(&menus); err != nil {
			return nil, err
		}
	} else {
		roleIds, err2 := dao.UserRole.Ctx(ctx).Fields(dao.UserRole.Columns().RoleId).
			Where(dao.UserRole.Columns().UserId, sessionUser.Id).Array()
		if err2 != nil {
			return nil, err2
		}
		if len(roleIds) == 0 {
			return
		}
		menuIds, err2 := dao.RoleMenu.Ctx(ctx).Fields(dao.RoleMenu.Columns().MenuId).
			WhereIn(dao.RoleMenu.Columns().RoleId, gconv.SliceInt64(roleIds)).Array()
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
	// 设置缓存
	if len(res) > 0 {
		err = cache.Instance().Set(ctx, fmt.Sprintf(consts.CacheUserPerm, sessionUser.Id), g.Map{consts.CacheData: res})
		if err != nil {
			return
		}
	}
	return
}

func (s *user) CheckPerm(ctx context.Context, perm string) bool {
	perms, err := getUserPerms(ctx, true)
	if err != nil {
		g.Log().Info(ctx, err)
		return false
	}
	for _, userPerm := range perms {
		if userPerm == perm {
			return true
		}
	}
	return false
}

// CheckUrl 菜单校验
// TODO 只有用户菜单权限，但是用户菜单以来组织机构和角色？？？
func (s *user) CheckUrl(ctx context.Context, path string) bool {
	menus, err := getUserMenus(ctx, true)
	if err != nil {
		g.Log().Info(ctx, err)
		return false
	}
	for _, e := range menus {
		if e.Type != consts.MenuTypeMenu {
			continue
		}
		if gstr.HasPrefix(path, "/admin/"+e.RoutePath) {
			return true
		}
	}
	return false
}
