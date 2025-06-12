// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`                                             // 主键
	Name      string      `json:"name"      orm:"name"       description:"名称/11111/"`                                      // 名称/11111/
	Code      string      `json:"code"      orm:"code"       description:"角色编码"`                                           // 角色编码
	DataScope int         `json:"dataScope" orm:"data_scope" description:"数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）"` // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`                                             // 排序
	Remark    string      `json:"remark"    orm:"remark"     description:"说明//textarea"`                                   // 说明//textarea
	Enable    int         `json:"enable"    orm:"enable"     description:"是否启用//radio/1,启用,2,禁用"`                          // 是否启用//radio/1,启用,2,禁用
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"更新时间"`                                           // 更新时间
	UpdateId  int64       `json:"updateId"  orm:"update_id"  description:"更新人"`                                            // 更新人
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`                                           // 创建时间
	CreateId  int64       `json:"createId"  orm:"create_id"  description:"创建者"`                                            // 创建者
}
