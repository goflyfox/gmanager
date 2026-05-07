// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GenTable is the golang structure for table gen_table.
type GenTable struct {
	Id             int64       `json:"id"             orm:"id"              description:"编号"`                           // 编号
	TableName      string      `json:"tableName"      orm:"table_name"      description:"表名称"`                          // 表名称
	TableComment   string      `json:"tableComment"   orm:"table_comment"   description:"表描述"`                          // 表描述
	ClassName      string      `json:"className"      orm:"class_name"      description:"实体类名称（首字母大写）"`                 // 实体类名称（首字母大写）
	PackageName    string      `json:"packageName"    orm:"package_name"    description:"生成包路径"`                        // 生成包路径
	ModuleName     string      `json:"moduleName"     orm:"module_name"     description:"生成模块名（如 system）"`              // 生成模块名（如 system）
	BusinessName   string      `json:"businessName"   orm:"business_name"   description:"生成业务名（如 post）"`                // 生成业务名（如 post）
	FunctionName   string      `json:"functionName"   orm:"function_name"   description:"生成功能名（如 岗位管理）"`                // 生成功能名（如 岗位管理）
	FunctionAuthor string      `json:"functionAuthor" orm:"function_author" description:"生成作者"`                         // 生成作者
	TplCategory    string      `json:"tplCategory"    orm:"tpl_category"    description:"模板类型（crud/tree/sub，一期仅 crud）"` // 模板类型（crud/tree/sub，一期仅 crud）
	GenType        string      `json:"genType"        orm:"gen_type"        description:"生成方式（0=ZIP压缩包 1=自定义路径）"`       // 生成方式（0=ZIP压缩包 1=自定义路径）
	GenPath        string      `json:"genPath"        orm:"gen_path"        description:"自定义生成路径"`                      // 自定义生成路径
	Options        string      `json:"options"        orm:"options"         description:"其它生成选项（JSON）"`                 // 其它生成选项（JSON）
	CreateBy       int64       `json:"createBy"       orm:"create_by"       description:"创建人"`                          // 创建人
	CreateAt       *gtime.Time `json:"createAt"       orm:"create_at"       description:"创建时间"`                         // 创建时间
	UpdateBy       int64       `json:"updateBy"       orm:"update_by"       description:"更新人"`                          // 更新人
	UpdateAt       *gtime.Time `json:"updateAt"       orm:"update_at"       description:"更新时间"`                         // 更新时间
}
