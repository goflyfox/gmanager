package log

import (
	"github.com/gogf/gf/util/gconv"
)

// Fill with you ideas below.

const (
	LOGIN      = "登录"
	LOGOUT     = "登出"
	INSERT     = "插入"
	UPDATE     = "更新"
	DELETE     = "删除"
	TypeEdit   = 2
	TypeSystem = 1
)

func (r *Entity) PkVal() int {
	return gconv.Int(r.Id)
}

func (r *Entity) TableName() string {
	return Table
}

func (m *arModel) Columns() string {
	sqlColumns := "t.id,t.log_type as logType,t.oper_object as operObject,t.oper_table as operTable,t.oper_id as operId,t.oper_type as operType,t.oper_remark as operRemark,t.enable,t.update_time as updateTime,t.update_id as updateId,t.create_time as createTime,t.create_id as createId"
	return sqlColumns
}
