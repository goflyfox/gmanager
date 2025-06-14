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
	Id         interface{} // ID
	ParentId   interface{} // 父菜单ID
	Name       interface{} // 菜单名称
	Type       interface{} // 菜单类型（1-菜单 2-目录 3-外链 4-按钮）
	RouteName  interface{} // 路由名称（Vue Router 中用于命名路由）
	RoutePath  interface{} // 路由路径（Vue Router 中定义的 URL 路径）
	Component  interface{} // 组件路径（组件页面完整路径，相对于 src/views/，缺省后缀 .vue）
	Perm       interface{} // 【按钮】权限标识
	AlwaysShow interface{} // 【目录】只有一个子路由是否始终显示（1-是 0-否）
	KeepAlive  interface{} // 【菜单】是否开启页面缓存（1-是 0-否）
	Sort       interface{} // 排序
	Icon       interface{} // 菜单图标
	Redirect   interface{} // 跳转路径
	Params     interface{} // 路由参数
	Enable     interface{} // 是否启用//radio/1,启用,2,禁用
	UpdateAt   *gtime.Time // 更新时间
	UpdateId   interface{} // 更新人
	CreateAt   *gtime.Time // 创建时间
	CreateId   interface{} // 创建者
}
