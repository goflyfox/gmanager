package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"gmanager/internal/admin/controller"
	"gmanager/internal/admin/middleware"
	common "gmanager/internal/common/controller"
	"gmanager/internal/library/cache"
	"gmanager/internal/library/gftoken"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.BindStatusHandler(301, controller.Error301) // 301错误页面
			s.BindStatusHandler(404, controller.Error404) // 404错误页面
			s.BindStatusHandler(500, controller.Error500) // 500错误页面

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Bind(
					controller.Ping,
				)
			})

			s.Group("/admin", func(group *ghttp.RouterGroup) {
				//跨域处理，安全起见正式环境请注释该行
				group.Middleware(func(r *ghttp.Request) {
					r.Response.CORSDefault()
					r.Middleware.Next()
				})
				group.Middleware(gftoken.MiddlewareAuth) // gtoken登陆认证
				group.Middleware(middleware.UserPerm)    // 用户权限认证
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(middleware.DemoNotice) // 演示环境提示
				group.Bind(
					common.Upload,
					controller.Login,
					controller.Dept,
					controller.User,
					controller.Role,
					controller.Config,
					controller.Menu,
					controller.Log,
				)
			})

			initData(ctx)

			s.Run()
			return nil
		},
	}
)

func initData(ctx context.Context) {
	cache.InitCache(ctx)
}
