package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/admin/logic"
	"strings"
)

type role struct{}

var Role = new(role)

func (c *role) List(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	res, err = logic.Role.List(ctx, req)
	return
}

func (c *role) Options(ctx context.Context, req *v1.RoleOptionsReq) (res *v1.RoleOptionsRes, err error) {
	res, err = logic.Role.Options(ctx, req)
	return
}

func (c *role) Get(ctx context.Context, req *v1.RoleGetReq) (res *v1.RoleGetRes, err error) {
	res, err = logic.Role.Get(ctx, req.Id)
	return
}

func (c *role) Save(ctx context.Context, req *v1.RoleSaveReq) (res *v1.RoleSaveRes, err error) {
	err = logic.Role.Save(ctx, req)
	return
}

func (c *role) Delete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error) {
	if req.Ids == "" {
		return
	}
	idArr := make([]int, 0)
	for _, v := range strings.Split(req.Ids, ",") {
		idArr = append(idArr, gconv.Int(v))
	}
	err = logic.Role.Delete(ctx, idArr)
	return
}

func (c *role) MenuIds(ctx context.Context, req *v1.RoleMenuIdsReq) (res *v1.RoleMenuIdsRes, err error) {
	res, err = logic.Role.MenuIds(ctx, req.Id)
	return
}

func (c *role) AddMenuIds(ctx context.Context, req *v1.RoleAddMenuIdsReq) (res *v1.RoleAddMenuIdsRes, err error) {
	err = logic.Role.AddMenuIds(ctx, req)
	return
}
