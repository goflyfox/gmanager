package hook

import (
	"fmt"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

func AuthAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		var params map[string]string
		if r.Method == "GET" {
			params = r.GetQueryMap()
		} else if r.Method == "POST" {
			params = r.GetPostMap()
		} else {
			r.Response.Writeln("Request Method is ERROR! ")
			return
		}

		no := gconv.String(gtime.Millisecond())

		glog.Info(fmt.Sprintf("[AUTH_%s_%d][url:%s][params:%s][data:%s]",
			no, r.Id, r.URL.Path, params, respData.Json()))
		respData.Msg = "请求错误或登录超时"
		r.Response.WriteJson(respData)
		r.ExitAll()
	}
}

func AuthBeforeFunc(r *ghttp.Request) bool {
	// 静态页面不拦截
	if r.IsFileRequest() {
		return false
	}
	// 静态页面不拦截
	if r.IsAjaxRequest() {
		return true
	}

	return false
}
