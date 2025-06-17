package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// UploadFileReq 上传文件
type UploadFileReq struct {
	g.Meta `path:"/upload/file" tags:"附件" method:"post" summary:"上传附件"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"文件"`
}

type UploadFileRes struct {
	Url  string `json:"url"   dc:"URL"`
	Name string `json:"name"   dc:"名称"`
}
