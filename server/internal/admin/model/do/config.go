// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Config is the golang structure of table sys_config for DAO operations like Where/Data.
type Config struct {
	g.Meta       `orm:"table:sys_config, do:true"`
	Id           any         // 主键
	Name         any         // 名称
	Key          any         // 键
	Value        any         // 值
	Code         any         // 编码
	DataType     any         // 数据类型//radio/1,KV配置,2,字典,3,字典数据
	ParentId     any         // 类型
	ParentKey    any         //
	Remark       any         // 备注
	Sort         any         // 排序号
	CopyStatus   any         // 拷贝状态 1 拷贝  2  不拷贝
	ChangeStatus any         // 1 可以更改 2 不可更改
	Enable       any         // 是否启用//radio/1,启用,2,禁用
	UpdateAt     *gtime.Time // 更新时间
	UpdateId     any         // 更新人
	CreateAt     *gtime.Time // 创建时间
	CreateId     any         // 创建者
}
