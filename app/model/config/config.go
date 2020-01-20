package config

import "github.com/gogf/gf/util/gconv"

// Fill with you ideas below.

func (r *Entity) PkVal() int {
	return gconv.Int(r.Id)
}

func (r *Entity) TableName() string {
	return Table
}

func (m *arModel) Columns() string {
	sqlColumns := "t.id,t.name,t.key,t.value,t.code,t.data_type as dataType,t.parent_id as parentId,t.parent_key as parentKey,t.sort,t.project_id as projectId,t.copy_status as copyStatus,t.change_status as changeStatus,t.enable,t.update_time as updateTime,t.update_id as updateId,t.create_time as createTime,t.create_id as createId"
	return sqlColumns
}
