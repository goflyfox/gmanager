package started

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/text/gstr"
	"github.com/gogf/gf/g/util/gconv"
)

var TableInfo g.MapStrStr

func Start() {
	var dbName = "gmanager"
	link := g.Config().GetString("database.link")
	if link != "" {
		dbName = gstr.Split(link, "/")[1]
	}
	r, err := g.DB().Table("INFORMATION_SCHEMA.TABLES").Fields(
		"table_name as name,table_comment as comment").Where(
		"table_schema = ?", dbName).Select()
	if err != nil {
		glog.Error("gstart tables error", err)
	} else {
		TableInfo = g.MapStrStr{}
		list := r.ToList()
		for _, value := range list {
			TableInfo[gconv.String(value["name"])] = gconv.String(value["comment"])
		}
		glog.Info("gstart table info finish", TableInfo)
	}
}
