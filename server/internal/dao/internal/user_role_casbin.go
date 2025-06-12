// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserRoleCasbinDao is the data access object for the table sys_user_role_casbin.
type UserRoleCasbinDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  UserRoleCasbinColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// UserRoleCasbinColumns defines and stores column names for the table sys_user_role_casbin.
type UserRoleCasbinColumns struct {
	Id    string //
	PType string //
	V0    string //
	V1    string //
	V2    string //
	V3    string //
	V4    string //
	V5    string //
}

// userRoleCasbinColumns holds the columns for the table sys_user_role_casbin.
var userRoleCasbinColumns = UserRoleCasbinColumns{
	Id:    "id",
	PType: "p_type",
	V0:    "v0",
	V1:    "v1",
	V2:    "v2",
	V3:    "v3",
	V4:    "v4",
	V5:    "v5",
}

// NewUserRoleCasbinDao creates and returns a new DAO object for table data access.
func NewUserRoleCasbinDao(handlers ...gdb.ModelHandler) *UserRoleCasbinDao {
	return &UserRoleCasbinDao{
		group:    "default",
		table:    "sys_user_role_casbin",
		columns:  userRoleCasbinColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserRoleCasbinDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserRoleCasbinDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserRoleCasbinDao) Columns() UserRoleCasbinColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserRoleCasbinDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserRoleCasbinDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UserRoleCasbinDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
