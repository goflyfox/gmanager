package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"gmanager/internal/model/entity"
	"gmanager/internal/model/input"
)

type ConfigListReq struct {
	g.Meta   `path:"/config/list" method:"POST" tags:"配置管理" summary:"配置列表"`
	Keywords string `json:"keywords" dc:"名称"`
	DataType int    `json:"dataType" dc:"数据类型"`
	Enable   int    `json:"enable" dc:"是否启用"`
	input.PageReq
}

type ConfigListRes struct {
	List []*entity.Config `json:"list" dc:"配置列表"`
	input.PageRes
}

type ConfigGetReq struct {
	g.Meta `path:"/config/get/:id" method:"get" tags:"配置管理" summary:"配置获取"`
	Id     int64 `json:"id" dc:"ID"`
}

type ConfigGetRes = entity.Config

type ConfigSaveReq struct {
	g.Meta       `path:"/config/save/:id" method:"post" tags:"配置管理" summary:"配置保存"`
	Id           int64  `json:"id"`
	Name         string `json:"name"  dc:"配置名称" v:"required#配置名称不能为空"`
	Key          string `json:"key"   dc:"配置键" v:"required#键不能为空"`
	Value        string `json:"value"   dc:"配置值" `
	Code         string `json:"code" dc:"编码"`
	DataType     int    `json:"dataType" dc:"数据类型"`
	ParentId     int64  `json:"parentId" dc:"类型"`
	Remark       string `json:"remark" dc:"备注"`
	Sort         string `json:"sort" dc:"排序"`
	Enable       int    `json:"enable" dc:"是否启用//radio/1,启用,2,禁用"`
	CopyStatus   int    `json:"copyStatus"   dc:"拷贝状态"`
	ChangeStatus int    `json:"changeStatus" dc:"更新状态"`
}

type ConfigSaveRes struct {
}

type ConfigDeleteReq struct {
	g.Meta `path:"/config/delete/:ids" method:"post" tags:"配置管理" summary:"配置删除"`
	Ids    string `json:"ids" dc:"删除id列表"`
}

type ConfigDeleteRes struct {
}

type ConfigRefreshReq struct {
	g.Meta `path:"/config/refresh" method:"post" tags:"配置管理" summary:"配置缓存刷新"`
}

type ConfigRefreshRes struct {
}

type ConfigDictOptionsReq struct {
	g.Meta `path:"/config/dict/options" method:"post" tags:"配置管理" summary:"字典下拉列表"`
	Enable int `json:"enable" dc:"是否启用"`
}

type ConfigDictOptionsRes = []*input.OptionVal

type ConfigValueReq struct {
	g.Meta `path:"/config/value/:key" method:"post" tags:"配置管理" summary:"获取配置对应信息"`
	Key    string `json:"key"   dc:"配置键" v:"required#键不能为空"`
}

type ConfigValueRes struct {
	Id    int64  `json:"id"`
	Value string `json:"value"`
	Code  string `json:"code"`
}

type ConfigDictItemsReq struct {
	g.Meta    `path:"/config/dict/items/:parentKey" method:"post" tags:"配置管理" summary:"获取配置字典对应数据列表"`
	ParentKey string `json:"parentKey" dc:"类型Key" v:"required#字典Key不能为空"`
}

type ConfigDictItemsRes = []*input.OptionVal
