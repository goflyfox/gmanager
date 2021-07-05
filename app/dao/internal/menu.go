// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// MenuDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type MenuDao struct {
	gmvc.M              // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB      // DB is the raw underlying database management object.
	Table   string      // Table is the table name of the DAO.
	Columns menuColumns // Columns contains all the columns of Table that for convenient usage.
}

// MenuColumns defines and stores column names for table sys_menu.
type menuColumns struct {
	Id         string // 主键
	ParentId   string // 父id
	Name       string // 名称/11111
	Icon       string // 菜单图标
	Urlkey     string // 菜单key
	Url        string // 链接地址
	Perms      string // 授权(多个用逗号分隔，如：user:list,user:create)
	Status     string // 状态//radio/2,隐藏,1,显示
	Type       string // 类型//select/1,目录,2,菜单,3,按钮
	Sort       string // 排序
	Level      string // 级别
	Enable     string // 是否启用//radio/1,启用,2,禁用
	UpdateTime string // 更新时间
	UpdateId   string // 更新人
	CreateTime string // 创建时间
	CreateId   string // 创建者
}

func NewMenuDao() *MenuDao {
	return &MenuDao{
		M:     g.DB("default").Model("sys_menu").Safe(),
		DB:    g.DB("default"),
		Table: "sys_menu",
		Columns: menuColumns{
			Id:         "id",
			ParentId:   "parent_id",
			Name:       "name",
			Icon:       "icon",
			Urlkey:     "urlkey",
			Url:        "url",
			Perms:      "perms",
			Status:     "status",
			Type:       "type",
			Sort:       "sort",
			Level:      "level",
			Enable:     "enable",
			UpdateTime: "update_time",
			UpdateId:   "update_id",
			CreateTime: "create_time",
			CreateId:   "create_id",
		},
	}
}