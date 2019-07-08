package common

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/crypto/gmd5"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"gmanager/module/constants"
	"gmanager/module/system"
	"gmanager/utils"
	"gmanager/utils/base"
	"gmanager/utils/bean"
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

	model := system.SysUser{Username: username}.GetByUsername()

	if model.Id <= 0 {
		base.Fail(r, "用户名或密码错误："+username)
	}

	if model.Enable != constants.EnableYes {
		base.Fail(r, "账号状态异常，请联系管理员")
	}

	reqPassword, err := gmd5.Encrypt(passwd + model.Salt)
	if err != nil {
		glog.Error("login password encrypt error", err)
		base.Error(r, "login password encrypt error")
	}

	if reqPassword != model.Password {
		base.Fail(r, "用户名或者密码错误："+username)
	}

	sessionUser := bean.SessionUser{
		Id:       model.Id,
		Uuid:     model.Uuid,
		RealName: model.RealName,
		Username: model.Username,
	}

	// 登录日志
	model.UpdateTime = utils.GetNow()
	model.UpdateId = model.Id
	system.LogSave(model, system.LOGIN)

	return username, sessionUser
}

// 登出
func LogoutBefore(r *ghttp.Request) bool {
	userId := base.GetUser(r).Id
	model := system.SysUser{Id: userId}.Get()
	if model.Id != userId {
		// 登出用户不存在
		glog.Warning("logout userId error", userId)
		return false
	}

	// 登出日志
	model.UpdateTime = utils.GetNow()
	model.UpdateId = userId
	system.LogSave(model, system.LOGOUT)
	return true
}
