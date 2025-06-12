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
	Id       interface{} // 主键
	Uuid     interface{} // UUID
	UserName interface{} // 登录名/11111
	Mobile   interface{} // 手机号
	Email    interface{} // email
	Password interface{} // 密码
	Salt     interface{} // 密码盐
	DeptId   interface{} // 部门/11111/dict
	UserType interface{} // 类型//select/1,管理员,2,普通用户,3,前台用户,4,第三方用户,5,API用户
	Status   interface{} // 状态
	Thirdid  interface{} // 第三方ID
	Endtime  interface{} // 结束时间
	NickName interface{} // 昵称
	Gender   interface{} // 性别;0:保密,1:男,2:女
	Address  interface{} // 地址
	Avatar   interface{} // 头像地址
	Birthday interface{} // 生日
	Remark   interface{} // 说明
	Enable   interface{} // 是否启用//radio/1,启用,2,禁用
	UpdateAt *gtime.Time // 更新时间
	UpdateId interface{} // 更新人
	CreateAt *gtime.Time // 创建时间
	CreateId interface{} // 创建者
}
