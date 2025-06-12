package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"gmanager/internal/model/entity"
	"gmanager/internal/model/input"
)

type LogListReq struct {
	g.Meta   `path:"/log/list" method:"post" tags:"日志管理" summary:"日志列表"`
	Keywords string      `json:"keywords" dc:"日志/手机号/昵称名称"`
	LogType  int         `json:"logType"  dc:"类型"`
	Enable   int         `json:"enable" dc:"是否启用"`
	CreateAt g.MapIntStr `json:"createAt" dc:"创建时间区间"`
	input.PageReq
}

type LogListRes struct {
	List []*entity.Log `json:"list" dc:"日志列表"`
	input.PageRes
}

type LogGetReq struct {
	g.Meta `path:"/log/get/:id" method:"get" tags:"日志管理" summary:"日志获取"`
	Id     int64 `json:"id" dc:"ID"`
}

type LogGetRes = entity.Log

type LogDeleteReq struct {
	g.Meta `path:"/log/delete/:ids" method:"post" tags:"日志管理" summary:"日志删除"`
	Ids    string `json:"ids" dc:"删除id列表"`
}

type LogDeleteRes struct {
}

type LogVisitTrendReq struct {
	g.Meta    `path:"/log/visit-trend" method:"post" tags:"日志管理" summary:"日志列表"`
	StartDate string `json:"startDate" dc:"开始日期"`
	EndDate   string `json:"endDate" dc:"结束日期"`
}

type LogVisitTrendRes struct {
	Dates  []string `json:"dates" dc:"日期列表"`
	PVList []int    `json:"pvList" dc:"PV列表"`
	IPList []int    `json:"ipList" dc:"IP列表"`
}
