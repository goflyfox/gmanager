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
	Id           interface{} // 主键
	Name         interface{} // 名称
	Key          interface{} // 键
	Value        interface{} // 值
	Code         interface{} // 编码
	DataType     interface{} // 数据类型//radio/1,KV配置,2,字典,3,字典数据
	ParentId     interface{} // 类型
	ParentKey    interface{} //
	Remark       interface{} // 备注
	Sort         interface{} // 排序号
	CopyStatus   interface{} // 拷贝状态 1 拷贝  2  不拷贝
	ChangeStatus interface{} // 1 可以更改 2 不可更改
	Enable       interface{} // 是否启用//radio/1,启用,2,禁用
	UpdateAt     *gtime.Time // 更新时间
	UpdateId     interface{} // 更新人
	CreateAt     *gtime.Time // 创建时间
	CreateId     interface{} // 创建者
}
