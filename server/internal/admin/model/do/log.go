// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Log is the golang structure of table sys_log for DAO operations like Where/Data.
type Log struct {
	g.Meta        `orm:"table:sys_log, do:true"`
	Id            interface{} // 主键
	LogType       interface{} // 类型
	OperObject    interface{} // 操作对象
	OperTable     interface{} // 操作表
	OperId        interface{} // 操作主键
	OperType      interface{} // 操作类型
	OperRemark    interface{} // 操作备注
	Url           interface{} // 提交url
	Method        interface{} // 请求方式
	Ip            interface{} // IP地址
	UserAgent     interface{} // UA信息
	ExecutionTime interface{} // 响应时间
	Operator      interface{} // 操作人
	Enable        interface{} // 是否启用//radio/1,启用,2,禁用
	UpdateAt      *gtime.Time // 更新时间
	UpdateId      interface{} // 更新人
	CreateAt      *gtime.Time // 创建时间
	CreateId      interface{} // 创建者
}
