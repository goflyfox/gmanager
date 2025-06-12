package v1

import "github.com/gogf/gf/v2/frame/g"

// PingReq ping
type PingReq struct {
	g.Meta `path:"/ping" method:"get" tags:"Ping" summary:"ping"`
}

type PingRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
