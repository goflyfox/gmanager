package input

import (
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
