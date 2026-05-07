// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GenTableColumn is the golang structure for table gen_table_column.
type GenTableColumn struct {
	Id            int64       `json:"id"            orm:"id"             description:"编号"`                                                         // 编号
	TableId       int64       `json:"tableId"       orm:"table_id"       description:"归属表编号"`                                                      // 归属表编号
	ColumnName    string      `json:"columnName"    orm:"column_name"    description:"列名称"`                                                        // 列名称
	ColumnComment string      `json:"columnComment" orm:"column_comment" description:"列描述"`                                                        // 列描述
	ColumnType    string      `json:"columnType"    orm:"column_type"    description:"列类型（如 varchar(64)）"`                                         // 列类型（如 varchar(64)）
	GoType        string      `json:"goType"        orm:"go_type"        description:"Go 类型（string/int64/time.Time 等）"`                            // Go 类型（string/int64/time.Time 等）
	GoField       string      `json:"goField"       orm:"go_field"       description:"Go 字段名（驼峰）"`                                                 // Go 字段名（驼峰）
	IsPk          string      `json:"isPk"          orm:"is_pk"          description:"是否主键（1=是）"`                                                  // 是否主键（1=是）
	IsIncrement   string      `json:"isIncrement"   orm:"is_increment"   description:"是否自增（1=是）"`                                                  // 是否自增（1=是）
	IsRequired    string      `json:"isRequired"    orm:"is_required"    description:"是否必填（1=是）"`                                                  // 是否必填（1=是）
	IsInsert      string      `json:"isInsert"      orm:"is_insert"      description:"是否为插入字段（1=是）"`                                               // 是否为插入字段（1=是）
	IsEdit        string      `json:"isEdit"        orm:"is_edit"        description:"是否编辑字段（1=是）"`                                                // 是否编辑字段（1=是）
	IsList        string      `json:"isList"        orm:"is_list"        description:"是否列表字段（1=是）"`                                                // 是否列表字段（1=是）
	IsQuery       string      `json:"isQuery"       orm:"is_query"       description:"是否查询字段（1=是）"`                                                // 是否查询字段（1=是）
	QueryType     string      `json:"queryType"     orm:"query_type"     description:"查询方式（EQ/NE/GT/LT/LIKE/BETWEEN）"`                             // 查询方式（EQ/NE/GT/LT/LIKE/BETWEEN）
	HtmlType      string      `json:"htmlType"      orm:"html_type"      description:"显示类型（input/textarea/select/radio/checkbox/datetime/switch）"` // 显示类型（input/textarea/select/radio/checkbox/datetime/switch）
	DictType      string      `json:"dictType"      orm:"dict_type"      description:"字典类型（绑定 sys_config 的 config_key）"`                           // 字典类型（绑定 sys_config 的 config_key）
	Sort          int         `json:"sort"          orm:"sort"           description:"排序"`                                                         // 排序
	CreateBy      int64       `json:"createBy"      orm:"create_by"      description:""`                                                           //
	CreateAt      *gtime.Time `json:"createAt"      orm:"create_at"      description:""`                                                           //
	UpdateBy      int64       `json:"updateBy"      orm:"update_by"      description:""`                                                           //
	UpdateAt      *gtime.Time `json:"updateAt"      orm:"update_at"      description:""`                                                           //
}
