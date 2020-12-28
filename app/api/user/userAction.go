package user

import (
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"gmanager/app/service/menu"
	"gmanager/app/service/role"
	"gmanager/app/service/user"
	"gmanager/library"
	"gmanager/library/base"
)

type Action struct {
	base.BaseRouter
}

// path: /index
func (action *Action) Index(r *ghttp.Request) {
	tplFile := "pages/system/user_index.html"
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
	model, err := user.GetById(id)
	if err != nil {
		base.Fail(r, err.Error())
	}

	base.Succ(r, model)
}

// path: /delete/{id}
func (action *Action) Delete(r *ghttp.Request) {
	id := r.GetInt64("id")
	_, err := user.Delete(id, base.GetUser(r).Id)
	if err != nil {
		base.Fail(r, err.Error())
	}

	base.Succ(r, "")
}

// path: /save
func (action *Action) Save(r *ghttp.Request) {
	request := new(user.Request)
	err := gconv.Struct(r.GetMap(), request)
	if err != nil {
		glog.Error("save struct error", err)
		base.Error(r, "save error")
	}

	request.UserId = base.GetUser(r).Id
	_, err = user.Save(request)
	if err != nil {
		base.Fail(r, "保存失败")
	}

	base.Succ(r, "")
}

// path: /list
func (action *Action) List(r *ghttp.Request) {
	form := base.NewForm(r.GetMap())

	list, err := user.List(&form)
	if err != nil {
		glog.Error("page error", err)
		base.Error(r, err.Error())
	}

	base.Succ(r, list)
}

// path: /page
func (action *Action) Page(r *ghttp.Request) {
	form := base.NewForm(r.GetMap())
	page, err := user.Page(&form)
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
	page, err := user.Page(&form)
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

// 获取用户角色信息
func (action *Action) RoleInfo(r *ghttp.Request) {
	userId := r.GetInt("userId")
	if userId == 0 {
		base.Fail(r, "参数错误")
	}
	form := base.NewForm(r.GetMap())
	form.SetParam("userId", gconv.String(userId))

	// 角色列表
	roleList, err := role.List(&form)
	if err != nil {
		glog.Error("RoleInfo error", err)
		base.Error(r, err.Error())
	}

	// 选择的列表
	userRoleList, err := user.ListUserRole(&form)
	if err != nil {
		glog.Error("RoleInfo error", err)
		base.Error(r, err.Error())
	}

	roleSlice := g.SliceStr{}
	for _, userRole := range userRoleList {
		roleSlice = append(roleSlice, gconv.String(userRole.RoleId))
	}
	roleIds := gstr.Join(roleSlice, ",")

	base.Succ(r, g.Map{
		"userId":  userId,
		"roleIds": roleIds,
		"list":    roleList,
	})
}

// 用户绑定角色
func (action *Action) RoleSave(r *ghttp.Request) {
	userId := r.GetInt("userid")
	roleIds := r.GetString("roleids")
	if userId == 0 {
		base.Fail(r, "参数错误")
	}

	// 保存角色信息
	user.SaveUserRole(userId, roleIds)

	base.Succ(r, "")
}

// 修改当前用户密码
func (action *Action) Password(r *ghttp.Request) {
	userId := base.GetUser(r).Id
	model, err := user.GetById(gconv.Int64(userId))
	if err != nil {
		glog.Error(" Menu error", err)
		base.Error(r, "登录信息获取异常")
	} else if model.Id <= 0 {
		base.Fail(r, "登录异常")
	}

	password := r.GetString("password")
	newPassword := r.GetString("newPassword")
	if password == "" || newPassword == "" {
		base.Fail(r, "参数错误")
	}

	reqPassword, err := gmd5.Encrypt(password + model.Salt)
	if err != nil {
		glog.Error(" Password error", err)
		base.Error(r, "密码错误.")
	}

	if reqPassword != model.Password {
		base.Fail(r, "原密码错误")
	}

	model.UpdateId = userId
	model.UpdateTime = library.GetNow()
	model.Password, err = gmd5.Encrypt(newPassword + model.Salt)
	if err != nil {
		glog.Error(" Password error", err)
		base.Error(r, "密码错误!")
	}

	_, err = model.Update()

	if err != nil {
		glog.Error(" Password error", err)
		base.Error(r, "密码错误.")
	}

	base.Succ(r, "")
}

// 获取当前用户信息
func (action *Action) Info(r *ghttp.Request) {
	userId := base.GetUser(r).Id
	model, err := user.GetById(gconv.Int64(userId))
	if err != nil {
		glog.Error(" Menu error", err)
		base.Error(r, "登录信息获取异常")
	} else if model.Id <= 0 {
		base.Fail(r, "登录异常")
	}

	model.Password = ""
	model.Salt = ""
	base.Succ(r, model)
}

// 获取当前用户信息
func (action *Action) Menu(r *ghttp.Request) {
	userId := base.GetUser(r).Id
	model, err := user.GetById(gconv.Int64(userId))
	if err != nil {
		glog.Error(" Menu error", err)
		base.Error(r, "登录信息获取异常")
	} else if model.Id <= 0 {
		base.Fail(r, "登录异常")
	}

	list, err := menu.ListUser(model.Id, model.UserType)
	if err != nil {
		glog.Error(" Menu error", err)
		base.Error(r, "数据获取异常")
	}

	base.Succ(r, list)
}
