package middle

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"gmanager/app/constants"
	"gmanager/library/base"
)

func MiddlewareLog(r *ghttp.Request) {
	var beforeTime int64
	var params map[string]interface{}
	var no string

	if constants.DEBUG {
		beforeTime = gtime.TimestampMilli()

		if r.IsFileRequest() {
			return
		}

		if r.Method == "GET" {
			params = r.GetMap()
		} else if r.Method == "POST" {
			params = r.GetMap()
		} else {
			base.Error(r, "Request Method is ERROR! ")
			return
		}

		no = gconv.String(params["no"])
		if no == "" {
			no = gconv.String(beforeTime)
		}

		glog.Info(fmt.Sprintf("[REQUEST_%s][url:%s][params:%s]",
			no, r.URL.Path, params))
	}

	r.Middleware.Next()

	// 青牛完成
	if constants.DEBUG {
		data := string(r.Response.Buffer())
		if r.URL.Path == "" || r.URL.Path == "/" || gstr.Contains(
			r.URL.Path, "index") || gstr.Contains(
			r.URL.Path, "html") || gstr.Contains(
			r.URL.Path, "login") {
			data = ""
		}

		afterTime := gtime.TimestampMilli()

		if r.IsFileRequest() {
			glog.Info(fmt.Sprintf("[FILE_%s][diff:%d][url:%s][params:%s]",
				no, afterTime-beforeTime, r.URL.Path, params))
		} else if afterTime-beforeTime > 1000 {
			glog.Warning(fmt.Sprintf("[RESPONSE_%s][diff:%d][url:%s][params:%s][data:%s]",
				no, afterTime-beforeTime, r.URL.Path, params, data))
		} else {
			glog.Info(fmt.Sprintf("[RESPONSE_%s][diff:%d][url:%s][params:%s][data:%s]",
				no, afterTime-beforeTime, r.URL.Path, params, data))
		}
	}
}

func MiddlewareCommon(r *ghttp.Request) {
	r.SetParam("BASE_PATH", "")

	r.Middleware.Next()
}
