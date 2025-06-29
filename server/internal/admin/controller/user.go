package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/admin/logic"
	"strings"
)

type user struct{}

var User = new(user)

func (c *user) List(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	res, err = logic.User.List(ctx, req)
	return
}

func (c *user) Get(ctx context.Context, req *v1.UserGetReq) (res *v1.UserGetRes, err error) {
	res, err = logic.User.Get(ctx, req.Id)
	return
}

func (c *user) Save(ctx context.Context, req *v1.UserSaveReq) (res *v1.UserSaveRes, err error) {
	err = logic.User.Save(ctx, req)
	return
}

func (c *user) Delete(ctx context.Context, req *v1.UserDeleteReq) (res *v1.UserDeleteRes, err error) {
	if req.Ids == "" {
		return
	}
	idArr := make([]int, 0)
	for _, v := range strings.Split(req.Ids, ",") {
		idArr = append(idArr, gconv.Int(v))
	}
	err = logic.User.Delete(ctx, idArr)
	return
}

func (c *user) PasswordReset(ctx context.Context, req *v1.UserPasswordResetReq) (res *v1.UserPasswordResetRes, err error) {
	err = logic.User.PasswordReset(ctx, req)
	return
}

// UserInfo 获取用户信息接口
func (c *user) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	res, err = logic.User.UserInfo(ctx, req)
	return
}

// UserMenus 获取用户菜单接口
func (c *user) UserMenus(ctx context.Context, req *v1.UserMenusReq) (res *v1.UserMenusRes, err error) {
	res, err = logic.User.UserMenus(ctx, req)
	return
}

func (c *user) Export(ctx context.Context, req *v1.UserExportReq) (res *v1.UserExportRes, err error) {
	err = logic.User.Export(ctx, req)
	return
}

func (c *user) Import(ctx context.Context, req *v1.UserImportReq) (res *v1.UserImportRes, err error) {
	res, err = logic.User.Import(ctx, req)
	return
}

func (c *user) Template(ctx context.Context, req *v1.UserTemplateReq) (res *v1.UserTemplateRes, err error) {
	err = logic.User.Template(ctx, req)
	return
}

func (c *user) Profile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error) {
	res, err = logic.User.Profile(ctx, req)
	return
}

func (c *user) SaveProfile(ctx context.Context, req *v1.UserSaveProfileReq) (res *v1.UserSaveProfileRes, err error) {
	err = logic.User.SaveProfile(ctx, req)
	return
}

func (c *user) ChangePassword(ctx context.Context, req *v1.UserChangePasswordReq) (res *v1.UserChangePasswordRes, err error) {
	if req.OldPassword == req.NewPassword {
		return nil, gerror.New("原密码和新密码一致")
	}
	err = logic.User.ChangePassword(ctx, req)
	return
}

func (c *user) SendMobileCode(ctx context.Context, req *v1.UserSendMobileCodeReq) (res *v1.UserSendMobileCodeRes, err error) {
	err = logic.User.SendMobileCode(ctx, req)
	return
}

func (c *user) SaveMobile(ctx context.Context, req *v1.UserSaveMobileReq) (res *v1.UserSaveMobileRes, err error) {
	err = logic.User.SaveMobile(ctx, req)
	return
}

func (c *user) SendEmailCode(ctx context.Context, req *v1.UserSendEmailCodeReq) (res *v1.UserSendEmailCodeRes, err error) {
	err = logic.User.SendEmailCode(ctx, req)
	return
}

func (c *user) SaveEmail(ctx context.Context, req *v1.UserSaveEmailReq) (res *v1.UserSaveEmailRes, err error) {
	err = logic.User.SaveEmail(ctx, req)
	return
}
