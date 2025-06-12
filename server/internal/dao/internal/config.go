// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ConfigDao is the data access object for the table sys_config.
type ConfigDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ConfigColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ConfigColumns defines and stores column names for the table sys_config.
type ConfigColumns struct {
	Id           string // 主键
	Name         string // 名称
	Key          string // 键
	Value        string // 值
	Code         string // 编码
	DataType     string // 数据类型//radio/1,KV配置,2,字典,3,字典数据
	ParentId     string // 类型
	ParentKey    string //
	Remark       string // 备注
	Sort         string // 排序号
	CopyStatus   string // 拷贝状态 1 拷贝  2  不拷贝
	ChangeStatus string // 1 可以更改 2 不可更改
	Enable       string // 是否启用//radio/1,启用,2,禁用
	UpdateAt     string // 更新时间
	UpdateId     string // 更新人
	CreateAt     string // 创建时间
	CreateId     string // 创建者
}

// configColumns holds the columns for the table sys_config.
var configColumns = ConfigColumns{
	Id:           "id",
	Name:         "name",
	Key:          "key",
	Value:        "value",
	Code:         "code",
	DataType:     "data_type",
	ParentId:     "parent_id",
	ParentKey:    "parent_key",
	Remark:       "remark",
	Sort:         "sort",
	CopyStatus:   "copy_status",
	ChangeStatus: "change_status",
	Enable:       "enable",
	UpdateAt:     "update_at",
	UpdateId:     "update_id",
	CreateAt:     "create_at",
	CreateId:     "create_id",
}

// NewConfigDao creates and returns a new DAO object for table data access.
func NewConfigDao(handlers ...gdb.ModelHandler) *ConfigDao {
	return &ConfigDao{
		group:    "default",
		table:    "sys_config",
		columns:  configColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ConfigDao) Columns() ConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ConfigDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
