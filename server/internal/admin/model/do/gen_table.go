// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GenTable is the golang structure of table sys_gen_table for DAO operations like Where/Data.
type GenTable struct {
	g.Meta         `orm:"table:sys_gen_table, do:true"`
	Id             any         // 编号
	TableName      any         // 表名称
	TableComment   any         // 表描述
	ClassName      any         // 实体类名称（首字母大写）
	PackageName    any         // 生成包路径
	ModuleName     any         // 生成模块名（如 system）
	BusinessName   any         // 生成业务名（如 post）
	FunctionName   any         // 生成功能名（如 岗位管理）
	FunctionAuthor any         // 生成作者
	TplCategory    any         // 模板类型（crud/tree/sub，一期仅 crud）
	GenType        any         // 生成方式（0=ZIP压缩包 1=自定义路径）
	GenPath        any         // 自定义生成路径
	Options        any         // 其它生成选项（JSON）
	CreateBy       any         // 创建人
	CreateAt       *gtime.Time // 创建时间
	UpdateBy       any         // 更新人
	UpdateAt       *gtime.Time // 更新时间
}
