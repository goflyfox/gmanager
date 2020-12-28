package role

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"gmanager/app/service/role"
	"gmanager/library/base"
)

type Action struct {
	base.BaseRouter
}

// path: /index
func (action *Action) Index(r *ghttp.Request) {
	tplFile := "pages/system/role_index.html"
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
	model, err := role.GetById(id)
	if err != nil {
		base.Fail(r, err.Error())
	}

	base.Succ(r, model)
}

// path: /delete/{id}
func (action *Action) Delete(r *ghttp.Request) {
	id := r.GetInt64("id")
	_, err := role.Delete(id, base.GetUser(r).Id)
	if err != nil {
		base.Fail(r, err.Error())
	}

	base.Succ(r, "")
}

// path: /save
func (action *Action) Save(r *ghttp.Request) {
	request := new(role.Request)
	err := gconv.Struct(r.GetMap(), request)
	if err != nil {
		glog.Error("save struct error", err)
		base.Error(r, "save error")
	}

	menus := r.GetString("menus")

	request.UserId = base.GetUser(r).Id
	_, err = role.Save(request)
	if err != nil {
		base.Fail(r, "保存失败")
	}

	// 保存菜单信息
	role.SaveRoleMenu(request.Id, menus)

	base.Succ(r, "")
}

// path: /list
func (action *Action) List(r *ghttp.Request) {
	form := base.NewForm(r.GetMap())

	list, err := role.List(&form)
	if err != nil {
		glog.Error("page error", err)
		base.Error(r, err.Error())
	}

	base.Succ(r, list)
}

// path: /page
func (action *Action) Page(r *ghttp.Request) {
	form := base.NewForm(r.GetMap())
	page, err := role.Page(&form)
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
	page, err := role.Page(&form)
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

// 角色绑定菜单信息
func (action *Action) Info(r *ghttp.Request) {
	roleId := r.GetInt("roleId")
	if roleId == 0 {
		base.Fail(r, "参数错误")
	}
	form := base.NewForm(r.GetMap())
	form.SetParam("roleId", gconv.String(roleId))

	// 选择的列表
	roleMenuList, err := role.ListRoleMenu(&form)
	if err != nil {
		glog.Error("info error", err)
		base.Error(r, err.Error())
	}

	menuSlice := g.SliceStr{}
	for _, roleMenu := range roleMenuList {
		menuSlice = append(menuSlice, gconv.String(roleMenu.MenuId))
	}

	base.Succ(r, g.Map{
		"roleId": roleId,
		"menus":  menuSlice,
	})
}

// 角色绑定菜单
func (action *Action) MenuSave(r *ghttp.Request) {
	roleId := r.GetInt("roleId")
	menuIds := r.GetString("menuIds")
	if roleId == 0 || menuIds == "" {
		base.Fail(r, "参数错误")
	}

	err := role.SaveRoleMenu(roleId, menuIds)
	if err != nil {
		glog.Error("MenuSave error", err)
		base.Error(r, err.Error())
	}

	base.Succ(r, "")
}
