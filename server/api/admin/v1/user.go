package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	input2 "gmanager/internal/admin/model/input"
)

type UserListReq struct {
	g.Meta   `path:"/user/list" method:"post" tags:"用户管理" summary:"用户列表"`
	Keywords string      `json:"keywords" dc:"用户/手机号/昵称名称"`
	DeptId   int64       `json:"deptId"  dc:"部门"`
	Status   int         `json:"code"  dc:"用户状态"`
	Enable   int         `json:"enable" dc:"是否启用"`
	CreateAt g.MapIntStr `json:"createAt" dc:"创建时间区间"`
	input2.PageReq
}

type UserListRes struct {
	List []*input2.User `json:"list" dc:"用户列表"`
	input2.PageRes
}

type UserSaveReq struct {
	g.Meta   `path:"/user/save/:id" method:"post" tags:"用户管理" summary:"用户保存"`
	Id       int64   `json:"id"`
	DeptId   int64   `json:"deptId"  dc:"部门id" v:"required#部门不能为空"`
	UserName string  `json:"userName"  dc:"用户名称" v:"required#用户名称不能为空"`
	NickName string  `json:"nickName" dc:"昵称" v:"required#昵称不能为空"`
	Mobile   string  `json:"mobile"   dc:"手机号"`
	Email    string  `json:"email"   dc:"邮箱"`
	Status   int     `json:"status" dc:"状态"`
	Endtime  string  `json:"endtime" dc:"结束时间"`
	Gender   int     `json:"gender" dc:"性别"`
	Address  string  `json:"address" dc:"地址"`
	Avatar   string  `json:"avatar" dc:"头像地址"`
	Birthday int     `json:"birthday" dc:"生日"`
	Remark   string  `json:"remark" dc:"备注"`
	Enable   int     `json:"enable" dc:"是否启用"`
	RoleIds  []int64 `json:"roleIds" dc:"角色id"`
}

type UserSaveRes struct {
}

type UserGetReq struct {
	g.Meta `path:"/user/get/:id" method:"get" tags:"用户管理" summary:"用户获取"`
	Id     int64 `json:"id" dc:"ID"`
}

type UserGetRes = input2.User

type UserDeleteReq struct {
	g.Meta `path:"/user/delete/:ids" method:"post" tags:"用户管理" summary:"用户删除"`
	Ids    string `json:"ids" dc:"删除id列表"`
}

type UserDeleteRes struct {
}

type UserPasswordResetReq struct {
	g.Meta   `path:"/user/password/reset/:id" method:"post" tags:"用户管理" summary:"用户密码重置"`
	Id       int64  `json:"id"`
	Password string `json:"password" dc:"密码" v:"required#密码不能为空"`
}

type UserPasswordResetRes struct {
}

type UserInfoReq struct {
	g.Meta `path:"/user/info" tags:"用户管理" method:"get" summary:"用户登陆信息"`
}
type UserInfoRes struct {
	UserID   string   `json:"userId"`
	Username string   `json:"username"`
	Nickname string   `json:"nickname"`
	Avatar   string   `json:"avatar"`
	Roles    []string `json:"roles"`
	Perms    []string `json:"perms"`
}

type UserMenusReq struct {
	g.Meta `path:"/user/menus" tags:"用户管理" method:"get" summary:"用户菜单信息"`
}

type UserMenusRes = []*input2.UserMenu

type UserExportReq struct {
	g.Meta   `path:"/user/export" method:"get" tags:"用户管理" summary:"用户数据导出"`
	Keywords string      `json:"keywords" dc:"用户/手机号/昵称名称"`
	DeptId   int64       `json:"deptId"  dc:"部门"`
	Status   int         `json:"code"  dc:"用户状态"`
	Enable   int         `json:"enable" dc:"是否启用"`
	CreateAt g.MapIntStr `json:"createAt" dc:"创建时间区间"`
	input2.PageReq
}

type UserExportRes struct {
}

type UserImportReq struct {
	g.Meta `path:"/user/import" method:"post" tags:"用户管理" summary:"批量导入用户"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"分片文件"`
}

type UserImportRes struct {
	Code         int      `json:"code"   dc:"状态码"`         // 状态码
	InvalidCount int      `json:"invalidCount"   dc:"手机号"` // 无效数据条数
	ValidCount   int      `json:"validCount"   dc:"手机号"`   // 有效数据条数
	MessageList  []string `json:"messageList"   dc:"手机号"`  // 错误信息
}

type UserTemplateReq struct {
	g.Meta `path:"/user/template" method:"get" tags:"用户管理" summary:"批量创建用户模版下载"`
}

type UserTemplateRes struct {
}
