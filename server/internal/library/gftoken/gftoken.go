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

func init() {
	options := &gtoken.Options{}
	ctx := gctx.New()
	err := g.Cfg().MustGet(ctx, "gToken").Struct(options)
	if err != nil {
		panic("options init fail")
	}
	// 创建gtoken对象
	GToken = gtoken.NewDefaultToken(*options)
}

func MiddlewareAuth(r *ghttp.Request) {
	gtoken.NewDefaultMiddleware(GToken,
		"/admin/login", "/admin/captcha/get").Auth(r)
}

func GetUserKey(ctx context.Context) string {
	return g.RequestFromCtx(ctx).GetCtxVar(gtoken.KeyUserKey).String()
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
	userKey := g.RequestFromCtx(ctx).GetCtxVar(gtoken.KeyUserKey).String()
	_, data, err = GToken.Get(ctx, userKey)
	return
}
