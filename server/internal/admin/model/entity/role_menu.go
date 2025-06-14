// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RoleMenu is the golang structure for table role_menu.
type RoleMenu struct {
	RoleId int64 `json:"roleId" orm:"role_id" description:"角色id"` // 角色id
	MenuId int64 `json:"menuId" orm:"menu_id" description:"菜单id"` // 菜单id
}
