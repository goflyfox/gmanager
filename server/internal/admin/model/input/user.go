package input

import (
	"github.com/gogf/gf/v2/os/gtime"
	"gmanager/internal/admin/model/entity"
)

type UserOptionRes struct {
	Value int64  `json:"value"`
	Label string `json:"label"`
}

type User struct {
	entity.User
	DeptName string  `json:"deptName"`
	RoleIds  []int64 `json:"roleIds"`
}

type UserProfile struct {
	Id        int64       `json:"id"       description:"主键"`
	UserName  string      `json:"userName" description:"登录名/11111"`
	Mobile    string      `json:"mobile"   description:"手机号"`
	Email     string      `json:"email"    description:"email"`
	NickName  string      `json:"nickName" description:"昵称"`
	Gender    int         `json:"gender"   description:"性别;0:保密,1:男,2:女"`
	Avatar    string      `json:"avatar"   description:"头像地址"`
	CreateAt  *gtime.Time `json:"createAt" description:"创建时间"`
	DeptName  string      `json:"deptName" description:"部门"`
	RoleNames string      `json:"roleNames" description:"角色"`
}
