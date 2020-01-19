package router

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"gmanager/app/api/common"
	"gmanager/app/api/config"
	"gmanager/app/api/department"
	"gmanager/app/api/log"
	"gmanager/app/api/menu"
	"gmanager/app/api/role"
	"gmanager/app/api/user"
	"gmanager/app/constants"
	"gmanager/module/component/middle"
	"gmanager/module/component/started"
	"gmanager/utils/base"
	"strings"
)

/*
绑定业务路由
*/
func bindRouter() {
	urlPath := g.Config().GetString("url-path")
	s := g.Server()
	// 首页
	s.BindHandler(urlPath+"/", common.Login)
	s.BindHandler(urlPath+"/main.html", common.Index)
	s.BindHandler(urlPath+"/login", common.Login)

	s.BindHandler(urlPath+"/welcome", common.Welcome)
	s.BindHandler(urlPath+"/admin/welcome.html", common.Welcome)
	// 中间件
	s.BindMiddleware(urlPath+"/*", middle.MiddlewareLog)
	s.BindMiddleware(urlPath+"/*", middle.MiddlewareCommon)

	s.Group(urlPath+"/system", func(g *ghttp.RouterGroup) {
		// 允许跨域
		g.Middleware(func(r *ghttp.Request) {
			r.Response.CORSDefault()
			r.Middleware.Next()
		})

		// 系统路由
		userAction := new(user.Action)
		g.ALL("user", userAction)
		g.GET("/user/get/{id}", userAction.Get)
		g.ALL("user/delete/{id}", userAction.Delete)

		departAction := new(department.Action)
		g.ALL("department", departAction)
		g.GET("/department/get/{id}", departAction.Get)
		g.ALL("/department/delete/{id}", departAction.Delete)

		logAction := new(log.Action)
		g.ALL("log", logAction)
		g.GET("/log/get/{id}", logAction.Get)
		g.ALL("/log/delete/{id}", logAction.Delete)

		menuAction := new(menu.Action)
		g.ALL("menu", menuAction)
		g.GET("/menu/get/{id}", menuAction.Get)
		g.ALL("/menu/delete/{id}", menuAction.Delete)

		roleAction := new(role.Action)
		g.ALL("role", roleAction)
		g.GET("/role/get/{id}", roleAction.Get)
		g.ALL("/role/delete/{id}", roleAction.Delete)

		configAction := new(config.Action)
		g.ALL("config", configAction)
		g.GET("/config/get/{id}", configAction.Get)
		g.ALL("/config/delete/{id}", configAction.Delete)

	})

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
	base.Token.Start()
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

	started.Start()

	glog.Info("########router finish.")
}
