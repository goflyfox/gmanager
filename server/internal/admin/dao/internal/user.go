// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for the table sys_user.
type UserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserColumns defines and stores column names for the table sys_user.
type UserColumns struct {
	Id       string // 主键
	Uuid     string // UUID
	UserName string // 登录名/11111
	Mobile   string // 手机号
	Email    string // email
	Password string // 密码
	Salt     string // 密码盐
	DeptId   string // 部门/11111/dict
	UserType string // 类型//select/1,管理员,2,普通用户,3,前台用户,4,第三方用户,5,API用户
	Status   string // 状态
	Thirdid  string // 第三方ID
	Endtime  string // 结束时间
	NickName string // 昵称
	Gender   string // 性别;0:保密,1:男,2:女
	Address  string // 地址
	Avatar   string // 头像地址
	Birthday string // 生日
	Remark   string // 说明
	Enable   string // 是否启用//radio/1,启用,2,禁用
	UpdateAt string // 更新时间
	UpdateId string // 更新人
	CreateAt string // 创建时间
	CreateId string // 创建者
}

// userColumns holds the columns for the table sys_user.
var userColumns = UserColumns{
	Id:       "id",
	Uuid:     "uuid",
	UserName: "user_name",
	Mobile:   "mobile",
	Email:    "email",
	Password: "password",
	Salt:     "salt",
	DeptId:   "dept_id",
	UserType: "user_type",
	Status:   "status",
	Thirdid:  "thirdid",
	Endtime:  "endtime",
	NickName: "nick_name",
	Gender:   "gender",
	Address:  "address",
	Avatar:   "avatar",
	Birthday: "birthday",
	Remark:   "remark",
	Enable:   "enable",
	UpdateAt: "update_at",
	UpdateId: "update_id",
	CreateAt: "create_at",
	CreateId: "create_id",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao(handlers ...gdb.ModelHandler) *UserDao {
	return &UserDao{
		group:    "default",
		table:    "sys_user",
		columns:  userColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
