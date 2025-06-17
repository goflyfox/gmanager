package logic

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "gmanager/api/common/v1"
	"gmanager/internal/common/model/input"
	"gmanager/internal/common/service"
)

// Upload 上传服务
var Upload = new(upload)

type upload struct{}

// Upload 上传文件
func (s *upload) Upload(ctx context.Context, in *v1.UploadFileReq) (res *v1.UploadFileRes, err error) {
	uploadRes, err := service.Storage().Upload(ctx, &input.UploadFile{File: in.File})
	if err != nil {
		return
	}
	if uploadRes == nil {
		err = gerror.New("上传失败")
		return
	}
	res = &v1.UploadFileRes{
		Name: uploadRes.Name,
		Url:  uploadRes.FileUrl,
	}
	return
}
