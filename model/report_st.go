package model

// 统计配置
type OwReportSt struct {
	Id             int64    `json:"id" bson:"_id" tb:"ow_report_st" mg:"true"`
	Symbol         string   `json:"symbol" bson:"symbol"`                 // 主币
	SymbolTag      string   `json:"symbolTag" bson:"symbolTag"`           // 币种-映射
	Token          string   `json:"token" bson:"token"`                   // 代币名称
	IsContract     int64    `json:"isContract" bson:"isContract"`         // 0.主币 1.代币
	ContractID     string   `json:"contractID" bson:"contractID"`         // 代币ID
	ExcludeAccount []string `json:"excludeAccount" bson:"excludeAccount"` // 排除在外的账号列表
	Price          int64    `json:"price" bson:"price"`                   // 单价换算/分
	Ctime          int64    `json:"ctime" bson:"ctime"`
	Utime          int64    `json:"utime" bson:"utime"`
	State          int64    `json:"state" bson:"state"`
}
