// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table sys_role for DAO operations like Where/Data.
type Role struct {
	g.Meta    `orm:"table:sys_role, do:true"`
	Id        any         // 主键
	Name      any         // 名称/11111/
	Code      any         // 角色编码
	DataScope any         // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Sort      any         // 排序
	Remark    any         // 说明//textarea
	Enable    any         // 是否启用//radio/1,启用,2,禁用
	UpdateAt  *gtime.Time // 更新时间
	UpdateId  any         // 更新人
	CreateAt  *gtime.Time // 创建时间
	CreateId  any         // 创建者
}
