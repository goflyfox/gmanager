package user_role

import "github.com/gogf/gf/util/gconv"

// Fill with you ideas below.

func (r *Entity) PkVal() int {
	return gconv.Int(r.Id)
}

func (r *Entity) TableName() string {
	return Table
}

func (m *arModel) Columns() string {
	sqlColumns := "t.id,t.user_id as userId,t.role_id as roleId"
	return sqlColumns
}
