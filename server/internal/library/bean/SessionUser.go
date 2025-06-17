package bean

type SessionUser struct {
	Id       int64  `json:"id"`       // 主键
	Uuid     string `json:"uuid"`     // UUID
	UserName string `json:"userName"` // 登录名/11111
	NickName string `json:"nickName"` // 昵称
	UserType int    `json:"userType"` // 类型
}
