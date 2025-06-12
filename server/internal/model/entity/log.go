// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Log is the golang structure for table log.
type Log struct {
	Id            int64       `json:"id"            orm:"id"             description:"主键"`                    // 主键
	LogType       int         `json:"logType"       orm:"log_type"       description:"类型"`                    // 类型
	OperObject    string      `json:"operObject"    orm:"oper_object"    description:"操作对象"`                  // 操作对象
	OperTable     string      `json:"operTable"     orm:"oper_table"     description:"操作表"`                   // 操作表
	OperId        int64       `json:"operId"        orm:"oper_id"        description:"操作主键"`                  // 操作主键
	OperType      string      `json:"operType"      orm:"oper_type"      description:"操作类型"`                  // 操作类型
	OperRemark    string      `json:"operRemark"    orm:"oper_remark"    description:"操作备注"`                  // 操作备注
	Url           string      `json:"url"           orm:"url"            description:"提交url"`                 // 提交url
	Method        string      `json:"method"        orm:"method"         description:"请求方式"`                  // 请求方式
	Ip            string      `json:"ip"            orm:"ip"             description:"IP地址"`                  // IP地址
	UserAgent     string      `json:"userAgent"     orm:"user_agent"     description:"UA信息"`                  // UA信息
	ExecutionTime int64       `json:"executionTime" orm:"execution_time" description:"响应时间"`                  // 响应时间
	Operator      string      `json:"operator"      orm:"operator"       description:"操作人"`                   // 操作人
	Enable        int         `json:"enable"        orm:"enable"         description:"是否启用//radio/1,启用,2,禁用"` // 是否启用//radio/1,启用,2,禁用
	UpdateAt      *gtime.Time `json:"updateAt"      orm:"update_at"      description:"更新时间"`                  // 更新时间
	UpdateId      int64       `json:"updateId"      orm:"update_id"      description:"更新人"`                   // 更新人
	CreateAt      *gtime.Time `json:"createAt"      orm:"create_at"      description:"创建时间"`                  // 创建时间
	CreateId      int64       `json:"createId"      orm:"create_id"      description:"创建者"`                   // 创建者
}
