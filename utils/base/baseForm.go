package base

import "github.com/gogf/gf/g/util/gconv"

type BaseForm struct {
	Page      int    `form:"page",json:"page"`           // 当前页码
	Rows      int    `form:"rows",json:"rows"`           // 每页多少条
	TotalPage int    `form:"totalPage",json:"totalPage"` // 总页数
	TotalSize int    `form:"totalSize",json:"totalSize"` // 总共数据条数
	OrderBy   string `form:"orderBy",json:"orderBy"`     // 排序
	Params    map[string]string
	Object    interface{}
}

func NewForm(params map[string]string) BaseForm {
	form := BaseForm{}
	form.Params = params
	//  第几页
	if value, ok := params["page"]; ok {
		form.Page = gconv.Int(value)
	}
	// 页数
	if value, ok := params["rows"]; ok {
		form.Rows = gconv.Int(value)
	}
	// 排序
	if value, ok := params["orderBy"]; ok {
		form.OrderBy = gconv.String(value)
	} else if value, ok := params["sidx"]; ok {
		form.OrderBy = gconv.String(value)
		if value2, ok := params["sord"]; ok {
			form.OrderBy += " " + gconv.String(value2)
		}
	}

	return form
}

func (form *BaseForm) SetParam(key string, value string) {
	form.Params[key] = value
}

func (form *BaseForm) SetParams(params map[string]string) {
	form.Page = gconv.Int(params["page"])
	form.Rows = gconv.Int(params["rows"])
	form.OrderBy = params["orderBy"]
	form.Params = params
}

func (form *BaseForm) SetObject(object interface{}) {
	form.Object = object
}
