package logic

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	v1 "gmanager/api/admin/v1"
	"gmanager/internal/admin/dao"
	"gmanager/internal/admin/model/do"
	"gmanager/internal/admin/model/entity"
	"gmanager/internal/library/gftoken"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// Generator 代码生成服务
var Generator = new(generator)

type generator struct{}

// ==================== 表管理 ====================

// List 获取代码生成表列表
func (s *generator) List(ctx context.Context, in *v1.GeneratorTableListReq) (res *v1.GeneratorTableListRes, err error) {
	if in == nil {
		return
	}
	m := dao.GenTable.Ctx(ctx)
	columns := dao.GenTable.Columns()
	res = &v1.GeneratorTableListRes{}

	if in.Keywords != "" {
		m = m.Where(m.Builder().WhereLike(columns.TableName, "%"+in.Keywords+"%").
			WhereOrLike(columns.TableComment, "%"+in.Keywords+"%"))
	}

	res.Total, err = m.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据数量失败！")
		return
	}
	res.CurrentPage = in.PageNum
	if res.Total == 0 {
		return
	}

	if in.NeedOrderBy() {
		m = m.Order(in.OrderBy)
	} else {
		m = m.Order("id desc")
	}
	var pageList []*entity.GenTable
	if err = m.Page(in.PageNum, in.PageSize).Scan(&pageList); err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	res.List = pageList
	return
}

// DbTableList 获取数据库表列表（information_schema）
func (s *generator) DbTableList(ctx context.Context, in *v1.GeneratorDbTableListReq) (res *v1.GeneratorDbTableListRes, err error) {
	res = &v1.GeneratorDbTableListRes{}

	config := g.Cfg().MustGet(ctx, "database.default.link").String()
	// 解析 link 获取数据库名，格式如: mysql:root:123456@tcp(127.0.0.1:3306)/gmanager
	var dbName string
	if idx := strings.LastIndex(config, "/"); idx > 0 {
		dbName = strings.TrimSpace(config[idx+1:])
		if idx2 := strings.Index(dbName, "?"); idx2 > 0 {
			dbName = dbName[:idx2]
		}
	}
	if dbName == "" {
		err = gerror.New("无法获取数据库名称，请检查 database.default.link 配置")
		return
	}

	m := g.DB().Model("information_schema.tables").
		Where("table_schema", dbName).
		Where("table_type", "BASE TABLE")

	if in.Keywords != "" {
		m = m.WhereLike("table_name", "%"+in.Keywords+"%")
	}

	// 排除已导入的表
	importedNames, err := dao.GenTable.Ctx(ctx).Fields(dao.GenTable.Columns().TableName).Array()
	if err != nil {
		return
	}
	if len(importedNames) > 0 {
		m = m.WhereNotIn("table_name", importedNames)
	}

	res.Total, err = m.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取数据数量失败！")
		return
	}
	res.CurrentPage = in.PageNum
	if res.Total == 0 {
		return
	}

	var list []g.Map
	err = m.Page(in.PageNum, in.PageSize).Scan(&list)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}

	for _, v := range list {
		res.List = append(res.List, &entity.GenTable{
			TableName:    gconv.String(v["TABLE_NAME"]),
			TableComment: gconv.String(v["TABLE_COMMENT"]),
		})
	}
	return
}

