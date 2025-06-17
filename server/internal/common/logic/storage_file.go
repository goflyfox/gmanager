package logic

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"gmanager/internal/common/consts"
	"gmanager/internal/common/model/input"
	"gmanager/internal/common/service"
)

type storageFile struct{}

func NewStorageFile() *storageFile {
	return &storageFile{}
}

func init() {
	service.RegisterStorage(NewStorageFile())
}

func (s *storageFile) Upload(ctx context.Context, file *input.UploadFile) (res *input.UploadFileRes, err error) {
	// TODO 上传信息需要配置化
	var (
		attachmentPath = "attachment"
		nowDate        = gtime.Date()
		storageType    = consts.StorageTypeLocal
	)

	fullDirPath := gfile.MainPkgPath() + "/resource/public/" + attachmentPath + "/" + nowDate
	fileName, err := file.File.Save(fullDirPath, true)
	if err != nil {
		return
	}
	// 不含静态文件夹的路径
	fullPath := "/" + attachmentPath + "/" + nowDate + "/" + fileName
	res = &input.UploadFileRes{
		Name:    fileName,
		OriName: file.File.Filename,
		Type:    storageType,
		Path:    fullDirPath,
		FileUrl: "http://localhost:8000" + fullPath,
		Size:    file.File.Size,
		Ext:     gfile.Ext(fullPath),
	}
	g.Log().Info(ctx, "upload file", res)

	return
}
