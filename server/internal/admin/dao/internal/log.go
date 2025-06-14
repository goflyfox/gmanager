// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LogDao is the data access object for the table sys_log.
type LogDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  LogColumns         // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// LogColumns defines and stores column names for the table sys_log.
type LogColumns struct {
	Id            string // 主键
	LogType       string // 类型
	OperObject    string // 操作对象
	OperTable     string // 操作表
	OperId        string // 操作主键
	OperType      string // 操作类型
	OperRemark    string // 操作备注
	Url           string // 提交url
	Method        string // 请求方式
	Ip            string // IP地址
	UserAgent     string // UA信息
	ExecutionTime string // 响应时间
	Operator      string // 操作人
	Enable        string // 是否启用//radio/1,启用,2,禁用
	UpdateAt      string // 更新时间
	UpdateId      string // 更新人
	CreateAt      string // 创建时间
	CreateId      string // 创建者
}

// logColumns holds the columns for the table sys_log.
var logColumns = LogColumns{
	Id:            "id",
	LogType:       "log_type",
	OperObject:    "oper_object",
	OperTable:     "oper_table",
	OperId:        "oper_id",
	OperType:      "oper_type",
	OperRemark:    "oper_remark",
	Url:           "url",
	Method:        "method",
	Ip:            "ip",
	UserAgent:     "user_agent",
	ExecutionTime: "execution_time",
	Operator:      "operator",
	Enable:        "enable",
	UpdateAt:      "update_at",
	UpdateId:      "update_id",
	CreateAt:      "create_at",
	CreateId:      "create_id",
}

// NewLogDao creates and returns a new DAO object for table data access.
func NewLogDao(handlers ...gdb.ModelHandler) *LogDao {
	return &LogDao{
		group:    "default",
		table:    "sys_log",
		columns:  logColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *LogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *LogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *LogDao) Columns() LogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *LogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *LogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *LogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
