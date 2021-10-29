package model

// 统计内容
type OwReportCt struct {
	Id         int64   `json:"id" bson:"_id" tb:"ow_report_ct" mg:"true"`
	Symbol     string  `json:"symbol" bson:"symbol"`         // 主币
	Token      string  `json:"token" bson:"token"`           // 代币名称
	IsContract int64   `json:"isContract" bson:"isContract"` // 0.主币 1.代币
	ContractID string  `json:"contractID" bson:"contractID"` // 代币ID
	Type       int64   `json:"type" bson:"type"`             // 1.充值 2.提币
	Price      float64 `json:"price" bson:"price"`           // 单价换算后最终价格
	Amount     float64 `json:"amount" bson:"amount"`         // 最终币种数量
	Depth      *Depth  `json:"depth" bson:"depth"`           // 买卖深度数据
	Date1      int64   `json:"date1" bson:"date1"`           // 日期 yyyyMMdd
	Date2      string  `json:"date2" bson:"date2"`           // 日期 HH:mm
	Ctime      int64   `json:"ctime" bson:"ctime"`
	Utime      int64   `json:"utime" bson:"utime"`
	State      int64   `json:"state" bson:"state"`
}

type Depth struct {
	BuyPrice     float64 `json:"buyPrice" bson:"buyPrice"`         // 可买1价格
	BuyDeep      float64 `json:"buyDeep" bson:"buyDeep"`           // 可买1数量
	BuyAvgPrice  float64 `json:"buyAvgPrice" bson:"buyAvgPrice"`   // 可买50平均价格
	BuyAvgDeep   float64 `json:"buyAvgDeep" bson:"buyAvgDeep"`     // 可买50平均数量
	SellPrice    float64 `json:"sellPrice" bson:"sellPrice"`       // 可卖1价格
	SellDeep     float64 `json:"sellDeep" bson:"sellDeep"`         // 可卖1数量
	SellAvgPrice float64 `json:"sellAvgPrice" bson:"sellAvgPrice"` // 可卖50平均价格
	SellAvgDeep  float64 `json:"sellAvgDeep" bson:"sellAvgDeep"`   // 可卖50平均数量
}
