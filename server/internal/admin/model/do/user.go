// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table sys_user for DAO operations like Where/Data.
type User struct {
	g.Meta   `orm:"table:sys_user, do:true"`
	Id       any         // 主键
	Uuid     any         // UUID
	UserName any         // 登录名/11111
	Mobile   any         // 手机号
	Email    any         // email
	Password any         // 密码
	Salt     any         // 密码盐
	DeptId   any         // 部门/11111/dict
	UserType any         // 类型//select/1,管理员,2,普通用户,3,前台用户,4,第三方用户,5,API用户
	Status   any         // 状态
	Thirdid  any         // 第三方ID
	Endtime  any         // 结束时间
	NickName any         // 昵称
	Gender   any         // 性别;0:保密,1:男,2:女
	Address  any         // 地址
	Avatar   any         // 头像地址
	Birthday any         // 生日
	Remark   any         // 说明
	Enable   any         // 是否启用//radio/1,启用,2,禁用
	UpdateAt *gtime.Time // 更新时间
	UpdateId any         // 更新人
	CreateAt *gtime.Time // 创建时间
	CreateId any         // 创建者
}
