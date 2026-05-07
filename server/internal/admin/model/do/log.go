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
	Id            any         // 主键
	LogType       any         // 类型
	OperObject    any         // 操作对象
	OperTable     any         // 操作表
	OperId        any         // 操作主键
	OperType      any         // 操作类型
	OperRemark    any         // 操作备注
	Url           any         // 提交url
	Method        any         // 请求方式
	Ip            any         // IP地址
	UserAgent     any         // UA信息
	ExecutionTime any         // 响应时间
	Operator      any         // 操作人
	Enable        any         // 是否启用//radio/1,启用,2,禁用
	UpdateAt      *gtime.Time // 更新时间
	UpdateId      any         // 更新人
	CreateAt      *gtime.Time // 创建时间
	CreateId      any         // 创建者
}
