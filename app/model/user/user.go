package user

import "github.com/gogf/gf/util/gconv"

// Fill with you ideas below.

func (r *Entity) PkVal() int {
	return gconv.Int(r.Id)
}

func (r *Entity) TableName() string {
	return Table
}

func (m *arModel) Columns() string {
	sqlColumns := "t.id,t.uuid,t.username,t.password,t.salt,t.real_name as realName,t.depart_id as departId,t.user_type as userType,t.status,t.thirdid,t.endtime,t.email,t.tel,t.address,t.title_url as titleUrl,t.remark,t.theme,t.back_site_id as backSiteId,t.create_site_id as createSiteId,t.project_id as projectId,t.project_name as projectName,t.enable,t.update_time as updateTime,t.update_id as updateId,t.create_time as createTime,t.create_id as createId"
	return sqlColumns
}
