package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"gmanager/module/component/started"
)

// 管理初始化顺序.
func init() {
	initConfig()
	initRouter()
	started.Start()
}

// 用于配置初始化.
func initConfig() {
	glog.Info("########service start...")

	v := g.View()
	c := g.Config()
	s := g.Server()

	// 配置对象及视图对象配置
	c.AddPath("config")

	v.SetDelimiters("${", "}")
	v.AddPath("template")

	glog.SetStdoutPrint(true)

	s.SetNameToUriType(ghttp.URI_TYPE_ALLLOWER)

	glog.Info("########service finish.")

}
