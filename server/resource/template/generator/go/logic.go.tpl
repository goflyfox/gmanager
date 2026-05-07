package logic

import (
	"context"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/admin/dao"
	"gmanager/internal/admin/model/do"
	"gmanager/internal/admin/model/entity"
	"gmanager/internal/library/gftoken"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// {{.className}} {{.functionName}}服务
var {{.className}} = new({{.businessName}})

type {{.businessName}} struct{}

// List 获取列表
func (s *{{.businessName}}) List(ctx context.Context, in *v1.{{.className}}ListReq) (res *v1.{{.className}}ListRes, err error) {
	if in == nil {
		return
	}
	m := dao.{{.className}}.Ctx(ctx)
	columns := dao.{{.className}}.Columns()
	res = &v1.{{.className}}ListRes{}

	if in.Keywords != "" {
		m = m.Where(m.Builder(){{range $i, $col := .columns}}{{if eq $col.IsQuery "1"}}{{if eq $col.QueryType "LIKE"}}.
			WhereOrLike(columns.{{$col.GoField}}, "%"+in.Keywords+"%"){{end}}{{end}}{{end}})
	}
{{- range .queryColumns}}
{{- if eq .QueryType "EQ"}}
	if in.{{.GoField}} != "" {
		m = m.Where(columns.{{.GoField}}, in.{{.GoField}})
	}
{{- else if eq .QueryType "LIKE"}}
	if in.{{.GoField}} != "" {
		m = m.WhereLike(columns.{{.GoField}}, "%"+in.{{.GoField}}+"%")
	}
{{- else if eq .QueryType "BETWEEN"}}
	if len(in.{{.GoField}}) > 0 && in.{{.GoField}}[0] != "" {
		m = m.WhereBetween(columns.{{.GoField}}, in.{{.GoField}}[0]+" 00:00:00", in.{{.GoField}}[1]+" 23:59:59")
	}
{{- end}}
{{- end}}

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
	var pageList []*entity.{{.className}}
	if err = m.Page(in.PageNum, in.PageSize).Scan(&pageList); err != nil {
		err = gerror.Wrap(err, "获取数据失败！")
		return
	}
	res.List = pageList
	return
}

// Get 获取详情
func (s *{{.businessName}}) Get(ctx context.Context, id int64) (res *v1.{{.className}}GetRes, err error) {
	err = dao.{{.className}}.Ctx(ctx).Where(dao.{{.className}}.Columns().Id, id).Scan(&res)
	return
}

// Save 保存
func (s *{{.businessName}}) Save(ctx context.Context, in *v1.{{.className}}SaveReq) error {
	var model do.{{.className}}
	err := gconv.Struct(in, &model)
	if err != nil {
		return gerror.Wrap(err, "数据转换错误")
	}

	m := dao.{{.className}}.Ctx(ctx)
	columns := dao.{{.className}}.Columns()

	userId := gftoken.GetSessionUser(ctx).Id
	model.UpdateId = userId
	model.UpdateAt = gtime.Now()
	if in.Id > 0 {
		_, err = m.Where(columns.Id, model.Id).Update(model)
		if err != nil {
			return err
		}
	} else {
		model.CreateId = userId
		model.CreateAt = gtime.Now()
		_, err = m.Insert(model)
		if err != nil {
			return err
		}
	}
	return nil
}

// Delete 删除
func (s *{{.businessName}}) Delete(ctx context.Context, ids []int) error {
	_, err := dao.{{.className}}.Ctx(ctx).WhereIn(dao.{{.className}}.Columns().Id, ids).Delete()
	return err
}