// ImportTable 导入表结构
func (s *generator) ImportTable(ctx context.Context, names []string) error {
	if len(names) == 0 {
		return gerror.New("至少选择一张表")
	}

	config := g.Cfg().MustGet(ctx, "database.default.link").String()
	var dbName string
	if idx := strings.LastIndex(config, "/"); idx > 0 {
		dbName = strings.TrimSpace(config[idx+1:])
		if idx2 := strings.Index(dbName, "?"); idx2 > 0 {
			dbName = dbName[:idx2]
		}
	}
	if dbName == "" {
		return gerror.New("无法获取数据库名称，请检查 database.default.link 配置")
	}

	userId := gftoken.GetSessionUser(ctx).Id
	now := gtime.Now()

	for _, tableName := range names {
		// 获取表信息
		var tableInfo g.Map
		err := g.DB().Model("information_schema.tables").
			Where("table_schema", dbName).
			Where("table_name", tableName).
			Scan(&tableInfo)
		if err != nil {
			return err
		}
		if len(tableInfo) == 0 {
			return gerror.New(fmt.Sprintf("表 %s 不存在", tableName))
		}

		// 插入 gen_table
		table := &do.GenTable{
			TableName:    tableName,
			TableComment: gconv.String(tableInfo["TABLE_COMMENT"]),
			ClassName:    s.toClassName(tableName),
			PackageName:  "gmanager",
			ModuleName:   "system",
			BusinessName: s.toBusinessName(tableName),
			FunctionName: gconv.String(tableInfo["TABLE_COMMENT"]),
			TplCategory:  "crud",
			GenType:      "0",
			GenPath:      "",
			CreateBy:     userId,
			CreateAt:     now,
			UpdateBy:     userId,
			UpdateAt:     now,
		}
		tableId, err := dao.GenTable.Ctx(ctx).InsertAndGetId(table)
		if err != nil {
			return err
		}

		// 获取列信息
		var columns []g.Map
		err = g.DB().Model("information_schema.columns").
			Where("table_schema", dbName).
			Where("table_name", tableName).
			Order("ORDINAL_POSITION").
			Scan(&columns)
		if err != nil {
			return err
		}

		// 查询主键
		var pks []g.Map
		_ = g.DB().Model("information_schema.key_column_usage").
			Where("table_schema", dbName).
			Where("table_name", tableName).
			Where("constraint_name", "PRIMARY").
			Scan(&pks)
		pkMap := make(map[string]bool)
		for _, pk := range pks {
			pkMap[gconv.String(pk["COLUMN_NAME"])] = true
		}

		for _, col := range columns {
			columnName := gconv.String(col["COLUMN_NAME"])
			columnType := gconv.String(col["DATA_TYPE"])
			columnComment := gconv.String(col["COLUMN_COMMENT"])
			if columnComment == "" {
				columnComment = columnName
			}
			isPk := "0"
			if pkMap[columnName] {
				isPk = "1"
			}
			isIncrement := "0"
			if gconv.String(col["EXTRA"]) == "auto_increment" {
				isIncrement = "1"
			}

			goType := s.sqlTypeToGoType(columnType, gconv.String(col["COLUMN_TYPE"]))
			htmlType := s.inferHtmlType(columnName, columnType, gconv.String(col["COLUMN_TYPE"]))

			// 默认规则
			isInsert := "1"
			isEdit := "1"
			isList := "1"
			isQuery := "0"
			if isPk == "1" || s.isAuditColumn(columnName) {
				isInsert = "0"
				isEdit = "0"
				isList = "0"
			}
			if columnName == "status" || columnName == "enable" || gstr.HasSuffix(columnName, "_status") || gstr.HasSuffix(columnName, "_type") {
				isQuery = "1"
			}

			column := &do.GenTableColumn{
				TableId:       tableId,
				ColumnName:    columnName,
				ColumnComment: columnComment,
				ColumnType:    gconv.String(col["COLUMN_TYPE"]),
				GoType:        goType,
				GoField:       s.toCamelCase(columnName),
				IsPk:          isPk,
				IsIncrement:   isIncrement,
				IsRequired:    "0",
				IsInsert:      isInsert,
				IsEdit:        isEdit,
				IsList:        isList,
				IsQuery:       isQuery,
				QueryType:     s.inferQueryType(htmlType, goType),
				HtmlType:      htmlType,
				DictType:      "",
				Sort:          gconv.Int(col["ORDINAL_POSITION"]),
				CreateBy:      userId,
				CreateAt:      now,
				UpdateBy:      userId,
				UpdateAt:      now,
			}
			_, err = dao.GenTableColumn.Ctx(ctx).Insert(column)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Get 获取表配置详情
func (s *generator) Get(ctx context.Context, id int64) (res *v1.GeneratorTableGetRes, err error) {
	res = &v1.GeneratorTableGetRes{}
	err = dao.GenTable.Ctx(ctx).Where(dao.GenTable.Columns().Id, id).Scan(&res.Info)
	if err != nil {
		return
	}
	err = dao.GenTableColumn.Ctx(ctx).Where(dao.GenTableColumn.Columns().TableId, id).Order("sort").Scan(&res.Columns)
	return
}

// Save 保存表配置
func (s *generator) Save(ctx context.Context, in *v1.GeneratorTableSaveReq) error {
	userId := gftoken.GetSessionUser(ctx).Id
	now := gtime.Now()

	var table do.GenTable
	err := gconv.Struct(in, &table)
	if err != nil {
		return gerror.Wrap(err, "数据转换错误")
	}
	table.UpdateBy = userId
	table.UpdateAt = now

	if in.Id > 0 {
		_, err = dao.GenTable.Ctx(ctx).Where(dao.GenTable.Columns().Id, in.Id).Update(table)
		if err != nil {
			return err
		}
		// 删除旧字段
		_, err = dao.GenTableColumn.Ctx(ctx).Where(dao.GenTableColumn.Columns().TableId, in.Id).Delete()
		if err != nil {
			return err
		}
	} else {
		table.CreateBy = userId
		table.CreateAt = now
		id, err := dao.GenTable.Ctx(ctx).InsertAndGetId(table)
		if err != nil {
			return err
		}
		in.Id = id
	}

	// 插入新字段
	for _, col := range in.Columns {
		column := &do.GenTableColumn{}
		_ = gconv.Struct(col, column)
		column.TableId = in.Id
		column.CreateBy = userId
		column.CreateAt = now
		column.UpdateBy = userId
		column.UpdateAt = now
		_, err = dao.GenTableColumn.Ctx(ctx).Insert(column)
		if err != nil {
			return err
		}
	}
	return nil
}

// Delete 删除表配置
func (s *generator) Delete(ctx context.Context, ids []int) error {
	_, err := dao.GenTableColumn.Ctx(ctx).WhereIn(dao.GenTableColumn.Columns().TableId, ids).Delete()
	if err != nil {
		return err
	}
	_, err = dao.GenTable.Ctx(ctx).WhereIn(dao.GenTable.Columns().Id, ids).Delete()
	return err
}

// ==================== 代码生成核心 ====================

// Preview 预览代码
func (s *generator) Preview(ctx context.Context, id int64) (res *v1.GeneratorPreviewRes, err error) {
	table, columns, err := s.getGenData(ctx, id)
	if err != nil {
		return
	}
	data, err := s.renderAll(table, columns)
	if err != nil {
		return
	}
	res = &v1.GeneratorPreviewRes{Data: data}
	return
}

// Download 下载 ZIP
func (s *generator) Download(ctx context.Context, id int64) (*bytes.Buffer, string, error) {
	table, columns, err := s.getGenData(ctx, id)
	if err != nil {
		return nil, "", err
	}
	data, err := s.renderAll(table, columns)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)
	for fileName, content := range data {
		w, err := zw.Create(fileName)
		if err != nil {
			return nil, "", err
		}
		_, _ = w.Write([]byte(content))
	}
	_ = zw.Close()

	zipName := table.BusinessName + ".zip"
	return buf, zipName, nil
}

// GenCode 写入磁盘
func (s *generator) GenCode(ctx context.Context, id int64, path string) error {
	table, columns, err := s.getGenData(ctx, id)
	if err != nil {
		return err
	}
	data, err := s.renderAll(table, columns)
	if err != nil {
		return err
	}

	genPath := path
	if genPath == "" {
		genPath = table.GenPath
	}
	if genPath == "" {
		// 默认项目根目录
		genPath = "."
	}

	// 安全检查：确保路径在项目根目录内
	absPath, err := filepath.Abs(genPath)
	if err != nil {
		return err
	}
	// 获取工作目录
	wd, err := os.Getwd()
	if err != nil {
		wd = "."
	}
	absWd, _ := filepath.Abs(wd)
	if !strings.HasPrefix(absPath, absWd) {
		return gerror.New("生成路径必须在项目目录内")
	}

	for fileName, content := range data {
		fullPath := filepath.Join(absPath, fileName)
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			return err
		}
	}
	return nil
}

// ==================== 辅助方法 ====================

func (s *generator) getGenData(ctx context.Context, id int64) (*entity.GenTable, []*entity.GenTableColumn, error) {
	var table *entity.GenTable
	err := dao.GenTable.Ctx(ctx).Where(dao.GenTable.Columns().Id, id).Scan(&table)
	if err != nil {
		return nil, nil, err
	}
	if table == nil {
		return nil, nil, gerror.New("表配置不存在")
	}
	var columns []*entity.GenTableColumn
	err = dao.GenTableColumn.Ctx(ctx).Where(dao.GenTableColumn.Columns().TableId, id).Order("sort").Scan(&columns)
	if err != nil {
		return nil, nil, err
	}
	return table, columns, nil
}

func (s *generator) renderAll(table *entity.GenTable, columns []*entity.GenTableColumn) (map[string]string, error) {
	result := make(map[string]string)
	mp := s.buildTemplateData(table, columns)

	files := map[string]string{
		"go/api.go.tpl":        fmt.Sprintf("server/api/admin/v1/%s.go", table.BusinessName),
		"go/controller.go.tpl": fmt.Sprintf("server/internal/admin/controller/%s.go", table.BusinessName),
		"go/logic.go.tpl":      fmt.Sprintf("server/internal/admin/logic/%s.go", table.BusinessName),
		"vue/index.vue.tpl":    fmt.Sprintf("web/src/views/%s/%s/index.vue", table.ModuleName, table.BusinessName),
		"vue/api.ts.tpl":       fmt.Sprintf("web/src/api/%s/%s.api.ts", table.ModuleName, table.BusinessName),
	}

	for tplName, outPath := range files {
		content, err := s.renderTemplate(tplName, mp)
		if err != nil {
			return nil, err
		}
		result[outPath] = content
	}
	return result, nil
}

func (s *generator) renderTemplate(tplName string, data map[string]interface{}) (string, error) {
	// 尝试从嵌入资源或文件读取模板
	// 开发时从文件读取，构建后从 embed 读取
	// 这里先实现文件读取方式（因为当前是开发环境）
	tplPath := "server/resource/template/generator/" + tplName
	content, err := os.ReadFile(tplPath)
	if err != nil {
		// 尝试从工作目录的相对路径读取
		content, err = os.ReadFile(tplPath)
		if err != nil {
			return "", gerror.Wrap(err, "读取模板失败: "+tplName)
		}
	}

	tplContent := string(content)
	// 对 Vue 模板，先将 Vue 的 {{ }} 占位符保护起来，避免与 Go template 冲突
	isVue := strings.HasPrefix(tplName, "vue/")
	if isVue {
		tplContent = strings.ReplaceAll(tplContent, "{{", "__VUE_OPEN__")
		tplContent = strings.ReplaceAll(tplContent, "}}", "__VUE_CLOSE__")
	}

	tpl, err := template.New(tplName).Funcs(s.templateFuncs()).Parse(tplContent)
	if err != nil {
		return "", gerror.Wrap(err, "解析模板失败: "+tplName)
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return "", gerror.Wrap(err, "渲染模板失败: "+tplName)
	}

	result := buf.String()
	if isVue {
		result = strings.ReplaceAll(result, "__VUE_OPEN__", "{{")
		result = strings.ReplaceAll(result, "__VUE_CLOSE__", "}}")
	}
	return result, nil
}

func (s *generator) buildTemplateData(table *entity.GenTable, columns []*entity.GenTableColumn) map[string]interface{} {
	// 过滤字段
	var queryColumns, listColumns, formColumns []*entity.GenTableColumn
	var pkColumn *entity.GenTableColumn
	for _, col := range columns {
		if col.IsPk == "1" {
			pkColumn = col
		}
		if col.IsQuery == "1" {
			queryColumns = append(queryColumns, col)
		}
		if col.IsList == "1" {
			listColumns = append(listColumns, col)
		}
		if col.IsEdit == "1" || col.IsInsert == "1" {
			formColumns = append(formColumns, col)
		}
	}

	return map[string]interface{}{
		"table":        table,
		"columns":      columns,
		"pkColumn":     pkColumn,
		"queryColumns": queryColumns,
		"listColumns":  listColumns,
		"formColumns":  formColumns,
		"className":    table.ClassName,
		"businessName": table.BusinessName,
		"moduleName":   table.ModuleName,
		"functionName": table.FunctionName,
		"author":       table.FunctionAuthor,
	}
}

func (s *generator) templateFuncs() template.FuncMap {
	return template.FuncMap{
		"camelCase":   s.toCamelCase,
		"lowerFirst":  s.lowerFirst,
		"upperFirst":  s.upperFirst,
		"goTag":       s.goTag,
		"vueField":    s.vueField,
		"queryMethod": s.queryMethod,
	}
}

// ==================== 类型转换工具 ====================

func (s *generator) toClassName(tableName string) string {
	name := s.toCamelCase(tableName)
	return s.upperFirst(name)
}

func (s *generator) toBusinessName(tableName string) string {
	// sys_user -> user
	parts := strings.Split(tableName, "_")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return tableName
}

func (s *generator) toCamelCase(name string) string {
	parts := strings.Split(name, "_")
	for i := 0; i < len(parts); i++ {
		parts[i] = s.upperFirst(parts[i])
	}
	return strings.Join(parts, "")
}

func (s *generator) upperFirst(s1 string) string {
	if s1 == "" {
		return ""
	}
	return strings.ToUpper(s1[:1]) + s1[1:]
}

func (s *generator) lowerFirst(s1 string) string {
	if s1 == "" {
		return ""
	}
	return strings.ToLower(s1[:1]) + s1[1:]
}

func (s *generator) sqlTypeToGoType(dataType, columnType string) string {
	switch dataType {
	case "tinyint", "smallint", "mediumint", "int":
		return "int"
	case "bigint":
		return "int64"
	case "float":
		return "float32"
	case "double", "decimal":
		return "float64"
	case "datetime", "timestamp", "date", "time":
		return "*gtime.Time"
	default:
		return "string"
	}
}

func (s *generator) inferHtmlType(columnName, dataType, columnType string) string {
	if dataType == "datetime" || dataType == "timestamp" {
		return "datetime"
	}
	if dataType == "tinyint" && strings.Contains(columnType, "(1)") {
		return "switch"
	}
	if strings.HasSuffix(columnName, "_status") || strings.HasSuffix(columnName, "_type") {
		return "select"
	}
	if dataType == "text" || dataType == "longtext" {
		return "textarea"
	}
	return "input"
}

func (s *generator) inferQueryType(htmlType, goType string) string {
	if htmlType == "datetime" {
		return "BETWEEN"
	}
	if goType == "string" {
		return "LIKE"
	}
	return "EQ"
}

func (s *generator) isAuditColumn(columnName string) bool {
	switch columnName {
	case "create_at", "update_at", "create_by", "update_by", "create_id", "update_id",
		"create_time", "update_time", "created_at", "updated_at":
		return true
	}
	return false
}

func (s *generator) goTag(col *entity.GenTableColumn) string {
	tag := fmt.Sprintf(`json:"%s"`, col.GoField)
	if col.IsRequired == "1" {
		tag += fmt.Sprintf(` v:"required#%s不能为空"`, col.ColumnComment)
	}
	return tag
}

func (s *generator) vueField(col *entity.GenTableColumn) string {
	return col.GoField
}

func (s *generator) queryMethod(col *entity.GenTableColumn) string {
	switch col.QueryType {
	case "EQ":
		return "Where"
	case "NE":
		return "WhereNot"
	case "GT":
		return "WhereGT"
	case "LT":
		return "WhereLT"
	case "LIKE":
		return "WhereLike"
	case "BETWEEN":
		return "WhereBetween"
	default:
		return "Where"
	}
}
