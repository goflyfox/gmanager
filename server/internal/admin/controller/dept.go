package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/admin/logic"
	"gmanager/internal/consts"
	"strings"
)

type dept struct{}

var Dept = new(dept)

func (c *dept) List(ctx context.Context, req *v1.DeptListReq) (res *v1.DeptListRes, err error) {
	res, err = logic.Dept.List(ctx, req)
	return
}

func (c *dept) Options(ctx context.Context, req *v1.DeptOptionsReq) (res *v1.DeptOptionsRes, err error) {
	if req != nil && req.Enable == 0 {
		req.Enable = consts.EnableYes
	}
	res, err = logic.Dept.Options(ctx, req)
	return
}

func (c *dept) Save(ctx context.Context, req *v1.DeptSaveReq) (res *v1.DeptSaveRes, err error) {
	err = logic.Dept.Save(ctx, req)
	return
}

func (c *dept) Delete(ctx context.Context, req *v1.DeptDeleteReq) (res *v1.DeptDeleteRes, err error) {
	if req.Ids == "" {
		return
	}
	idArr := make([]int, 0)
	for _, v := range strings.Split(req.Ids, ",") {
		idArr = append(idArr, gconv.Int(v))
	}
	err = logic.Dept.Delete(ctx, idArr)
	return
}

func (c *dept) Get(ctx context.Context, req *v1.DeptGetReq) (res *v1.DeptGetRes, err error) {
	res, err = logic.Dept.Get(ctx, req.Id)
	return
}
