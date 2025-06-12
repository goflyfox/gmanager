package bean

type SessionUser struct {
	Id       int64  `form:"id" json:"id"`             // 主键
	Uuid     string `form:"uuid" json:"uuid"`         // UUID
	Username string `form:"username" json:"username"` // 登录名/11111
	NickName string `form:"nickname" json:"nickname"` // 昵称
}
