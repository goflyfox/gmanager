// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

// UserRole is the golang structure for table sys_user_role.
type UserRole struct {
	Id     int `orm:"id,primary" json:"id"`     // 主键
	UserId int `orm:"user_id"    json:"userId"` // 用户id
	RoleId int `orm:"role_id"    json:"roleId"` // 角色id
}
