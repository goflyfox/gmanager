package boot

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"gmanager/gtoken"
	"gmanager/module/common"
	"gmanager/module/component/hook"
	"gmanager/module/constants"
	"gmanager/module/system"
	"gmanager/utils/base"
)

/*
绑定业务路由
*/
func bindRouter() {
	urlPath := g.Config().GetString("url-path")

	// 首页
	g.Server().BindHandler(urlPath+"/", common.Login)
	g.Server().BindHandler(urlPath+"/main.html", common.Index)
	g.Server().BindHandler(urlPath+"/login", common.Login)

	g.Server().BindHandler(urlPath+"/admin/welcome.html", common.Welcome)
	// 调试日志
	//g.Server().BindObject(urlPath+"/tmp", new(adminAction.TmpAction))
	//

	// 系统路由
	userAction := new(system.UserAction)
	g.Server().BindObject(urlPath+"/system/user", userAction)
	g.Server().BindObjectMethod(urlPath+"/system/user/get/{id}", userAction, "Get")
	g.Server().BindObjectMethod(urlPath+"/system/user/delete/{id}", userAction, "Delete")

	departAction := new(system.DepartmentAction)
	g.Server().BindObject(urlPath+"/system/department", departAction)
	g.Server().BindObjectMethod(urlPath+"/system/department/get/{id}", departAction, "Get")
	g.Server().BindObjectMethod(urlPath+"/system/department/delete/{id}", departAction, "Delete")

	logAction := new(system.LogAction)
	g.Server().BindObject(urlPath+"/system/log", logAction)
	g.Server().BindObjectMethod(urlPath+"/system/log/get/{id}", logAction, "Get")
	g.Server().BindObjectMethod(urlPath+"/system/log/delete/{id}", logAction, "Delete")

	menuAction := new(system.MenuAction)
	g.Server().BindObject(urlPath+"/system/menu", menuAction)
	g.Server().BindObjectMethod(urlPath+"/system/menu/get/{id}", menuAction, "Get")
	g.Server().BindObjectMethod(urlPath+"/system/menu/delete/{id}", menuAction, "Delete")

	roleAction := new(system.RoleAction)
	g.Server().BindObject(urlPath+"/system/role", roleAction)
	g.Server().BindObjectMethod(urlPath+"/system/role/get/{id}", roleAction, "Get")
	g.Server().BindObjectMethod(urlPath+"/system/role/delete/{id}", roleAction, "Delete")

	configAction := new(system.ConfigAction)
	g.Server().BindObject(urlPath+"/system/config", configAction)
	g.Server().BindObjectMethod(urlPath+"/system/config/get/{id}", configAction, "Get")
	g.Server().BindObjectMethod(urlPath+"/system/config/delete/{id}", configAction, "Delete")

	authPaths := g.SliceStr{"/user/*", "/system/*"}

	// 启动gtoken
	base.Token = &gtoken.GfToken{
		//Timeout:         10 * 1000,
		CacheMode:        g.Config().GetInt8("cache-mode"),
		LoginPath:        "/login/submit",
		LoginBeforeFunc:  common.LoginSubmit,
		LogoutPath:       "/user/logout",
		LogoutBeforeFunc: common.LogoutBefore,
		AuthPaths:        authPaths,
		AuthBeforeFunc:   hook.AuthBeforeFunc,
		AuthAfterFunc:    hook.AuthAfterFunc,
	}
	base.Token.Start()
}

/*
统一路由注册
*/
func initRouter() {

	s := g.Server()

	//// 通用设置
	s.BindHookHandler("/*any", ghttp.HOOK_BEFORE_SERVE, hook.CommonBefore)

	// 日志拦截
	s.BindHookHandlerByMap("/*any", map[string]ghttp.HandlerFunc{
		ghttp.HOOK_BEFORE_SERVE: hook.LogBeforeServe,
		ghttp.HOOK_AFTER_SERVE:  hook.LogBeforeOutput,
	})

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
	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")

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
}
