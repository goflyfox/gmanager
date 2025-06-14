// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Menu is the golang structure for table menu.
type Menu struct {
	Id         int64       `json:"id"         orm:"id"          description:"ID"`                                      // ID
	ParentId   int64       `json:"parentId"   orm:"parent_id"   description:"父菜单ID"`                                   // 父菜单ID
	Name       string      `json:"name"       orm:"name"        description:"菜单名称"`                                    // 菜单名称
	Type       int         `json:"type"       orm:"type"        description:"菜单类型（1-菜单 2-目录 3-外链 4-按钮）"`               // 菜单类型（1-菜单 2-目录 3-外链 4-按钮）
	RouteName  string      `json:"routeName"  orm:"route_name"  description:"路由名称（Vue Router 中用于命名路由）"`                // 路由名称（Vue Router 中用于命名路由）
	RoutePath  string      `json:"routePath"  orm:"route_path"  description:"路由路径（Vue Router 中定义的 URL 路径）"`            // 路由路径（Vue Router 中定义的 URL 路径）
	Component  string      `json:"component"  orm:"component"   description:"组件路径（组件页面完整路径，相对于 src/views/，缺省后缀 .vue）"` // 组件路径（组件页面完整路径，相对于 src/views/，缺省后缀 .vue）
	Perm       string      `json:"perm"       orm:"perm"        description:"【按钮】权限标识"`                                // 【按钮】权限标识
	AlwaysShow int         `json:"alwaysShow" orm:"always_show" description:"【目录】只有一个子路由是否始终显示（1-是 0-否）"`              // 【目录】只有一个子路由是否始终显示（1-是 0-否）
	KeepAlive  int         `json:"keepAlive"  orm:"keep_alive"  description:"【菜单】是否开启页面缓存（1-是 0-否）"`                   // 【菜单】是否开启页面缓存（1-是 0-否）
	Sort       int         `json:"sort"       orm:"sort"        description:"排序"`                                      // 排序
	Icon       string      `json:"icon"       orm:"icon"        description:"菜单图标"`                                    // 菜单图标
	Redirect   string      `json:"redirect"   orm:"redirect"    description:"跳转路径"`                                    // 跳转路径
	Params     string      `json:"params"     orm:"params"      description:"路由参数"`                                    // 路由参数
	Enable     int         `json:"enable"     orm:"enable"      description:"是否启用//radio/1,启用,2,禁用"`                   // 是否启用//radio/1,启用,2,禁用
	UpdateAt   *gtime.Time `json:"updateAt"   orm:"update_at"   description:"更新时间"`                                    // 更新时间
	UpdateId   int64       `json:"updateId"   orm:"update_id"   description:"更新人"`                                     // 更新人
	CreateAt   *gtime.Time `json:"createAt"   orm:"create_at"   description:"创建时间"`                                    // 创建时间
	CreateId   int64       `json:"createId"   orm:"create_id"   description:"创建者"`                                     // 创建者
}
