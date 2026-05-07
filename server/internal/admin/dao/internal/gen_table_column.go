// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GenTableColumnDao is the data access object for the table sys_gen_table_column.
type GenTableColumnDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  GenTableColumnColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// GenTableColumnColumns defines and stores column names for the table sys_gen_table_column.
type GenTableColumnColumns struct {
	Id            string // 编号
	TableId       string // 归属表编号
	ColumnName    string // 列名称
	ColumnComment string // 列描述
	ColumnType    string // 列类型（如 varchar(64)）
	GoType        string // Go 类型（string/int64/time.Time 等）
	GoField       string // Go 字段名（驼峰）
	IsPk          string // 是否主键（1=是）
	IsIncrement   string // 是否自增（1=是）
	IsRequired    string // 是否必填（1=是）
	IsInsert      string // 是否为插入字段（1=是）
	IsEdit        string // 是否编辑字段（1=是）
	IsList        string // 是否列表字段（1=是）
	IsQuery       string // 是否查询字段（1=是）
	QueryType     string // 查询方式（EQ/NE/GT/LT/LIKE/BETWEEN）
	HtmlType      string // 显示类型（input/textarea/select/radio/checkbox/datetime/switch）
	DictType      string // 字典类型（绑定 sys_config 的 config_key）
	Sort          string // 排序
	CreateBy      string //
	CreateAt      string //
	UpdateBy      string //
	UpdateAt      string //
}

// genTableColumnColumns holds the columns for the table sys_gen_table_column.
var genTableColumnColumns = GenTableColumnColumns{
	Id:            "id",
	TableId:       "table_id",
	ColumnName:    "column_name",
	ColumnComment: "column_comment",
	ColumnType:    "column_type",
	GoType:        "go_type",
	GoField:       "go_field",
	IsPk:          "is_pk",
	IsIncrement:   "is_increment",
	IsRequired:    "is_required",
	IsInsert:      "is_insert",
	IsEdit:        "is_edit",
	IsList:        "is_list",
	IsQuery:       "is_query",
	QueryType:     "query_type",
	HtmlType:      "html_type",
	DictType:      "dict_type",
	Sort:          "sort",
	CreateBy:      "create_by",
	CreateAt:      "create_at",
	UpdateBy:      "update_by",
	UpdateAt:      "update_at",
}

// NewGenTableColumnDao creates and returns a new DAO object for table data access.
func NewGenTableColumnDao(handlers ...gdb.ModelHandler) *GenTableColumnDao {
	return &GenTableColumnDao{
		group:    "default",
		table:    "sys_gen_table_column",
		columns:  genTableColumnColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *GenTableColumnDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *GenTableColumnDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *GenTableColumnDao) Columns() GenTableColumnColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *GenTableColumnDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *GenTableColumnDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *GenTableColumnDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
