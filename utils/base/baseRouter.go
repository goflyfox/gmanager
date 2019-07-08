package base

import (
	"github.com/gogf/gf/g/encoding/gjson"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"gmanager/gtoken"
	"gmanager/utils/bean"
	"gmanager/utils/resp"
)

var Token *gtoken.GfToken

// baseRouter implemented global settings for all other routers.
type BaseRouter struct {
}

func GetUser(r *ghttp.Request) bean.SessionUser {
	resp := Token.GetTokenData(r)
	if !resp.Success() {
		return bean.SessionUser{}
	}

	var sessionUser bean.SessionUser
	err := gjson.DecodeTo(resp.GetString("data"), &sessionUser)
	if err != nil {
		glog.Error("get session user error", err)
	}

	return sessionUser
}

func Succ(r *ghttp.Request, data interface{}) {
	r.Response.WriteJson(resp.Succ(data))
	r.Exit()
}

func Fail(r *ghttp.Request, msg string) {
	r.Response.WriteJson(resp.Fail(msg))
	r.Exit()
}

func Error(r *ghttp.Request, msg string) {
	r.Response.WriteJson(resp.Error(msg))
	r.Exit()
}
func Resp(r *ghttp.Request, code int, msg string, data interface{}) {
	r.Response.WriteJson(resp.Resp{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	r.Exit()
}
