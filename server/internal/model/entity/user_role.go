// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserRole is the golang structure for table user_role.
type UserRole struct {
	UserId int64 `json:"userId" orm:"user_id" description:"用户id"` // 用户id
	RoleId int64 `json:"roleId" orm:"role_id" description:"角色id"` // 角色id
}
