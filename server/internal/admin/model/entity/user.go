// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id       int64       `json:"id"       orm:"id"        description:"主键"`                                             // 主键
	Uuid     string      `json:"uuid"     orm:"uuid"      description:"UUID"`                                           // UUID
	UserName string      `json:"userName" orm:"user_name" description:"登录名/11111"`                                      // 登录名/11111
	Mobile   string      `json:"mobile"   orm:"mobile"    description:"手机号"`                                            // 手机号
	Email    string      `json:"email"    orm:"email"     description:"email"`                                          // email
	Password string      `json:"password" orm:"password"  description:"密码"`                                             // 密码
	Salt     string      `json:"salt"     orm:"salt"      description:"密码盐"`                                            // 密码盐
	DeptId   int64       `json:"deptId"   orm:"dept_id"   description:"部门/11111/dict"`                                  // 部门/11111/dict
	UserType int         `json:"userType" orm:"user_type" description:"类型//select/1,管理员,2,普通用户,3,前台用户,4,第三方用户,5,API用户"` // 类型//select/1,管理员,2,普通用户,3,前台用户,4,第三方用户,5,API用户
	Status   int         `json:"status"   orm:"status"    description:"状态"`                                             // 状态
	Thirdid  string      `json:"thirdid"  orm:"thirdid"   description:"第三方ID"`                                          // 第三方ID
	Endtime  string      `json:"endtime"  orm:"endtime"   description:"结束时间"`                                           // 结束时间
	NickName string      `json:"nickName" orm:"nick_name" description:"昵称"`                                             // 昵称
	Gender   int         `json:"gender"   orm:"gender"    description:"性别;0:保密,1:男,2:女"`                                // 性别;0:保密,1:男,2:女
	Address  string      `json:"address"  orm:"address"   description:"地址"`                                             // 地址
	Avatar   string      `json:"avatar"   orm:"avatar"    description:"头像地址"`                                           // 头像地址
	Birthday int         `json:"birthday" orm:"birthday"  description:"生日"`                                             // 生日
	Remark   string      `json:"remark"   orm:"remark"    description:"说明"`                                             // 说明
	Enable   int         `json:"enable"   orm:"enable"    description:"是否启用//radio/1,启用,2,禁用"`                          // 是否启用//radio/1,启用,2,禁用
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:"更新时间"`                                           // 更新时间
	UpdateId int64       `json:"updateId" orm:"update_id" description:"更新人"`                                            // 更新人
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:"创建时间"`                                           // 创建时间
	CreateId int64       `json:"createId" orm:"create_id" description:"创建者"`                                            // 创建者
}
