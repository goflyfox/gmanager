// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// DepartmentDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type DepartmentDao struct {
	gmvc.M                    // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB            // DB is the raw underlying database management object.
	Table   string            // Table is the table name of the DAO.
	Columns departmentColumns // Columns contains all the columns of Table that for convenient usage.
}

// DepartmentColumns defines and stores column names for table sys_department.
type departmentColumns struct {
	Id         string // 主键
	ParentId   string // 上级机构
	Name       string // 部门/11111
	Code       string // 机构编码
	Sort       string // 序号
	Linkman    string // 联系人
	LinkmanNo  string // 联系人电话
	Remark     string // 机构描述
	Enable     string // 是否启用//radio/1,启用,2,禁用
	UpdateTime string // 更新时间
	UpdateId   string // 更新人
	CreateTime string // 创建时间
	CreateId   string // 创建者
}

func NewDepartmentDao() *DepartmentDao {
	return &DepartmentDao{
		M:     g.DB("default").Model("sys_department").Safe(),
		DB:    g.DB("default"),
		Table: "sys_department",
		Columns: departmentColumns{
			Id:         "id",
			ParentId:   "parent_id",
			Name:       "name",
			Code:       "code",
			Sort:       "sort",
			Linkman:    "linkman",
			LinkmanNo:  "linkman_no",
			Remark:     "remark",
			Enable:     "enable",
			UpdateTime: "update_time",
			UpdateId:   "update_id",
			CreateTime: "create_time",
			CreateId:   "create_id",
		},
	}
}
