package system

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/gogf/gf/g/text/gstr"
	"github.com/gogf/gf/g/util/gconv"
	"gmanager/utils"
	"gmanager/utils/base"
)

type RoleAction struct {
	base.BaseRouter
}

var (
	actionNameRole = "RoleAction"
)

// path: /index
func (action *RoleAction) Index(r *ghttp.Request) {
	tplFile := "pages/system/role_index.html"
	err := r.Response.WriteTpl(tplFile, g.Map{
		"now": gtime.Datetime(),
	})

	if err != nil {
		glog.Error(err)
	}
}

// path: /get/{id}
func (action *RoleAction) Get(r *ghttp.Request) {
	id := r.GetInt("id")
	model := SysRole{Id: id}.Get()
	if model.Id <= 0 {
		base.Fail(r, actionNameRole+" get fail")
	}

	base.Succ(r, model)
}

// path: /delete/{id}
func (action *RoleAction) Delete(r *ghttp.Request) {
	id := r.GetInt("id")

	model := SysRole{Id: id}
	model.UpdateId = base.GetUser(r).Id
	model.UpdateTime = utils.GetNow()

	num := model.Delete()
	if num <= 0 {
		base.Fail(r, actionNameRole+" delete fail")
	}

	base.Succ(r, "")
}

// path: /save
func (action *RoleAction) Save(r *ghttp.Request) {
	model := SysRole{}
	err := gconv.Struct(r.GetPostMap(), &model)
	if err != nil {
		glog.Error(actionNameRole+" save struct error", err)
		base.Error(r, "save error")
	}

	menus := r.GetPostString("menus")
	if menus == "" {
		base.Fail(r, "参数错误")
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
		base.Fail(r, actionNameRole+" save fail")
	}

	// 保存菜单信息
	SysRoleMenu{RoleId: model.Id}.DeleteByRoleId()
	menuIdArray := gstr.Split(menus, ",")
	for _, menuId := range menuIdArray {
		model := SysRoleMenu{RoleId: model.Id, MenuId: gconv.Int(menuId)}
		model.Insert()
	}

	base.Succ(r, "")
}

// path: /list
func (action *RoleAction) List(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := SysRole{}

	list := model.List(&form)
	base.Succ(r, list)
}

// path: /page
func (action *RoleAction) Page(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := SysRole{}

	page := model.Page(&form)
	base.Succ(r, g.Map{"list": page, "form": form})
}

// path: /jqgrid
func (action *RoleAction) Jqgrid(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := SysRole{}

	page := model.Page(&form)
	r.Response.WriteJson(g.Map{
		"page":    form.Page,
		"rows":    page,
		"total":   form.TotalPage,
		"records": form.TotalSize,
	})
}

// 角色绑定菜单信息
func (action *RoleAction) Info(r *ghttp.Request) {
	roleId := r.GetPostInt("roleId")
	if roleId == 0 {
		base.Fail(r, "参数错误")
	}
	form := base.NewForm(r.GetPostMap())
	form.SetParam("roleId", gconv.String(roleId))

	// 选择的列表
	roleMenuList := SysRoleMenu{}.List(&form)
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
func (action *RoleAction) MenuSave(r *ghttp.Request) {
	roleId := r.GetPostInt("roleId")
	menuIds := r.GetPostString("menuIds")
	if roleId == 0 || menuIds == "" {
		base.Fail(r, "参数错误")
	}

	SysRoleMenu{RoleId: roleId}.DeleteByRoleId()
	menuIdArray := gstr.Split(menuIds, ",")
	for _, menuId := range menuIdArray {
		model := SysRoleMenu{RoleId: roleId, MenuId: gconv.Int(menuId)}
		model.Insert()
	}

	base.Succ(r, "")
}
