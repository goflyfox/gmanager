package system

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/gogf/gf/g/util/gconv"
	"gmanager/utils"
	"gmanager/utils/base"
)

type LogAction struct {
	base.BaseRouter
}

var (
	actionNameLog = "LogAction"
)

// path: /index
func (action *LogAction) Index(r *ghttp.Request) {
	tplFile := "pages/system/log_index.html"
	err := r.Response.WriteTpl(tplFile, g.Map{
		"now": gtime.Datetime(),
	})

	if err != nil {
		glog.Error(err)
	}
}

// path: /get/{id}
func (action *LogAction) Get(r *ghttp.Request) {
	id := r.GetInt("id")
	model := SysLog{Id: id}.Get()
	if model.Id <= 0 {
		base.Fail(r, actionNameLog+" get fail")
	}

	base.Succ(r, model)
}

// path: /delete/{id}
func (action *LogAction) Delete(r *ghttp.Request) {
	id := r.GetInt("id")

	model := SysLog{Id: id}
	model.UpdateId = base.GetUser(r).Id
	model.UpdateTime = utils.GetNow()

	num := model.Delete()
	if num <= 0 {
		base.Fail(r, actionNameLog+" delete fail")
	}

	base.Succ(r, "")
}

// path: /save
func (action *LogAction) Save(r *ghttp.Request) {
	model := SysLog{}
	err := gconv.Struct(r.GetPostMap(), &model)
	if err != nil {
		glog.Error(actionNameLog+" save struct error", err)
		base.Error(r, "save error")
	}

	userId := base.GetUser(r).Id

	model.UpdateId = userId
	model.UpdateTime = utils.GetNow()

	var num int64
	if model.Id <= 0 {
		model.CreateId = userId
		model.CreateTime = utils.GetNow()
		num = model.Insert()
	} else {
		num = model.Update()
	}

	if num <= 0 {
		base.Fail(r, actionNameLog+" save fail")
	}

	base.Succ(r, "")
}

// path: /list
func (action *LogAction) List(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := SysLog{}

	list := model.List(&form)
	base.Succ(r, list)
}

// path: /page
func (action *LogAction) Page(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := SysLog{}

	page := model.Page(&form)
	base.Succ(r, g.Map{"list": page, "form": form})
}

// path: /jqgrid
func (action *LogAction) Jqgrid(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := SysLog{}

	page := model.Page(&form)
	r.Response.WriteJson(g.Map{
		"page":    form.Page,
		"rows":    page,
		"total":   form.TotalPage,
		"records": form.TotalSize,
	})
}
