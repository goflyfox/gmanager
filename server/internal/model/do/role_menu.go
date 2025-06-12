// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoleMenu is the golang structure of table sys_role_menu for DAO operations like Where/Data.
type RoleMenu struct {
	g.Meta `orm:"table:sys_role_menu, do:true"`
	RoleId interface{} // 角色id
	MenuId interface{} // 菜单id
}
