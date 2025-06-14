// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Dept is the golang structure for table dept.
type Dept struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`                    // 主键
	ParentId  int64       `json:"parentId"  orm:"parent_id"  description:"上级机构"`                  // 上级机构
	Name      string      `json:"name"      orm:"name"       description:"部门/11111"`              // 部门/11111
	Code      string      `json:"code"      orm:"code"       description:"机构编码"`                  // 机构编码
	Sort      int         `json:"sort"      orm:"sort"       description:"序号"`                    // 序号
	Linkman   string      `json:"linkman"   orm:"linkman"    description:"联系人"`                   // 联系人
	LinkmanNo string      `json:"linkmanNo" orm:"linkman_no" description:"联系人电话"`                 // 联系人电话
	Remark    string      `json:"remark"    orm:"remark"     description:"机构描述"`                  // 机构描述
	Enable    int         `json:"enable"    orm:"enable"     description:"是否启用//radio/1,启用,2,禁用"` // 是否启用//radio/1,启用,2,禁用
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"更新时间"`                  // 更新时间
	UpdateId  int64       `json:"updateId"  orm:"update_id"  description:"更新人"`                   // 更新人
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`                  // 创建时间
	CreateId  int64       `json:"createId"  orm:"create_id"  description:"创建者"`                   // 创建者
}
