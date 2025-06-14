// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserRole is the golang structure of table sys_user_role for DAO operations like Where/Data.
type UserRole struct {
	g.Meta `orm:"table:sys_user_role, do:true"`
	UserId interface{} // 用户id
	RoleId interface{} // 角色id
}
