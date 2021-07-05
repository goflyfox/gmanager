// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gmanager/app/dao/internal"
)

// userDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type userDao struct {
	*internal.UserDao
}

var (
	// User is globally public accessible object for table sys_user operations.
	User userDao
)

func init() {
	User = userDao{
		internal.NewUserDao(),
	}
}

// Fill with you ideas below.