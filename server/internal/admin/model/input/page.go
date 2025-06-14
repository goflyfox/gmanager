package input

// PageReq 公共请求参数
type PageReq struct {
	PageNum  int    `json:"pageNum" example:"1" d:"1" v:"min:1#页码最小值不能低于1" dc:"当前页码"`                         //当前页码
	PageSize int    `json:"pageSize" example:"10" d:"10" v:"min:1|max:500#每页数量最小值不能低于1|最大值不能大于500" dc:"每页数量"` //每页数
	OrderBy  string //排序方式
}

// PageRes 列表公共返回
type PageRes struct {
	CurrentPage int `json:"currentPage" example:"0" dc:"当前页码"`
	Total       int `json:"total" example:"0" dc:"数据总行数"`
}
