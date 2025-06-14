package logic

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/admin/consts"
	"gmanager/internal/admin/dao"
	"gmanager/internal/admin/model/do"
	entity2 "gmanager/internal/admin/model/entity"
	input2 "gmanager/internal/admin/model/input"
	"gmanager/internal/library/bean"
	"gmanager/internal/library/captcha"
	"gmanager/internal/library/gftoken"
)

var Login = login{}

type login struct{}

// Login 登录接口
func (s *login) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	if !captcha.Verify(req.CodeId, req.Code) {
		err = gerror.NewCode(gcode.CodeValidationFailed, "验证码错误")
		return
	}

	var model *entity2.User
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().UserName, req.Username).Scan(&model)
	if err != nil {
		return
	}

	if model == nil || model.Id <= 0 {
		err = gerror.NewCode(gcode.CodeValidationFailed, "用户名或者密码错误!")
		return
	}

	if model.Enable != consts.EnableYes {
		err = gerror.NewCode(gcode.CodeValidationFailed, "账号状态异常，请联系管理")
		return
	}

	reqPassword, err := gmd5.Encrypt(req.Password + model.Salt)
	if err != nil {
		err = gerror.Wrap(err, "用户名或者密码错误!")
		return
	}

	if reqPassword != model.Password {
		err = gerror.NewCode(gcode.CodeValidationFailed, "用户名或者密码错误!")
		return
	}

	res = &v1.LoginRes{
		Username:  req.Username,
		TokenType: "Bearer",
	}
	sessionUser := bean.SessionUser{
		Id:       model.Id,
		Uuid:     model.Uuid,
		NickName: model.NickName,
		Username: model.UserName,
	}
	// 认证成功调用Generate生成Token
	res.AccessToken, err = gftoken.GToken.Generate(ctx, req.Username, sessionUser)
	if err != nil {
		return
	}
	// 记录日志
	_ = Log.SaveLog(ctx, &input2.LogData{
		Model: do.User{
			Id:       model.Id,
			UserName: model.UserName,
			UpdateAt: gtime.Now(),
			UpdateId: model.Id,
		},
		OperRemark: model.NickName,
		Operator:   model.UserName,
		OperType:   consts.LOGIN,
	})
	return
}

// Logout 登出接口
func (s *login) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	userName := gftoken.GetUserKey(ctx)
	if userName == "" {
		return
	}
	// 记录日志
	var model *entity2.User
	userErr := dao.User.Ctx(ctx).Where(dao.User.Columns().UserName, userName).Scan(&model)
	if userErr != nil {
		return
	}
	_ = Log.SaveLog(ctx, &input2.LogData{
		Model: do.User{
			Id:       model.Id,
			UserName: model.UserName,
			UpdateAt: gtime.Now(),
			UpdateId: model.Id,
		},
		OperRemark: model.NickName,
		Operator:   model.UserName,
		OperType:   consts.LOGOUT,
	})
	// 销毁Token
	_ = gftoken.GToken.Destroy(ctx, userName)
	return
}

// GetTree 菜单树形菜单
func (s *login) GetTree(pid int64, list []*entity2.Menu) (tree []*input2.UserMenu) {
	tree = make([]*input2.UserMenu, 0, len(list))
	for _, v := range list {
		if v.ParentId == pid {
			name := v.RouteName
			if name == "" {
				name = v.RoutePath
			}
			t := &input2.UserMenu{
				Id:        v.Id,
				Name:      name,
				Component: v.Component,
				Path:      v.RoutePath,
				Redirect:  v.Redirect,
				Meta: input2.Meta{
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

func (s *login) CaptchaGet(ctx context.Context, req *v1.CaptchaReq) (res *v1.CaptchaRes, err error) {
	res = &v1.CaptchaRes{}
	res.CodeId, res.Img, err = captcha.Generate(ctx)
	return
}

// CaptchaVerify 验证输入的验证码是否正确
func (s *login) CaptchaVerify(id, answer string) bool {
	return captcha.Verify(id, answer)
}
