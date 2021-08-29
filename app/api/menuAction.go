package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"gmanager/app/service"
	"gmanager/library/base"
)

var Menu = menuApi{}

type menuApi struct{ base.BaseRouter }

// path: /index
func (action *menuApi) Index(r *ghttp.Request) {
	tplFile := "pages/system/menu_index.html"
	err := r.Response.WriteTpl(tplFile, g.Map{
		"now": gtime.Datetime(),
	})

	if err != nil {
		glog.Error(err)
	}
}

// path: /get/{id}
func (action *menuApi) Get(r *ghttp.Request) {
	id := r.GetInt64("id")
	model, err := service.Menu.GetById(r.Context(), id)
	if err != nil {
		base.Fail(r, err.Error())
	}

	base.Succ(r, model)
}

// path: /delete/{id}
func (action *menuApi) Delete(r *ghttp.Request) {
	id := r.GetInt64("id")

	err := service.Menu.Delete(r.Context(), id, base.GetUser(r).Id)
	if err != nil {
		base.Fail(r, err.Error())
	}

	base.Succ(r, "")
}

// path: /save
func (action *menuApi) Save(r *ghttp.Request) {
	request := new(service.MenuReq)
	err := gconv.Struct(r.GetMap(), request)
	if err != nil {
		glog.Error("save struct error", err)
		base.Error(r, "save error")
	}

	request.UserId = base.GetUser(r).Id
	_, err = service.Menu.Save(r.Context(), request)
	if err != nil {
		base.Fail(r, "保存失败")
	}

	base.Succ(r, "")
}

// path: /list
func (action *menuApi) List(r *ghttp.Request) {
	form := base.NewForm(r.GetMap())

	list, err := service.Menu.List(r.Context(), &form)
	if err != nil {
		glog.Error("page error", err)
		base.Error(r, err.Error())
	}

	base.Succ(r, list)
}

// path: /page
func (action *menuApi) Page(r *ghttp.Request) {
	form := base.NewForm(r.GetMap())
	page, err := service.Menu.Page(r.Context(), &form)
	if err != nil {
		glog.Error("page error", err)
		base.Error(r, err.Error())
	}

	base.Succ(r,
		g.Map{
			"page":    form.Page,
			"rows":    page,
			"total":   form.TotalPage,
			"records": form.TotalSize,
		})
}

// path: /jqgrid
func (action *menuApi) Jqgrid(r *ghttp.Request) {
	form := base.NewForm(r.GetMap())
	page, err := service.Menu.Page(r.Context(), &form)
	if err != nil {
		glog.Error("jqgrid error", err)
		base.Error(r, err.Error())
	}

	r.Response.WriteJson(g.Map{
		"page":    form.Page,
		"rows":    page,
		"total":   form.TotalPage,
		"records": form.TotalSize,
	})
}
