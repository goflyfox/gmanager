package common

import (
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"gmanager/app/constants"
	"gmanager/app/service/log"
	"gmanager/app/service/user"
	"gmanager/library"
	"gmanager/library/base"
	"gmanager/library/bean"
)

// Login 登录页面
func Login(r *ghttp.Request) {
	err := r.Response.WriteTpl("pages/login.html", g.Map{})

	if err != nil {
		glog.Error(err)
	}
}

func Index(r *ghttp.Request) {
	err := r.Response.WriteTpl("pages/home.html", g.Map{
		"id":    1,
		"name":  "flyfox",
		"title": g.Config().GetString("setting.title"),
	})

	if err != nil {
		glog.Error(err)
	}

}

// LoginSubmit 登录认证
func LoginSubmit(r *ghttp.Request) (string, interface{}) {
	username := r.GetString("username")
	passwd := r.GetString("passwd")
	if username == "" || passwd == "" {
		base.Fail(r, "用户名或密码为空")
	}

	model, err := user.GetByUsername(username)
	if err != nil {
		base.Error(r, "服务异常，请联系管理员")
	}

	if model == nil || model.Id <= 0 {
		base.Fail(r, "用户名或密码错误.")
	}

	if model.Enable != constants.EnableYes {
		base.Fail(r, "账号状态异常，请联系管理员")
	}

	reqPassword, err2 := gmd5.Encrypt(passwd + model.Salt)
	if err2 != nil {
		glog.Error("login password encrypt error", err2)
		base.Error(r, "用户名或者密码错误："+username)
	}

	if reqPassword != model.Password {
		base.Fail(r, "用户名或者密码错误!")
	}

	sessionUser := bean.SessionUser{
		Id:       model.Id,
		Uuid:     model.Uuid,
		RealName: model.RealName,
		Username: model.Username,
	}

	// 登录日志
	model.UpdateTime = library.GetNow()
	model.UpdateId = model.Id
	log.SaveLog(model, constants.LOGIN)

	return username, sessionUser
}

// 登出
func LogoutBefore(r *ghttp.Request) bool {
	userId := base.GetUser(r).Id
	model, err := user.GetById(gconv.Int64(userId))
	if err != nil {
		glog.Warning("logout getUser error", err)
		return false
	} else if model.Id != userId {
		// 登出用户不存在
		glog.Warning("logout userId error", userId)
		return false
	}

	// 登出日志
	model.UpdateTime = library.GetNow()
	model.UpdateId = userId
	log.SaveLog(model, constants.LOGOUT)
	return true
}
