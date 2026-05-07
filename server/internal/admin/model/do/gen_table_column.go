// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GenTableColumn is the golang structure of table sys_gen_table_column for DAO operations like Where/Data.
type GenTableColumn struct {
	g.Meta        `orm:"table:sys_gen_table_column, do:true"`
	Id            any         // 编号
	TableId       any         // 归属表编号
	ColumnName    any         // 列名称
	ColumnComment any         // 列描述
	ColumnType    any         // 列类型（如 varchar(64)）
	GoType        any         // Go 类型（string/int64/time.Time 等）
	GoField       any         // Go 字段名（驼峰）
	IsPk          any         // 是否主键（1=是）
	IsIncrement   any         // 是否自增（1=是）
	IsRequired    any         // 是否必填（1=是）
	IsInsert      any         // 是否为插入字段（1=是）
	IsEdit        any         // 是否编辑字段（1=是）
	IsList        any         // 是否列表字段（1=是）
	IsQuery       any         // 是否查询字段（1=是）
	QueryType     any         // 查询方式（EQ/NE/GT/LT/LIKE/BETWEEN）
	HtmlType      any         // 显示类型（input/textarea/select/radio/checkbox/datetime/switch）
	DictType      any         // 字典类型（绑定 sys_config 的 config_key）
	Sort          any         // 排序
	CreateBy      any         //
	CreateAt      *gtime.Time //
	UpdateBy      any         //
	UpdateAt      *gtime.Time //
}
