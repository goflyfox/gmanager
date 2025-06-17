package controller

import (
	"context"
	v1 "gmanager/api/common/v1"
	"gmanager/internal/common/logic"
)

type upload struct{}

var Upload = new(upload)

func (c *upload) Upload(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error) {
	res, err = logic.Upload.Upload(ctx, req)
	return
}
