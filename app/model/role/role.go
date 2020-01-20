package role

import "github.com/gogf/gf/util/gconv"

// Fill with you ideas below.

func (r *Entity) PkVal() int {
	return gconv.Int(r.Id)
}

func (r *Entity) TableName() string {
	return Table
}

func (m *arModel) Columns() string {
	sqlColumns := "t.id,t.name,t.status,t.sort,t.remark,t.enable,t.update_time as updateTime,t.update_id as updateId,t.create_time as createTime,t.create_id as createId"
	return sqlColumns
}
