package system

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/crypto/gmd5"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/gogf/gf/g/text/gstr"
	"github.com/gogf/gf/g/util/gconv"
	"github.com/gogf/gf/g/util/grand"
	"gmanager/utils"
	"gmanager/utils/base"
)

type UserAction struct {
	base.BaseRouter
}

var (
	actionNameUser = "UserAction"
)

// path: /index
func (action *UserAction) Index(r *ghttp.Request) {
	tplFile := "pages/system/user_index.html"
	err := r.Response.WriteTpl(tplFile, g.Map{
		"now": gtime.Datetime(),
	})

	if err != nil {
		glog.Error(err)
	}
}

// path: /get/{id}
func (action *UserAction) Get(r *ghttp.Request) {
	id := r.GetInt("id")
	model := SysUser{Id: id}.Get()
	if model.Id <= 0 {
		base.Fail(r, actionNameUser+" get fail")
	}

	base.Succ(r, model)
}

// path: /delete/{id}
func (action *UserAction) Delete(r *ghttp.Request) {
	id := r.GetInt("id")

	model := SysUser{Id: id}
	model.UpdateId = base.GetUser(r).Id
	model.UpdateTime = utils.GetNow()

	num := model.Delete()
	if num <= 0 {
		base.Fail(r, actionNameUser+" delete fail")
	}

	base.Succ(r, "")
}

// path: /save
func (action *UserAction) Save(r *ghttp.Request) {
	model := SysUser{}
	err := gconv.Struct(r.GetPostMap(), &model)
	if err != nil {
		glog.Error(actionNameUser+" save struct error", err)
		base.Error(r, "save error")
	}

	userId := base.GetUser(r).Id

	model.UpdateId = userId
	model.UpdateTime = utils.GetNow()

	var num int64
	if model.Id <= 0 {
		model.CreateId = userId
		model.CreateTime = utils.GetNow()
		model.Uuid = grand.Str(32)
		model.Salt = grand.Str(6)
		if model.Password == "" {
			tmpStr, err := gmd5.Encrypt("123456")
			if err != nil {
				glog.Error(actionNameUser+" Password Encrypt error", err)
				base.Error(r, "Password Encrypt")
			}

			model.Password, err = gmd5.Encrypt(tmpStr + model.Salt)
			if err != nil {
				glog.Error(actionNameUser+" Password Encrypt Salt error", err)
				base.Error(r, "Password Encrypt Salt error")
			}
		}

		num = model.Insert()
	} else {
		num = model.Update()
	}

	if num <= 0 {
		base.Fail(r, actionNameUser+" save fail")
	}

	base.Succ(r, "")
}

// path: /list
func (action *UserAction) List(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := SysUser{}

	list := model.List(&form)
	base.Succ(r, list)
}

// path: /page
func (action *UserAction) Page(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := SysUser{}

	page := model.Page(&form)
	base.Succ(r, g.Map{"list": page, "form": form})
}

// path: /jqgrid
func (action *UserAction) Jqgrid(r *ghttp.Request) {
	form := base.NewForm(r.GetPostMap())
	model := SysUser{}

	page := model.Page(&form)
	r.Response.WriteJson(g.Map{
		"page":    form.Page,
		"rows":    page,
		"total":   form.TotalPage,
		"records": form.TotalSize,
	})
}

// 获取用户角色信息
func (action *UserAction) RoleInfo(r *ghttp.Request) {
	userId := r.GetPostInt("userId")
	if userId == 0 {
		base.Fail(r, "参数错误")
	}
	form := base.NewForm(r.GetPostMap())
	form.SetParam("userId", gconv.String(userId))

	// 角色列表
	roleList := SysRole{}.List(&form)

	// 选择的列表
	userRoleList := SysUserRole{}.List(&form)
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
func (action *UserAction) RoleSave(r *ghttp.Request) {
	userId := r.GetPostInt("userid")
	roleIds := r.GetPostString("roleids")
	if userId == 0 {
		base.Fail(r, "参数错误")
	}

	SysUserRole{UserId: userId}.DeleteByUserId()
	if roleIds != "" {
		roleIdArray := gstr.Split(roleIds, ",")
		for _, roleId := range roleIdArray {
			SysUserRole{UserId: userId, RoleId: gconv.Int(roleId)}.Insert()
		}
	}

	base.Succ(r, "")
}

// 修改当前用户密码
func (action *UserAction) Password(r *ghttp.Request) {
	userId := base.GetUser(r).Id
	model := SysUser{Id: userId}.Get()
	if model.Id <= 0 {
		base.Fail(r, "登录异常")
	}

	password := r.GetPostString("password")
	newPassword := r.GetPostString("newPassword")
	if password == "" || newPassword == "" {
		base.Fail(r, "参数错误")
	}

	reqPassword, err := gmd5.Encrypt(password + model.Salt)
	if err != nil {
		glog.Error(actionNameUser+" Password error", err)
		base.Error(r, "Password error")
	}

	if reqPassword != model.Password {
		base.Fail(r, "原密码错误")
	}

	model.UpdateId = userId
	model.UpdateTime = utils.GetNow()
	model.Password, err = gmd5.Encrypt(newPassword + model.Salt)
	if err != nil {
		glog.Error(actionNameUser+" Password error", err)
		base.Error(r, "Password error")
	}

	num := model.Update()

	if num <= 0 {
		base.Fail(r, actionNameUser+" Password fail")
	}

	base.Succ(r, "")
}

// 获取当前用户信息
func (action *UserAction) Info(r *ghttp.Request) {
	userId := base.GetUser(r).Id
	model := SysUser{Id: userId}.Get()
	if model.Id <= 0 {
		base.Fail(r, "登录异常")
	}

	model.Password = ""
	model.Salt = ""
	base.Succ(r, model)
}

// 获取当前用户信息
func (action *UserAction) Menu(r *ghttp.Request) {
	userId := base.GetUser(r).Id
	model := SysUser{Id: userId}.Get()
	if model.Id <= 0 {
		base.Fail(r, "登录异常")
	}

	list := SysMenu{}.ListUser(model.Id, model.UserType)

	base.Succ(r, list)
}
