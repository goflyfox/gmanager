// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GenTableDao is the data access object for the table sys_gen_table.
type GenTableDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  GenTableColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// GenTableColumns defines and stores column names for the table sys_gen_table.
type GenTableColumns struct {
	Id             string // 编号
	TableName      string // 表名称
	TableComment   string // 表描述
	ClassName      string // 实体类名称（首字母大写）
	PackageName    string // 生成包路径
	ModuleName     string // 生成模块名（如 system）
	BusinessName   string // 生成业务名（如 post）
	FunctionName   string // 生成功能名（如 岗位管理）
	FunctionAuthor string // 生成作者
	TplCategory    string // 模板类型（crud/tree/sub，一期仅 crud）
	GenType        string // 生成方式（0=ZIP压缩包 1=自定义路径）
	GenPath        string // 自定义生成路径
	Options        string // 其它生成选项（JSON）
	CreateBy       string // 创建人
	CreateAt       string // 创建时间
	UpdateBy       string // 更新人
	UpdateAt       string // 更新时间
}

// genTableColumns holds the columns for the table sys_gen_table.
var genTableColumns = GenTableColumns{
	Id:             "id",
	TableName:      "table_name",
	TableComment:   "table_comment",
	ClassName:      "class_name",
	PackageName:    "package_name",
	ModuleName:     "module_name",
	BusinessName:   "business_name",
	FunctionName:   "function_name",
	FunctionAuthor: "function_author",
	TplCategory:    "tpl_category",
	GenType:        "gen_type",
	GenPath:        "gen_path",
	Options:        "options",
	CreateBy:       "create_by",
	CreateAt:       "create_at",
	UpdateBy:       "update_by",
	UpdateAt:       "update_at",
}

// NewGenTableDao creates and returns a new DAO object for table data access.
func NewGenTableDao(handlers ...gdb.ModelHandler) *GenTableDao {
	return &GenTableDao{
		group:    "default",
		table:    "sys_gen_table",
		columns:  genTableColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *GenTableDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *GenTableDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *GenTableDao) Columns() GenTableColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *GenTableDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *GenTableDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *GenTableDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
