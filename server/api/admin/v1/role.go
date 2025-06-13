package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"gmanager/internal/model/entity"
	"gmanager/internal/model/input"
)

type RoleListReq struct {
	g.Meta   `path:"/role/list" method:"post" tags:"角色管理" summary:"角色列表"`
	Keywords string `json:"keywords" dc:"角色或编码名称"`
	Name     string `json:"name" dc:"角色名称"`
	Enable   int    `json:"enable" dc:"是否启用"`
	input.PageReq
}

type RoleListRes struct {
	List []*entity.Role `json:"list" dc:"角色列表"`
	input.PageRes
}

type RoleOptionsReq struct {
	g.Meta `path:"/role/options" method:"post" tags:"角色管理" summary:"角色下拉列表"`
	Enable int `json:"enable" dc:"是否启用"`
}

type RoleOptionsRes = []*input.OptionVal

type RoleSaveReq struct {
	g.Meta    `path:"/role/save/:id" method:"post" tags:"角色管理" summary:"角色保存"`
	Id        int64  `json:"id"`
	Name      string `json:"name"  dc:"名称" v:"required#名称不能为空"`
	Code      string `json:"code" dc:"编码"`
	DataScope int    `json:"dataScope" dc:"数据范围"`
	Sort      int    `json:"sort"      dc:"排序"`
	Remark    string `json:"remark"    dc:"说明"`
	Enable    int    `json:"enable" dc:"是否启用"`
}

type RoleSaveRes struct {
}

type RoleGetReq struct {
	g.Meta `path:"/role/get/:id" method:"get" tags:"角色管理" summary:"角色获取"`
	Id     int64 `json:"id" dc:"ID"`
}

type RoleGetRes = entity.Role

type RoleDeleteReq struct {
	g.Meta `path:"/role/delete/:ids" method:"post" tags:"角色管理" summary:"角色删除"`
	Ids    string `json:"ids" dc:"删除id列表"`
}

type RoleDeleteRes struct {
}

type RoleMenuIdsReq struct {
	g.Meta `path:"/role/menuIds/:id" method:"post" tags:"角色管理" summary:"获取角色对应的菜单列表"`
	Id     int64 `json:"id" dc:"ID"`
}

type RoleMenuIdsRes = []int64

type RoleAddMenuIdsReq struct {
	g.Meta  `path:"/role/addMenus/:id" method:"post" tags:"角色管理" summary:"添加角色对应的菜单列表"`
	Id      int64   `json:"id" dc:"ID"`
	MenuIds []int64 `json:"menuIds" dc:"菜单列表"`
}

type RoleAddMenuIdsRes struct {
}
