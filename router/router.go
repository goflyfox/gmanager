package router

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"gmanager/app/api"
	"gmanager/app/api/common"
	"gmanager/app/api/config"
	"gmanager/app/api/log"
	"gmanager/app/api/menu"
	"gmanager/app/api/role"
	"gmanager/app/api/user"
	"gmanager/app/component/middle"
	"gmanager/app/component/started"
	"gmanager/app/constants"
	"gmanager/library/base"
	"strings"
)

/*
绑定业务路由
*/
func bindRouter() {
	urlPath := g.Config().GetString("url-path")
	s := g.Server()

	s.Group(urlPath+"/", func(group *ghttp.RouterGroup) {
		// 允许跨域
		group.Middleware(func(r *ghttp.Request) {
			r.Response.CORSDefault()
			r.Middleware.Next()
		})
		// 日志拦截
		group.Middleware(middle.MiddlewareLog)
		// 通用属性
		group.Middleware(middle.MiddlewareCommon)

		// 首页
		group.ALL("/", common.Login)
		group.ALL("/main.html", common.Index)
		group.ALL("/login", common.Login)

		group.ALL("/welcome", common.Welcome)
		group.ALL("/admin/welcome.html", common.Welcome)

		// 启动gtoken
		base.Token = &gtoken.GfToken{
			//Timeout:         10 * 1000,
			CacheMode:        g.Config().GetInt8("gtoken.cache-mode"),
			MultiLogin:       g.Config().GetBool("gtoken.multi-login"),
			LoginPath:        "/login/submit",
			LoginBeforeFunc:  common.LoginSubmit,
			LogoutPath:       "/user/logout",
			LogoutBeforeFunc: common.LogoutBefore,
			AuthPaths:        g.SliceStr{"/user", "/system"},
			GlobalMiddleware: true,
			AuthBeforeFunc: func(r *ghttp.Request) bool {
				// 静态页面不拦截
				if r.IsFileRequest() {
					return false
				}

				if strings.HasSuffix(r.URL.Path, "index") {
					return false
				}

				return true
			},
		}
		// 需要认证
		group.Group(urlPath+"/system", func(group *ghttp.RouterGroup) {
			// gtoken认证中间件
			base.Token.Middleware(group)

			// 系统路由
			userAction := new(user.Action)
			group.ALL("user", userAction)
			group.GET("/user/get/{id}", userAction.Get)
			group.ALL("user/delete/{id}", userAction.Delete)

			group.ALL("department", api.Department)
			group.GET("/department/get/{id}", api.Department.Get)
			group.ALL("/department/delete/{id}", api.Department.Delete)

			logAction := new(log.Action)
			group.ALL("log", logAction)
			group.GET("/log/get/{id}", logAction.Get)
			group.ALL("/log/delete/{id}", logAction.Delete)

			menuAction := new(menu.Action)
			group.ALL("menu", menuAction)
			group.GET("/menu/get/{id}", menuAction.Get)
			group.ALL("/menu/delete/{id}", menuAction.Delete)

			roleAction := new(role.Action)
			group.ALL("role", roleAction)
			group.GET("/role/get/{id}", roleAction.Get)
			group.ALL("/role/delete/{id}", roleAction.Delete)

			configAction := new(config.Action)
			group.ALL("config", configAction)
			group.GET("/config/get/{id}", configAction.Get)
			group.ALL("/config/delete/{id}", configAction.Delete)
		})
	})

}

/*
统一路由注册
*/
func init() {
	glog.Info("########router start...")

	s := g.Server()

	// 绑定路由
	bindRouter()

	if constants.DEBUG {
		g.DB().SetDebug(constants.DEBUG)
	}

	// 上线建议关闭
	s.BindHandler("/debug", common.Debug)

	// 301错误页面
	s.BindStatusHandler(301, common.Error301)
	// 404错误页面
	s.BindStatusHandler(404, common.Error404)
	// 500错误页面
	s.BindStatusHandler(500, common.Error500)

	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	s.SetRewrite("/favicon.ico", "/resources/images/favicon.ico")

	// 管理接口
	s.EnableAdmin("/admin")

	// 为平滑重启管理页面设置HTTP Basic账号密码
	//s.BindHookHandler("/admin/*", ghttp.HOOK_BEFORE_SERVE, func(r *ghttp.Request) {
	//	user := g.Config().GetString("admin.user")
	//	pass := g.Config().GetString("admin.pass")
	//	if !r.BasicAuth(user, pass) {
	//		r.ExitAll()
	//	}
	//})

	// 强制跳转到HTTPS访问
	//g.Server().BindHookHandler("/*", ghttp.HOOK_BEFORE_SERVE, func(r *ghttp.Request) {
	//    if !r.IsFileServe() && r.TLS == nil {
	//        r.Response.RedirectTo(fmt.Sprintf("https://%s%s", r.Host, r.URL.String()))
	//        r.ExitAll()
	//    }
	//})

	started.StartLog()

	glog.Info("########router finish.")
}
