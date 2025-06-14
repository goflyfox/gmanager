package input

import (
	"gmanager/internal/admin/model/entity"
)

type DeptTreeRes struct {
	*entity.Dept
	Children []*DeptTreeRes `json:"children"`
}
