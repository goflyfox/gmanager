// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gmanager/app/dao/internal"
)

// userRoleDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type userRoleDao struct {
	*internal.UserRoleDao
}

var (
	// UserRole is globally public accessible object for table sys_user_role operations.
	UserRole = userRoleDao{
		internal.NewUserRoleDao(),
	}
)

// Fill with you ideas below.
