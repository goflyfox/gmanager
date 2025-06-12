package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/admin/logic"
	"strings"
)

type config struct{}

var Config = new(config)

func (c *config) List(ctx context.Context, req *v1.ConfigListReq) (res *v1.ConfigListRes, err error) {
	res, err = logic.Config.List(ctx, req)
	return
}

func (c *role) DictOptions(ctx context.Context, req *v1.ConfigDictOptionsReq) (res *v1.ConfigDictOptionsRes, err error) {
	res, err = logic.Config.DictOptions(ctx, req)
	return
}

func (c *role) Value(ctx context.Context, req *v1.ConfigValueReq) (res *v1.ConfigValueRes, err error) {
	res, err = logic.Config.Value(ctx, req)
	return
}

func (c *role) DictItems(ctx context.Context, req *v1.ConfigDictItemsReq) (res *v1.ConfigDictItemsRes, err error) {
	res, err = logic.Config.DictItems(ctx, req)
	return
}

func (c *config) Get(ctx context.Context, req *v1.ConfigGetReq) (res *v1.ConfigGetRes, err error) {
	res, err = logic.Config.Get(ctx, req.Id)
	return
}

func (c *config) Save(ctx context.Context, req *v1.ConfigSaveReq) (res *v1.ConfigSaveRes, err error) {
	err = logic.Config.Save(ctx, req)
	return
}

func (c *config) Delete(ctx context.Context, req *v1.ConfigDeleteReq) (res *v1.ConfigDeleteRes, err error) {
	if req.Ids == "" {
		return
	}
	idArr := make([]int, 0)
	for _, v := range strings.Split(req.Ids, ",") {
		idArr = append(idArr, gconv.Int(v))
	}
	err = logic.Config.Delete(ctx, idArr)
	return
}

func (c *config) Refresh(ctx context.Context, req *v1.ConfigRefreshReq) (res *v1.ConfigRefreshRes, err error) {
	err = logic.Config.Refresh(ctx)
	return
}
