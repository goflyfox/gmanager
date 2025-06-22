package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	input2 "gmanager/internal/admin/model/input"
)

type MenuListReq struct {
	g.Meta `path:"/menu/list" method:"post" perms:"admin:menu:query" tags:"菜单管理" summary:"菜单列表"`
	Name   string `json:"name" dc:"菜单名称"`
	Enable int    `json:"enable" dc:"是否启用"`
	input2.PageReq
}

type MenuListRes struct {
	List []*input2.MenuTreeRes `json:"list" dc:"菜单列表"`
	input2.PageRes
}

type MenuOptionsReq struct {
	g.Meta     `path:"/menu/options" method:"post" tags:"菜单管理" summary:"菜单下拉列表"`
	OnlyParent bool `json:"onlyParent" dc:"是否仅父节点"`
	Enable     int  `json:"enable" dc:"是否启用"`
}

type MenuOptionsRes = []*input2.OptionVal

type MenuGetReq struct {
	g.Meta `path:"/menu/get/:id" method:"get"  perms:"admin:menu:query" tags:"菜单管理" summary:"菜单获取"`
	Id     int `json:"id" dc:"ID"`
}

type MenuGetRes = input2.Menu

type MenuSaveReq struct {
	g.Meta     `path:"/menu/save/:id" method:"post"  perms:"admin:menu:save" tags:"菜单管理" summary:"菜单保存"`
	Id         int                `json:"id"`
	ParentId   int                `json:"parentId"  v:"required#父级不能为空"`
	Name       string             `json:"name"   dc:"菜单名称" v:"required#菜单名称不能为空"`
	Code       string             `json:"code"   dc:"菜单编码"`
	Sort       int                `json:"sort" dc:"排序序号" v:"required#菜单序号不能为空"`
	Type       int                `json:"type"            dc:"菜单类型"`
	RouteName  string             `json:"routeName"  dc:"路由名称"`
	RoutePath  string             `json:"routePath"  dc:"路由路径"`
	Component  string             `json:"component"  dc:"组件路径"`
	Perm       string             `json:"perm"            dc:"按钮权限标识"`
	AlwaysShow int                `json:"alwaysShow" dc:"只有一个子路由是否始终显示"`
	KeepAlive  int                `json:"keepAlive"  dc:"【菜单】是否开启页面缓存（1-是 0-否）"`
	Icon       string             `json:"icon"            dc:"菜单图标"`
	Redirect   string             `json:"redirect"    dc:"跳转路径"`
	Params     string             `json:"params"        dc:"路由参数"`
	ParamList  []*input2.KeyValue `json:"paramList"  description:"路由参数列表"`
	Enable     int                `json:"enable"        dc:"是否启用//radio/1,启用,2,禁用"`
}

type MenuSaveRes struct {
}

type MenuDeleteReq struct {
	g.Meta `path:"/menu/delete/:ids" method:"post" perms:"admin:menu:delete" tags:"菜单管理" summary:"菜单删除"`
	Ids    string `json:"ids" dc:"删除id列表"`
}

type MenuDeleteRes struct {
}
