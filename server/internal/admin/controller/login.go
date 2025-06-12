package controller

import (
	"context"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/admin/logic"
)

var Login = new(login)

type login struct{}

// Login 登录接口
func (c *login) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	res, err = logic.Login.Login(ctx, req)
	return
}

// Logout 登出接口
func (c *login) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	res, err = logic.Login.Logout(ctx, req)
	return
}

// CaptchaGet 验证码获取
func (c *login) CaptchaGet(ctx context.Context, req *v1.CaptchaReq) (res *v1.CaptchaRes, err error) {
	res, err = logic.Login.CaptchaGet(ctx, req)
	return
}
