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
	Id        any         // 主键
	ParentId  any         // 上级机构
	Name      any         // 部门/11111
	Code      any         // 机构编码
	Sort      any         // 序号
	Linkman   any         // 联系人
	LinkmanNo any         // 联系人电话
	Remark    any         // 机构描述
	Enable    any         // 是否启用//radio/1,启用,2,禁用
	UpdateAt  *gtime.Time // 更新时间
	UpdateId  any         // 更新人
	CreateAt  *gtime.Time // 创建时间
	CreateId  any         // 创建者
}
