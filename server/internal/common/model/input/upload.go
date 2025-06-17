package input

import "github.com/gogf/gf/v2/net/ghttp"

type UploadFile struct {
	File *ghttp.UploadFile `json:"file" type:"file" dc:"文件"`
}

type UploadFileRes struct {
	Name    string `json:"name"          description:"文件名"`
	OriName string `json:"oriName"  description:"文件原始名"`
	Path    string `json:"path"     description:"本地路径"`
	FileUrl string `json:"fileUrl"  description:"url"`
	Type    int    `json:"storageType" description:"存储类型"`
	Size    int64  `json:"size"     description:"文件大小"`
	Ext     string `json:"ext"      description:"扩展名"`
}
