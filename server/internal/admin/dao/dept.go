// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"gmanager/internal/admin/dao/internal"
)

// deptDao is the data access object for the table sys_dept.
// You can define custom methods on it to extend its functionality as needed.
type deptDao struct {
	*internal.DeptDao
}

var (
	// Dept is a globally accessible object for table sys_dept operations.
	Dept = deptDao{internal.NewDeptDao()}
)

// Add your custom methods and functionality below.
