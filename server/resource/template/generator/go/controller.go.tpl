package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/admin/logic"
	"strings"
)

type {{.businessName}} struct{}

var {{.className}} = new({{.businessName}})

func (c *{{.businessName}}) List(ctx context.Context, req *v1.{{.className}}ListReq) (res *v1.{{.className}}ListRes, err error) {
	res, err = logic.{{.className}}.List(ctx, req)
	return
}

func (c *{{.businessName}}) Get(ctx context.Context, req *v1.{{.className}}GetReq) (res *v1.{{.className}}GetRes, err error) {
	res, err = logic.{{.className}}.Get(ctx, req.Id)
	return
}

func (c *{{.businessName}}) Save(ctx context.Context, req *v1.{{.className}}SaveReq) (res *v1.{{.className}}SaveRes, err error) {
	err = logic.{{.className}}.Save(ctx, req)
	return
}

func (c *{{.businessName}}) Delete(ctx context.Context, req *v1.{{.className}}DeleteReq) (res *v1.{{.className}}DeleteRes, err error) {
	if req.Ids == "" {
		return
	}
	idArr := make([]int, 0)
	for _, v := range strings.Split(req.Ids, ",") {
		idArr = append(idArr, gconv.Int(v))
	}
	err = logic.{{.className}}.Delete(ctx, idArr)
	return
}
