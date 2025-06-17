package gftoken

import (
	"context"
	"github.com/goflyfox/gtoken/v2/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"gmanager/internal/library/bean"
)

var GToken gtoken.Token
var m gtoken.Middleware

func init() {
	options := &gtoken.Options{}
	ctx := gctx.New()
	err := g.Cfg().MustGet(ctx, "gToken").Struct(options)
	if err != nil {
		panic("options init fail")
	}
	// 创建gtoken对象
	GToken = gtoken.NewDefaultToken(*options)
	m = gtoken.NewDefaultMiddleware(GToken,
		"/admin/login", "/admin/captcha/get")
}

func MiddlewareAuth(r *ghttp.Request) {
	m.Auth(r)
}

func HasExcludePath(r *ghttp.Request) bool {
	return m.HasExcludePath(r)
}

func GetUserKey(ctx context.Context) string {
	return g.RequestFromCtx(ctx).GetCtxVar(gtoken.KeyUserKey).String()
}

func GetUserId(ctx context.Context) int64 {
	user := GetSessionUser(ctx)
	if user == nil {
		return 0
	}
	return user.Id
}

func GetSessionUser(ctx context.Context) *bean.SessionUser {
	var user *bean.SessionUser
	data, err := GetData(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil
	}
	err = gconv.Struct(data, &user)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil
	}
	return user
}

func GetData(ctx context.Context) (data any, err error) {
	userKey := g.RequestFromCtx(ctx).GetCtxVar(gtoken.KeyUserKey)
	if userKey.IsNil() {
		return
	}
	_, data, err = GToken.Get(ctx, userKey.String())
	return
}
