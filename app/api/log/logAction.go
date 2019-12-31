package log

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	logM "gmanager/app/model/log"
	"gmanager/app/service/log"
	"gmanager/utils"
	"gmanager/utils/base"
)

type Action struct {
	base.BaseRouter
}

// path: /index
func (action *Action) Index(r *ghttp.Request) {
	tplFile := "pages/system/log_index.html"
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
	model, err := log.GetById(id)
	if err != nil {
		base.Fail(r, err.Error())
	}

	base.Succ(r, model)
}

// path: /delete/{id}
func (action *Action) Delete(r *ghttp.Request) {
	id := r.GetInt64("id")
	_, err := log.Delete(id)
	if err != nil {
		base.Fail(r, err.Error())
	}

	base.Succ(r, "")
}

// path: /save
func (action *Action) Save(r *ghttp.Request) {
	model := logM.Entity{}
	err := gconv.Struct(r.GetPostMap(), &model)
	if err != nil {
		glog.Error("save struct error", err)
		base.Error(r, "save error")
	}

	userId := base.GetUser(r).Id

	model.UpdateId = userId
	model.UpdateTime = utils.GetNow()

	if model.Id <= 0 {
		model.CreateId = userId
		model.CreateTime = utils.GetNow()
		_, err = log.Insert(&model)
	} else {
		_, err = log.Update(&model)
	}

	if err != nil {
		base.Fail(r, "save fail")
	}

	base.Succ(r, "")
}

// path: /list
func (action *Action) List(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())

	list, err := log.List(&form)
	if err != nil {
		glog.Error("page error", err)
		base.Error(r, err.Error())
	}

	base.Succ(r, list)
}

// path: /page
func (action *Action) Page(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	page, err := log.Page(&form)
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
	form := base.NewForm(r.GetPostMap())
	page, err := log.Page(&form)
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
