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

type DepartmentAction struct {
	base.BaseRouter
}

var (
	actionNameDepartment = "DepartmentAction"
)

// path: /index
func (action *DepartmentAction) Index(r *ghttp.Request) {
	tplFile := "pages/system/department_index.html"
	err := r.Response.WriteTpl(tplFile, g.Map{
		"now": gtime.Datetime(),
	})

	if err != nil {
		glog.Error(err)
	}
}

// path: /get/{id}
func (action *DepartmentAction) Get(r *ghttp.Request) {
	id := r.GetInt("id")
	model := SysDepartment{Id: id}.Get()
	if model.Id <= 0 {
		base.Fail(r, actionNameDepartment+" get fail")
	}

	base.Succ(r, model)
}

// path: /delete/{id}
func (action *DepartmentAction) Delete(r *ghttp.Request) {
	id := r.GetInt("id")

	model := SysDepartment{Id: id}
	model.UpdateId = base.GetUser(r).Id
	model.UpdateTime = utils.GetNow()

	num := model.Delete()
	if num <= 0 {
		base.Fail(r, actionNameDepartment+" delete fail")
	}

	base.Succ(r, "")
}

// path: /save
func (action *DepartmentAction) Save(r *ghttp.Request) {
	model := SysDepartment{}
	err := gconv.Struct(r.GetPostMap(), &model)
	if err != nil {
		glog.Error(actionNameDepartment+" save struct error", err)
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
		base.Fail(r, actionNameDepartment+" save fail")
	}

	base.Succ(r, "")
}

// path: /list
func (action *DepartmentAction) List(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := SysDepartment{}

	list := model.List(&form)
	base.Succ(r, list)
}

// path: /page
func (action *DepartmentAction) Page(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := SysDepartment{}

	page := model.Page(&form)
	base.Succ(r, g.Map{"list": page, "form": form})
}

// path: /jqgrid
func (action *DepartmentAction) Jqgrid(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := SysDepartment{}

	page := model.Page(&form)
	r.Response.WriteJson(g.Map{
		"page":    form.Page,
		"rows":    page,
		"total":   form.TotalPage,
		"records": form.TotalSize,
	})
}
