package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "gmanager/api/admin/v1"
)

var Ping = new(ping)

type ping struct{}

func (c *ping) Ping(ctx context.Context, req *v1.PingReq) (res *v1.PingRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("pong")
	return
}
