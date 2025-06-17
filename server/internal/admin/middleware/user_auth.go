package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"gmanager/internal/admin/consts"
	"gmanager/internal/admin/logic"
	"gmanager/internal/library/gftoken"
)

// UserPerm 权限校验
func UserPerm(r *ghttp.Request) {
	// 不认证URL
	if gftoken.HasExcludePath(r) {
		r.Middleware.Next()
		return
	}

	ctx := r.Context()
	user := gftoken.GetSessionUser(ctx)
	// 用户登陆校验部分需要处理
	if user == nil {
		r.Middleware.Next()
		return
	}

	// 管理员不校验
	if user.UserType == consts.UserTypeAdmin {
		r.Middleware.Next()
		return
	}

	// 按钮权限校验
	perms := r.GetServeHandler().GetMetaTag("perms")
	if perms != "" { // 需要权限校验
		permArray := gstr.Split(perms, ",")
		for _, perm := range permArray {
			if !logic.User.CheckPerm(ctx, perm) {
				r.Response.WriteJson(ghttp.DefaultHandlerResponse{
					Code:    gcode.CodeSecurityReason.Code(),
					Message: "按钮权限不足",
				})
				r.ExitAll()
			}
		}
	}

	r.Middleware.Next()
}
