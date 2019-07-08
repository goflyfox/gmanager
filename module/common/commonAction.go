package common

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"gmanager/module/constants"
)

// Login 登录页面
func Welcome(r *ghttp.Request) {
	err := r.Response.WriteTpl("pages/welcome.html", g.Map{})

	if err != nil {
		glog.Error(err)
	}
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
