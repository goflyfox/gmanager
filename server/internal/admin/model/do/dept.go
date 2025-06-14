// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Dept is the golang structure of table sys_dept for DAO operations like Where/Data.
type Dept struct {
	g.Meta    `orm:"table:sys_dept, do:true"`
	Id        interface{} // 主键
	ParentId  interface{} // 上级机构
	Name      interface{} // 部门/11111
	Code      interface{} // 机构编码
	Sort      interface{} // 序号
	Linkman   interface{} // 联系人
	LinkmanNo interface{} // 联系人电话
	Remark    interface{} // 机构描述
	Enable    interface{} // 是否启用//radio/1,启用,2,禁用
	UpdateAt  *gtime.Time // 更新时间
	UpdateId  interface{} // 更新人
	CreateAt  *gtime.Time // 创建时间
	CreateId  interface{} // 创建者
}
