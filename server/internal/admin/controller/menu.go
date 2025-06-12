package controller

import (
	"context"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/admin/logic"
)

type menu struct{}

var Menu = new(menu)

func (c *menu) List(ctx context.Context, req *v1.MenuListReq) (res *v1.MenuListRes, err error) {
	res, err = logic.Menu.List(ctx, req)
	return
}

func (c *menu) Get(ctx context.Context, req *v1.MenuGetReq) (res *v1.MenuGetRes, err error) {
	res = &v1.MenuGetRes{}
	res, err = logic.Menu.Get(ctx, req.Id)
	return
}

func (c *menu) Options(ctx context.Context, req *v1.MenuOptionsReq) (res *v1.MenuOptionsRes, err error) {
	res, err = logic.Menu.Options(ctx, req)
	return
}

func (c *menu) Save(ctx context.Context, req *v1.MenuSaveReq) (res *v1.MenuSaveRes, err error) {
	err = logic.Menu.Save(ctx, req)
	return
}

func (c *menu) Delete(ctx context.Context, req *v1.MenuDeleteReq) (res *v1.MenuDeleteRes, err error) {
	err = logic.Menu.Delete(ctx, req.Ids)
	return
}
