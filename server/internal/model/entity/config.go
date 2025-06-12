// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Config is the golang structure for table config.
type Config struct {
	Id           int64       `json:"id"           orm:"id"            description:"主键"`                             // 主键
	Name         string      `json:"name"         orm:"name"          description:"名称"`                             // 名称
	Key          string      `json:"key"          orm:"key"           description:"键"`                              // 键
	Value        string      `json:"value"        orm:"value"         description:"值"`                              // 值
	Code         string      `json:"code"         orm:"code"          description:"编码"`                             // 编码
	DataType     int         `json:"dataType"     orm:"data_type"     description:"数据类型//radio/1,KV配置,2,字典,3,字典数据"` // 数据类型//radio/1,KV配置,2,字典,3,字典数据
	ParentId     int64       `json:"parentId"     orm:"parent_id"     description:"类型"`                             // 类型
	ParentKey    string      `json:"parentKey"    orm:"parent_key"    description:""`                               //
	Remark       string      `json:"remark"       orm:"remark"        description:"备注"`                             // 备注
	Sort         int         `json:"sort"         orm:"sort"          description:"排序号"`                            // 排序号
	CopyStatus   int         `json:"copyStatus"   orm:"copy_status"   description:"拷贝状态 1 拷贝  2  不拷贝"`              // 拷贝状态 1 拷贝  2  不拷贝
	ChangeStatus int         `json:"changeStatus" orm:"change_status" description:"1 可以更改 2 不可更改"`                  // 1 可以更改 2 不可更改
	Enable       int         `json:"enable"       orm:"enable"        description:"是否启用//radio/1,启用,2,禁用"`          // 是否启用//radio/1,启用,2,禁用
	UpdateAt     *gtime.Time `json:"updateAt"     orm:"update_at"     description:"更新时间"`                           // 更新时间
	UpdateId     int64       `json:"updateId"     orm:"update_id"     description:"更新人"`                            // 更新人
	CreateAt     *gtime.Time `json:"createAt"     orm:"create_at"     description:"创建时间"`                           // 创建时间
	CreateId     int64       `json:"createId"     orm:"create_id"     description:"创建者"`                            // 创建者
}
