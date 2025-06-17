package middleware

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

// DemoNotice 演示环境提示
func DemoNotice(r *ghttp.Request) {
	r.Middleware.Next()

	var (
		msg string
		err = r.GetError()
	)
	if err != nil {
		msg = err.Error()
		if gstr.Contains(msg, "denied to user") {
			r.SetError(gerror.New("演示环境，禁止操作"))
		}
	}
}
