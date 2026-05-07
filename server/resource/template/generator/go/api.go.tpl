package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"gmanager/internal/admin/model/entity"
	"gmanager/internal/admin/model/input"
)

// ==================== {{.functionName}} ====================

type {{.className}}ListReq struct {
	g.Meta   `path:"/{{.businessName}}/list" method:"post" perms:"admin:{{.businessName}}:query" tags:"{{.functionName}}" summary:"{{.functionName}}列表"`
	Keywords string `json:"keywords" dc:"关键字"`
{{- range .queryColumns}}
{{- if or (eq .GoType "string") (eq .GoType "*gtime.Time")}}
	{{.GoField}} string `json:"{{.GoField}}" dc:"{{.ColumnComment}}"`
{{- else}}
	{{.GoField}} {{.GoType}} `json:"{{.GoField}}" dc:"{{.ColumnComment}}"`
{{- end}}
{{- end}}
	input.PageReq
}

type {{.className}}ListRes struct {
	List []*entity.{{.className}} `json:"list" dc:"列表"`
	input.PageRes
}

type {{.className}}GetReq struct {
	g.Meta `path:"/{{.businessName}}/get/:id" method:"get" perms:"admin:{{.businessName}}:query" tags:"{{.functionName}}" summary:"获取详情"`
	Id     int64 `json:"id" dc:"ID"`
}

type {{.className}}GetRes = entity.{{.className}}

type {{.className}}SaveReq struct {
	g.Meta `path:"/{{.businessName}}/save/:id" method:"post" perms:"admin:{{.businessName}}:save" tags:"{{.functionName}}" summary:"保存"`
	Id     int64 `json:"id"`
{{- range .formColumns}}
	{{.GoField}} {{.GoType}} `json:"{{.GoField}}" dc:"{{.ColumnComment}}"{{if eq .IsRequired "1"}} v:"required#{{.ColumnComment}}不能为空"{{end}}`
{{- end}}
}

type {{.className}}SaveRes struct {
}

type {{.className}}DeleteReq struct {
	g.Meta `path:"/{{.businessName}}/delete/:ids" method:"post" perms:"admin:{{.businessName}}:delete" tags:"{{.functionName}}" summary:"删除"`
	Ids    string `json:"ids" dc:"删除id列表"`
}

type {{.className}}DeleteRes struct {
}
