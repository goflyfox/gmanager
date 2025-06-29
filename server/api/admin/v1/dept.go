package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"gmanager/internal/admin/model/entity"
	input2 "gmanager/internal/admin/model/input"
)

type DeptListReq struct {
	g.Meta   `path:"/dept/list" method:"POST" perms:"admin:dept:query" tags:"部门管理" summary:"部门列表"`
	Keywords string `json:"keywords" dc:"名称"`
	Code     string `json:"code"  dc:"部门编码"`
	Enable   int    `json:"enable" dc:"是否启用"`
	input2.PageReq
}

type DeptListRes struct {
	List []*input2.DeptTreeRes `json:"list" dc:"部门列表"`
	input2.PageRes
}

type DeptOptionsReq struct {
	g.Meta `path:"/dept/options" method:"post" tags:"部门管理" summary:"部门下拉列表"`
	Enable int `json:"enable" dc:"是否启用"`
}

type DeptOptionsRes = []*input2.OptionVal

type DeptGetReq struct {
	g.Meta `path:"/dept/get/:id" method:"get" perms:"admin:dept:query" tags:"部门管理" summary:"部门获取"`
	Id     int64 `json:"id" dc:"ID"`
}

type DeptGetRes = entity.Dept

type DeptSaveReq struct {
	g.Meta    `path:"/dept/save/:id" method:"post" perms:"admin:dept:save" tags:"部门管理" summary:"部门保存"`
	Id        int64  `json:"id"`
	ParentId  int64  `json:"parentId"  v:"required#父级不能为空"`
	Name      string `json:"name"   dc:"部门名称" v:"required#部门名称不能为空"`
	Code      string `json:"code"   dc:"部门编码" v:"required#部门编码不能为空"`
	Sort      int    `json:"sort" dc:"排序序号" v:"required#部门序号不能为空"`
	Linkman   string `json:"linkman" dc:"联系人"`
	LinkmanNo string `json:"linkmanNo" dc:"联系人号码"`
	Remark    string `json:"remark" dc:"备注"`
	Enable    int    `json:"enable" dc:"是否启用" v:"required#是否启不能为空"`
}

type DeptSaveRes struct {
}

type DeptDeleteReq struct {
	g.Meta `path:"/dept/delete/:ids" method:"post" perms:"admin:dept:delete" tags:"部门管理" summary:"部门删除"`
	Ids    string `json:"ids" dc:"删除id列表"`
}

type DeptDeleteRes struct {
}
