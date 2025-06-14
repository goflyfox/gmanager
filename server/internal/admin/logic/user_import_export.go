package logic

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "gmanager/api/admin/v1"
	"gmanager/internal/admin/consts"
	"gmanager/internal/admin/model/input"
	"gmanager/utility/excel"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

// Export 导出用户列表到Excel
func (s *user) Export(ctx context.Context, in *v1.UserExportReq) error {
	// 获取用户列表数据
	listReq := &v1.UserListReq{
		PageReq: input.PageReq{
			PageNum:  1,
			PageSize: 1000, // 导出全部数据
		},
		Keywords: in.Keywords,
		DeptId:   in.DeptId,
		Status:   in.Status,
		Enable:   in.Enable,
		CreateAt: in.CreateAt,
	}
	listRes, err := s.List(ctx, listReq)
	if err != nil {
		return gerror.Wrap(err, "获取用户数据失败")
	}

	//  导出格式
	type exportData struct {
		Id       int64  `json:"id"           description:"ID"`
		DeptName string `json:"deptId"  description:"部门" `
		UserName string `json:"userName"  description:"用户名称"`
		NickName string `json:"nickName" description:"昵称"`
		Mobile   string `json:"mobile"   description:"手机号"`
		Email    string `json:"email"   description:"邮箱"`
		Enable   string `json:"enable" description:"是否启用"`
		CreateAt string `json:"createAt"  description:"创建时间"`
	}

	var (
		titleList  = []string{"ID", "用户名", "昵称", "部门", "手机号", "邮箱", "状态", "创建时间"} // 设置表头
		fileName   = "用户信息导出-" + gctx.CtxId(ctx)
		sheetName  = "用户列表" // 创建工作表
		exportList []exportData
		row        exportData
	)

	deptNameMap, _ := Dept.DeptMap(ctx)
	// 格式化格式
	for _, v := range listRes.List {
		row.Id = v.Id
		row.DeptName = deptNameMap[v.DeptId].Name
		row.UserName = v.UserName
		row.NickName = v.NickName
		row.Mobile = v.Mobile
		row.Email = v.Email
		if v.Enable == consts.EnableYes {
			row.Enable = "正常"
		} else {
			row.Enable = "禁用"
		}
		row.CreateAt = v.CreateAt.String()
		exportList = append(exportList, row)
	}

	return excel.ExportByStructs(ctx, titleList, exportList, fileName, sheetName)
}

// Import 从Excel导入用户列表
// https://goframe.org/docs/web/http-client-file-uploading
func (s *user) Import(ctx context.Context, in *v1.UserImportReq) (res *v1.UserImportRes, err error) {
	res = &v1.UserImportRes{Code: consts.CodeOK}
	file, err := in.File.Open()
	if err != nil {
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			g.Log().Error(ctx, err)
		}
	}()

	excelData, err := excel.ParseExcel(ctx, file, "")
	if err != nil {
		return
	}
	if excelData == nil || len(excelData.Rows) == 0 {
		return
	}

	// 获取部门映射
	deptMap, err := Dept.DeptNameMap(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "获取部门数据失败")
	}

	// 遍历数据行
	for i, row := range excelData.Rows {
		errorMsg := "第" + gconv.String(i) + "行数据校验失败："
		if i == 0 { // 跳过表头
			continue
		}

		if len(row) < 7 { // 至少需要用户名、昵称、部门、手机号
			res.MessageList = append(res.MessageList, errorMsg+"导出数据格式不正确")
			res.InvalidCount++
			continue
		}

		// 解析行数据
		userName := row[0]
		nickName := row[1]
		gender := row[2]
		genderValue := ""
		mobile := row[3]
		email := row[4]
		roleName := row[5]
		deptName := row[6]
		// 验证必填字段
		if userName == "" || nickName == "" || deptName == "" || roleName == "" {
			res.MessageList = append(res.MessageList, errorMsg+"用户名、昵称、部门、角色不能为空")
			res.InvalidCount++
			continue
		}
		// 获取部门ID
		deptInfo, ok := deptMap[deptName]
		if !ok {
			res.MessageList = append(res.MessageList, errorMsg+"部门数据不存在")
			res.InvalidCount++
			continue
		}
		// 性别处理
		if gender != "" {
			configRes, err := Config.Value(ctx, &v1.ConfigValueReq{
				Name: gender,
			})
			if err != nil {
				g.Log().Error(ctx, err)
				res.MessageList = append(res.MessageList, errorMsg+"性别数据异常,"+err.Error())
				res.InvalidCount++
				continue
			}
			if configRes.Value == "" {
				res.MessageList = append(res.MessageList, errorMsg+"性别数据不存在")
				res.InvalidCount++
				continue
			}
			genderValue = configRes.Value
		}
		// 角色处理
		roleRes, err := Role.List(ctx, &v1.RoleListReq{Name: roleName})
		if err != nil {
			g.Log().Error(ctx, err)
			res.MessageList = append(res.MessageList, errorMsg+"角色数据异常,"+err.Error())
			res.InvalidCount++
			continue
		}
		if len(roleRes.List) == 0 {
			res.MessageList = append(res.MessageList, errorMsg+"角色数据不存在")
			res.InvalidCount++
			continue
		}
		var roleIds []int64
		for _, e := range roleRes.List {
			roleIds = append(roleIds, e.Id)
		}

		model := &v1.UserSaveReq{
			Id:       0,
			DeptId:   deptInfo.Id,
			UserName: userName,
			NickName: nickName,
			Gender:   gconv.Int(genderValue),
			Mobile:   mobile,
			Email:    email,
			Enable:   consts.EnableYes,
			RoleIds:  roleIds,
		}
		err = s.Save(ctx, model)
		if err != nil {
			g.Log().Error(ctx, err)
			res.MessageList = append(res.MessageList, errorMsg+"用户数据保存异常,"+err.Error())
			res.InvalidCount++
			continue
		}
		res.ValidCount++
	}
	if res.InvalidCount > 0 {
		res.Code = gcode.CodeBusinessValidationFailed.Code()
	}

	return res, nil
}

func (s *user) Template(ctx context.Context, req *v1.UserTemplateReq) error {
	// 模板文件路径
	fileName := "用户导入模板.xlsx"
	filePath := gfile.MainPkgPath() + "/resource/public/excel/" + fileName

	// 检查文件是否存在，如果存在则直接返回
	if gfile.Exists(filePath) {
		r := ghttp.RequestFromCtx(ctx)
		if r == nil {
			return gerror.New("ctx not http request")
		}
		r.Response.ServeFileDownload(filePath, fileName)
		return nil
	}

	return gerror.New("用户导入模板文件不存在")
}
