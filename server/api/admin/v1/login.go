package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" tags:"用户登陆" perms:"user:login" method:"post" summary:"用户登陆"`
	Username string `json:"username" v:"required#请输入用户名" dc:"用户名"`
	Password string `json:"password" v:"required#请输入密码" dc:"密码"`
	Code     string `json:"code" dc:"验证码"`
	CodeId   string `json:"codeId" dc:"验证码ID"`
}

type LoginRes struct {
	Username    string `json:"username"`
	TokenType   string `json:"tokenType"`
	AccessToken string `json:"accessToken"`
}

type LogoutReq struct {
	g.Meta `path:"/user/logout" tags:"用户登陆" method:"post" summary:"用户登出"`
}
type LogoutRes struct {
}

type CaptchaReq struct {
	g.Meta `path:"/captcha/get" tags:"用户登陆" method:"get" summary:"获取验证码"`
}
type CaptchaRes struct {
	g.Meta `mime:"application/json"`
	CodeId string `json:"codeId"`
	Img    string `json:"img"`
}
