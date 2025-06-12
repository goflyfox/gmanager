package input

type LogData struct {
	Model      any    `json:"model"`
	OperType   string `json:"oper_type"`
	OperObject string `json:"oper_object,omitempty"`
	OperRemark string `json:"operRemark,omitempty"`
	Operator   string `json:"operator,omitempty"`
	PkVal      int64  `json:"pk_val,omitempty"`
	TableName  string `json:"table_name,omitempty"`
}
