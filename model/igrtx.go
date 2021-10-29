package model

// 忽略的交易单
type OwIgrtx struct {
	Id     int64  `json:"id" bson:"_id" tb:"ow_igrtx" mg:"true"`
	AppID  string `json:"appID" bson:"appID"`   // appid
	Symbol string `json:"symbol" bson:"symbol"` // 币种名称
	Txid   string `json:"txid" bson:"txid"`     // txid
	Ctime  int64  `json:"ctime" bson:"ctime"`
	Utime  int64  `json:"utime" bson:"utime"`
	State  int64  `json:"state" bson:"state"`
}
