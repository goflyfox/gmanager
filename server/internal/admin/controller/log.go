package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/admin/logic"
	"strings"
)

type log struct{}

var Log = new(log)

func (c *log) List(ctx context.Context, req *v1.LogListReq) (res *v1.LogListRes, err error) {
	res, err = logic.Log.List(ctx, req)
	return
}

func (c *log) Delete(ctx context.Context, req *v1.LogDeleteReq) (res *v1.LogDeleteRes, err error) {
	if req.Ids == "" {
		return
	}
	idArr := make([]int, 0)
	for _, v := range strings.Split(req.Ids, ",") {
		idArr = append(idArr, gconv.Int(v))
	}
	err = logic.Log.Delete(ctx, idArr)
	return
}

func (c *log) Get(ctx context.Context, req *v1.LogGetReq) (res *v1.LogGetRes, err error) {
	res, err = logic.Log.Get(ctx, req.Id)
	return
}

func (c *log) VisitTrend(ctx context.Context, req *v1.LogVisitTrendReq) (res *v1.LogVisitTrendRes, err error) {
	res, err = logic.Log.VisitTrend(ctx, req)
	return
}
