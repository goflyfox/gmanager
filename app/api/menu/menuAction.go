package menu

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"gmanager/app/service/menu"
	"gmanager/library/base"
)

type Action struct {
	base.BaseRouter
}

// path: /index
func (action *Action) Index(r *ghttp.Request) {
	tplFile := "pages/system/menu_index.html"
	err := r.Response.WriteTpl(tplFile, g.Map{
		"now": gtime.Datetime(),
	})

	if err != nil {
		glog.Error(err)
	}
}

// path: /get/{id}
func (action *Action) Get(r *ghttp.Request) {
	id := r.GetInt64("id")
	model, err := menu.GetById(id)
	if err != nil {
		base.Fail(r, err.Error())
	}

	base.Succ(r, model)
}

// path: /delete/{id}
func (action *Action) Delete(r *ghttp.Request) {
	id := r.GetInt64("id")

	form := base.NewForm(g.Map{"parentId": id})
	childModel, err := menu.GetOne(&form)
	if err != nil {
		base.Fail(r, err.Error())
	} else if childModel != nil && childModel.Id > 0 {
		base.Fail(r, "请先删除子菜单")
	}

	_, err1 := menu.Delete(id, base.GetUser(r).Id)
	if err1 != nil {
		base.Fail(r, err1.Error())
	}

	base.Succ(r, "")
}

// path: /save
func (action *Action) Save(r *ghttp.Request) {
	request := new(menu.Request)
	err := gconv.Struct(r.GetMap(), request)
	if err != nil {
		glog.Error("save struct error", err)
		base.Error(r, "save error")
	}

	request.UserId = base.GetUser(r).Id
	_, err = menu.Save(request)
	if err != nil {
		base.Fail(r, "保存失败")
	}

	base.Succ(r, "")
}

// path: /tree
func (action *Action) Tree(r *ghttp.Request) {
	action.List(r)
}

// path: /list
func (action *Action) List(r *ghttp.Request) {
	form := base.NewForm(r.GetMap())

	list, err := menu.List(&form)
	if err != nil {
		glog.Error("page error", err)
		base.Error(r, err.Error())
	}

	base.Succ(r, list)
}

// path: /page
func (action *Action) Page(r *ghttp.Request) {
	form := base.NewForm(r.GetMap())
	page, err := menu.Page(&form)
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
func (action *Action) Jqgrid(r *ghttp.Request) {
	form := base.NewForm(r.GetMap())
	page, err := menu.Page(&form)
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
