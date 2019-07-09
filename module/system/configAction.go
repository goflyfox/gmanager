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

type ConfigAction struct {
	base.BaseRouter
}

var (
	actionNameConfig = "ConfigAction"
)

// path: /index
func (action *ConfigAction) Index(r *ghttp.Request) {
	tplFile := "pages/system/config_index.html"
	err := r.Response.WriteTpl(tplFile, g.Map{
		"now": gtime.Datetime(),
	})

	if err != nil {
		glog.Error(err)
	}
}

// path: /get/{id}
func (action *ConfigAction) Get(r *ghttp.Request) {
	id := r.GetInt("id")
	model := SysConfig{Id: id}.Get()
	if model.Id <= 0 {
		base.Fail(r, actionNameConfig+" get fail")
	}

	base.Succ(r, model)
}

// path: /delete/{id}
func (action *ConfigAction) Delete(r *ghttp.Request) {
	id := r.GetInt("id")

	model := SysConfig{Id: id}
	model.UpdateId = base.GetUser(r).Id
	model.UpdateTime = utils.GetNow()

	num := model.Delete()
	if num <= 0 {
		base.Fail(r, actionNameConfig+" delete fail")
	}

	base.Succ(r, "")
}

// path: /save
func (action *ConfigAction) Save(r *ghttp.Request) {
	model := SysConfig{}
	err := gconv.Struct(r.GetPostMap(), &model)
	if err != nil {
		glog.Error(actionNameConfig+" save struct error", err)
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
		base.Fail(r, actionNameConfig+" save fail")
	}

	base.Succ(r, "")
}

// path: /list
func (action *ConfigAction) List(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := SysConfig{}

	list := model.List(&form)
	base.Succ(r, list)
}

// path: /page
func (action *ConfigAction) Page(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := SysConfig{}

	page := model.Page(&form)
	base.Succ(r, g.Map{"list": page, "form": form})
}

// path: /jqgrid
func (action *ConfigAction) Jqgrid(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := SysConfig{}

	page := model.Page(&form)
	r.Response.WriteJson(g.Map{
		"page":    form.Page,
		"rows":    page,
		"total":   form.TotalPage,
		"records": form.TotalSize,
	})
}

// path: /type
func (action *ConfigAction) Type(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := SysConfig{}

	//userId := base.GetUser(r).Id
	//user := SysUser{Id: userId}.Get()
	form.SetParam("parentId", "0")
	form.OrderBy = "sort asc,create_time desc"

	list := model.List(&form)
	base.Succ(r, list)
}
