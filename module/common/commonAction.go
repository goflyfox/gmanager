package common

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"gmanager/module/constants"
)

// Login 登录页面
func Welcome(r *ghttp.Request) {
	r.Response.WriteTpl("pages/welcome.html", g.Map{})
}

// Login 登录页面
func Debug(r *ghttp.Request) {
	if constants.DEBUG {
		constants.DEBUG = false
		g.DB().SetDebug(false)
		r.Response.Writeln("debug close ~!~ ")
	} else {
		constants.DEBUG = true
		g.DB().SetDebug(true)
		r.Response.Writeln("debug open ~!~ ")
	}
}
