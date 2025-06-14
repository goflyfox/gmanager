package input

import (
	"gmanager/internal/admin/model/entity"
)

type Menu struct {
	*entity.Menu
	ParamList []*KeyValue `json:"paramList"  description:"路由参数列表"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MenuTreeRes struct {
	*entity.Menu
	Children []*MenuTreeRes `json:"children"`
}

type UserMenu struct {
	Id        int64       `json:"id"`
	Path      string      `json:"path"`
	Component string      `json:"component"`
	Redirect  string      `json:"redirect"`
	Name      string      `json:"name"`
	Meta      Meta        `json:"meta"`
	Children  []*UserMenu `json:"children,omitempty"`
}

type Meta struct {
	Title      string `json:"title"`
	Icon       string `json:"icon"`
	Hidden     bool   `json:"hidden"`
	KeepAlive  bool   `json:"keepAlive"`
	AlwaysShow bool   `json:"alwaysShow"`
}
