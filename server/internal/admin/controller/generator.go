package controller

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/admin/logic"
	"strings"
)

type generator struct{}

var Generator = new(generator)

func (c *generator) TableList(ctx context.Context, req *v1.GeneratorTableListReq) (res *v1.GeneratorTableListRes, err error) {
	res, err = logic.Generator.List(ctx, req)
	return
}

func (c *generator) DbTableList(ctx context.Context, req *v1.GeneratorDbTableListReq) (res *v1.GeneratorDbTableListRes, err error) {
	res, err = logic.Generator.DbTableList(ctx, req)
	return
}

func (c *generator) TableImport(ctx context.Context, req *v1.GeneratorTableImportReq) (res *v1.GeneratorTableImportRes, err error) {
	err = logic.Generator.ImportTable(ctx, req.Names)
	return
}

func (c *generator) TableGet(ctx context.Context, req *v1.GeneratorTableGetReq) (res *v1.GeneratorTableGetRes, err error) {
	res, err = logic.Generator.Get(ctx, req.Id)
	return
}

func (c *generator) TableSave(ctx context.Context, req *v1.GeneratorTableSaveReq) (res *v1.GeneratorTableSaveRes, err error) {
	err = logic.Generator.Save(ctx, req)
	return
}

func (c *generator) TableDelete(ctx context.Context, req *v1.GeneratorTableDeleteReq) (res *v1.GeneratorTableDeleteRes, err error) {
	if req.Ids == "" {
		return
	}
	idArr := make([]int, 0)
	for _, v := range strings.Split(req.Ids, ",") {
		idArr = append(idArr, gconv.Int(v))
	}
	err = logic.Generator.Delete(ctx, idArr)
	return
}

func (c *generator) Preview(ctx context.Context, req *v1.GeneratorPreviewReq) (res *v1.GeneratorPreviewRes, err error) {
	res, err = logic.Generator.Preview(ctx, req.Id)
	return
}

func (c *generator) Download(ctx context.Context, req *v1.GeneratorDownloadReq) (res *v1.GeneratorDownloadRes, err error) {
	buf, fileName, err := logic.Generator.Download(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	// 直接操作 response，GoFrame 的 controller 可以通过 ghttp.RequestFromCtx 获取
	request := ghttp.RequestFromCtx(ctx)
	request.Response.Header().Set("Content-Type", "application/zip")
	request.Response.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	request.Response.Write(buf.Bytes())
	// 返回 nil 阻止框架再次写入
	return nil, nil
}

func (c *generator) GenCode(ctx context.Context, req *v1.GeneratorGenCodeReq) (res *v1.GeneratorGenCodeRes, err error) {
	err = logic.Generator.GenCode(ctx, req.Id, req.Path)
	return
}
