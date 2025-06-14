// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MenuDao is the data access object for the table sys_menu.
type MenuDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MenuColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MenuColumns defines and stores column names for the table sys_menu.
type MenuColumns struct {
	Id         string // ID
	ParentId   string // 父菜单ID
	Name       string // 菜单名称
	Type       string // 菜单类型（1-菜单 2-目录 3-外链 4-按钮）
	RouteName  string // 路由名称（Vue Router 中用于命名路由）
	RoutePath  string // 路由路径（Vue Router 中定义的 URL 路径）
	Component  string // 组件路径（组件页面完整路径，相对于 src/views/，缺省后缀 .vue）
	Perm       string // 【按钮】权限标识
	AlwaysShow string // 【目录】只有一个子路由是否始终显示（1-是 0-否）
	KeepAlive  string // 【菜单】是否开启页面缓存（1-是 0-否）
	Sort       string // 排序
	Icon       string // 菜单图标
	Redirect   string // 跳转路径
	Params     string // 路由参数
	Enable     string // 是否启用//radio/1,启用,2,禁用
	UpdateAt   string // 更新时间
	UpdateId   string // 更新人
	CreateAt   string // 创建时间
	CreateId   string // 创建者
}

// menuColumns holds the columns for the table sys_menu.
var menuColumns = MenuColumns{
	Id:         "id",
	ParentId:   "parent_id",
	Name:       "name",
	Type:       "type",
	RouteName:  "route_name",
	RoutePath:  "route_path",
	Component:  "component",
	Perm:       "perm",
	AlwaysShow: "always_show",
	KeepAlive:  "keep_alive",
	Sort:       "sort",
	Icon:       "icon",
	Redirect:   "redirect",
	Params:     "params",
	Enable:     "enable",
	UpdateAt:   "update_at",
	UpdateId:   "update_id",
	CreateAt:   "create_at",
	CreateId:   "create_id",
}

// NewMenuDao creates and returns a new DAO object for table data access.
func NewMenuDao(handlers ...gdb.ModelHandler) *MenuDao {
	return &MenuDao{
		group:    "default",
		table:    "sys_menu",
		columns:  menuColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MenuDao) Columns() MenuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MenuDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *MenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
