package service

import (
	"context"
	"gmanager/internal/common/model/input"
)

type (
	IStorage interface {
		// Upload 上传文件
		Upload(ctx context.Context, file *input.UploadFile) (res *input.UploadFileRes, err error)
	}
)

var (
	storage IStorage
)

func Storage() IStorage {
	if storage == nil {
		panic("implement not found for interface IStorage, forgot register?")
	}
	return storage
}

func RegisterStorage(i IStorage) {
	storage = i
}
