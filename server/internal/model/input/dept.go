package input

import "gmanager/internal/model/entity"

type DeptTreeRes struct {
	*entity.Dept
	Children []*DeptTreeRes `json:"children"`
}
