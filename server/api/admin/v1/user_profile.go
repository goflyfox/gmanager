package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"gmanager/internal/admin/model/input"
)

type UserProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" tags:"个人中心" summary:"个人中心信息获取"`
}

type UserProfileRes = input.UserProfile

type UserSaveProfileReq struct {
	g.Meta   `path:"/user/saveProfile" method:"post" tags:"个人中心" summary:"个人中心信息保存"`
	Id       int64  `json:"id"`
	NickName string `json:"nickName" dc:"昵称"`
	Mobile   string `json:"mobile"   dc:"手机号"`
	Email    string `json:"email"   dc:"邮箱"`
	Gender   int    `json:"gender" dc:"性别"`
	Address  string `json:"address" dc:"地址"`
	Avatar   string `json:"avatar" dc:"头像地址"`
}

type UserSaveProfileRes struct {
}

type UserChangePasswordReq struct {
	g.Meta      `path:"/user/changePassword" method:"post" tags:"个人中心" summary:"个人中心修改密码"`
	OldPassword string `json:"oldPassword" dc:"原密码" v:"required#原密码不能为空"`
	NewPassword string `json:"newPassword" dc:"新密码" v:"required#新密码不能为空"`
}

type UserChangePasswordRes struct {
}

type UserSendMobileCodeReq struct {
	g.Meta `path:"/user/sendMobileCode" method:"post" tags:"个人中心" summary:"个人中心发送手机验证码"`
	Mobile string `json:"mobile" dc:"手机号" v:"required#手机号不能为空"`
}

type UserSendMobileCodeRes struct {
}

type UserSaveMobileReq struct {
	g.Meta `path:"/user/saveMobile" method:"post" tags:"个人中心" summary:"个人中心修改手机号"`
	Mobile string `json:"mobile" dc:"手机号" v:"required#手机号不能为空"`
	Code   string `json:"code" dc:"验证码" v:"required#验证码不能为空"`
}

type UserSaveMobileRes struct {
}

type UserSendEmailCodeReq struct {
	g.Meta `path:"/user/sendEmailCode" method:"post" tags:"个人中心" summary:"个人中心发送Email验证码"`
	Email  string `json:"mobile" dc:"邮箱" v:"required#邮箱不能为空"`
}

type UserSendEmailCodeRes struct {
}

type UserSaveEmailReq struct {
	g.Meta `path:"/user/saveEmail" method:"post" tags:"个人中心" summary:"个人中心修改Email"`
	Email  string `json:"mobile" dc:"邮箱" v:"required#邮箱不能为空"`
	Code   string `json:"code" dc:"验证码" v:"required#验证码不能为空"`
}

type UserSaveEmailRes struct {
}
