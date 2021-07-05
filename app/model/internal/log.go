// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

// Log is the golang structure for table sys_log.
type Log struct {
	Id         int    `orm:"id,primary"  json:"id"`         // 主键
	LogType    int    `orm:"log_type"    json:"logType"`    // 类型
	OperObject string `orm:"oper_object" json:"operObject"` // 操作对象
	OperTable  string `orm:"oper_table"  json:"operTable"`  // 操作表
	OperId     int    `orm:"oper_id"     json:"operId"`     // 操作主键
	OperType   string `orm:"oper_type"   json:"operType"`   // 操作类型
	OperRemark string `orm:"oper_remark" json:"operRemark"` // 操作备注
	Enable     int    `orm:"enable"      json:"enable"`     // 是否启用//radio/1,启用,2,禁用
	UpdateTime string `orm:"update_time" json:"updateTime"` // 更新时间
	UpdateId   int    `orm:"update_id"   json:"updateId"`   // 更新人
	CreateTime string `orm:"create_time" json:"createTime"` // 创建时间
	CreateId   int    `orm:"create_id"   json:"createId"`   // 创建者
}
