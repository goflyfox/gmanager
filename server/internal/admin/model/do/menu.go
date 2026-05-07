// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure of table sys_menu for DAO operations like Where/Data.
type Menu struct {
	g.Meta     `orm:"table:sys_menu, do:true"`
	Id         any         // ID
	ParentId   any         // 父菜单ID
	Name       any         // 菜单名称
	Type       any         // 菜单类型（1-菜单 2-目录 3-外链 4-按钮）
	RouteName  any         // 路由名称（Vue Router 中用于命名路由）
	RoutePath  any         // 路由路径（Vue Router 中定义的 URL 路径）
	Component  any         // 组件路径（组件页面完整路径，相对于 src/views/，缺省后缀 .vue）
	Perm       any         // 【按钮】权限标识
	AlwaysShow any         // 【目录】只有一个子路由是否始终显示（1-是 0-否）
	KeepAlive  any         // 【菜单】是否开启页面缓存（1-是 0-否）
	Sort       any         // 排序
	Icon       any         // 菜单图标
	Redirect   any         // 跳转路径
	Params     any         // 路由参数
	Enable     any         // 是否启用//radio/1,启用,2,禁用
	UpdateAt   *gtime.Time // 更新时间
	UpdateId   any         // 更新人
	CreateAt   *gtime.Time // 创建时间
	CreateId   any         // 创建者
}
