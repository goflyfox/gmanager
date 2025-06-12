// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeptDao is the data access object for the table sys_dept.
type DeptDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  DeptColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// DeptColumns defines and stores column names for the table sys_dept.
type DeptColumns struct {
	Id        string // 主键
	ParentId  string // 上级机构
	Name      string // 部门/11111
	Code      string // 机构编码
	Sort      string // 序号
	Linkman   string // 联系人
	LinkmanNo string // 联系人电话
	Remark    string // 机构描述
	Enable    string // 是否启用//radio/1,启用,2,禁用
	UpdateAt  string // 更新时间
	UpdateId  string // 更新人
	CreateAt  string // 创建时间
	CreateId  string // 创建者
}

// deptColumns holds the columns for the table sys_dept.
var deptColumns = DeptColumns{
	Id:        "id",
	ParentId:  "parent_id",
	Name:      "name",
	Code:      "code",
	Sort:      "sort",
	Linkman:   "linkman",
	LinkmanNo: "linkman_no",
	Remark:    "remark",
	Enable:    "enable",
	UpdateAt:  "update_at",
	UpdateId:  "update_id",
	CreateAt:  "create_at",
	CreateId:  "create_id",
}

// NewDeptDao creates and returns a new DAO object for table data access.
func NewDeptDao(handlers ...gdb.ModelHandler) *DeptDao {
	return &DeptDao{
		group:    "default",
		table:    "sys_dept",
		columns:  deptColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DeptDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DeptDao) Columns() DeptColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DeptDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DeptDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
