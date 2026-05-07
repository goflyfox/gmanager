package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"gmanager/internal/admin/model/entity"
	"gmanager/internal/admin/model/input"
)

// ==================== 代码生成表管理 ====================

type GeneratorTableListReq struct {
	g.Meta   `path:"/generator/table/list" method:"post" perms:"admin:generator:query" tags:"代码生成" summary:"代码生成表列表"`
	Keywords string `json:"keywords" dc:"表名称/表描述"`
	input.PageReq
}

type GeneratorTableListRes struct {
	List []*entity.GenTable `json:"list" dc:"表列表"`
	input.PageRes
}

type GeneratorTableImportReq struct {
	g.Meta `path:"/generator/table/import" method:"post" perms:"admin:generator:save" tags:"代码生成" summary:"导入表结构"`
	Names  []string `json:"names" v:"required#至少选择一张表"`
}

type GeneratorTableImportRes struct {
}

type GeneratorTableGetReq struct {
	g.Meta `path:"/generator/table/get/:id" method:"get" perms:"admin:generator:query" tags:"代码生成" summary:"获取表配置详情"`
	Id     int64 `json:"id" dc:"ID"`
}

type GeneratorTableGetRes struct {
	Info    *entity.GenTable         `json:"info"`
	Columns []*entity.GenTableColumn `json:"columns"`
}

type GeneratorTableSaveReq struct {
	g.Meta         `path:"/generator/table/save" method:"post" perms:"admin:generator:save" tags:"代码生成" summary:"保存表配置"`
	Id             int64                    `json:"id"`
	TableName      string                   `json:"tableName" v:"required#表名称不能为空"`
	TableComment   string                   `json:"tableComment"`
	ClassName      string                   `json:"className" v:"required#实体类名不能为空"`
	PackageName    string                   `json:"packageName"`
	ModuleName     string                   `json:"moduleName" v:"required#模块名不能为空"`
	BusinessName   string                   `json:"businessName" v:"required#业务名不能为空"`
	FunctionName   string                   `json:"functionName" v:"required#功能名不能为空"`
	FunctionAuthor string                   `json:"functionAuthor"`
	TplCategory    string                   `json:"tplCategory"`
	GenType        string                   `json:"genType"`
	GenPath        string                   `json:"genPath"`
	Options        string                   `json:"options"`
	Columns        []*entity.GenTableColumn `json:"columns"`
}

type GeneratorTableSaveRes struct {
}

type GeneratorTableDeleteReq struct {
	g.Meta `path:"/generator/table/delete/:ids" method:"post" perms:"admin:generator:delete" tags:"代码生成" summary:"删除表配置"`
	Ids    string `json:"ids" dc:"删除id列表"`
}

type GeneratorTableDeleteRes struct {
}

// 获取数据库表列表（用于导入选择）
type GeneratorDbTableListReq struct {
	g.Meta   `path:"/generator/db/table/list" method:"post" perms:"admin:generator:query" tags:"代码生成" summary:"数据库表列表"`
	Keywords string `json:"keywords" dc:"表名称"`
	input.PageReq
}

type GeneratorDbTableListRes struct {
	List []*entity.GenTable `json:"list" dc:"表列表"`
	input.PageRes
}

// ==================== 代码生成 ====================

type GeneratorPreviewReq struct {
	g.Meta `path:"/generator/preview/:id" method:"get" perms:"admin:generator:query" tags:"代码生成" summary:"预览代码"`
	Id     int64 `json:"id" dc:"ID"`
}

type GeneratorPreviewRes struct {
	Data map[string]string `json:"data" dc:"文件名->代码内容"`
}

type GeneratorDownloadReq struct {
	g.Meta `path:"/generator/download/:id" method:"get" perms:"admin:generator:query" tags:"代码生成" summary:"下载ZIP"`
	Id     int64 `json:"id" dc:"ID"`
}

type GeneratorDownloadRes struct {
	g.Meta `mime:"application/zip"`
}

type GeneratorGenCodeReq struct {
	g.Meta `path:"/generator/gen/:id" method:"post" perms:"admin:generator:save" tags:"代码生成" summary:"生成代码（写入磁盘）"`
	Id     int64  `json:"id"`
	Path   string `json:"path" dc:"生成路径，空则使用配置路径"`
}

type GeneratorGenCodeRes struct {
}
