package base

type BaseModel struct {
	Enable     int    `json:"enable" gconv:"enable,omitempty"`
	UpdateTime string `json:"updateTime" gconv:"update_time,omitempty"`
	UpdateId   int    `json:"updateId" gconv:"update_id,omitempty"`
	UpdateName string `json:"updateName,omitempty" gconv:"updateName,omitempty"`
	CreateTime string `json:"createTime" gconv:"create_time,omitempty"`
	CreateId   int    `json:"createId" gconv:"create_id,omitempty"`
	CreateName string `json:"createName,omitempty" gconv:"createName,omitempty"`
}

// 定义model interface
type IModel interface {
	// 获取表明
	TableName() string
	// 获取主键值
	PkVal() int
}
