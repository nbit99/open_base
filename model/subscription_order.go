package model

// 服务订单
type OwSubscriptionOrder struct {
	Id         int64  `json:"id" bson:"_id" tb:"ow_subscription_order" mg:"true"`
	OrderNo    string `json:"orderNo" bson:"orderNo"`
	Uid        string `json:"uid" bson:"uid"`
	TotalPrice string `json:"totalPrice" bson:"totalPrice"`
	Level      int64  `json:"level" bson:"level"`       // 套餐级别 SubscriptionLevel 1：试用版，2：开发版，3：企业版
	Duration   int64  `json:"duration" bson:"duration"` // 月数
	Expiration int64  `json:"expiration" bson:"expiration"`
	Ctime      int64  `json:"ctime" bson:"ctime"`
	Utime      int64  `json:"utime" bson:"utime"`
	State      int64  `json:"state" bson:"state"` // -1: 已过期， 0:已创建，1:已转账，用户标志已支付，2:转账已广播，3；转账完成，区块已确认

	Symbol            string `json:"symbol" bson:"symbol"`
	ContractID        string `json:"contractID" bson:"contractID"`
	Confirm           int64  `json:"confirm" bson:"confirm"`
	ReceivableAddress string `json:"receivableAddress" bson:"receivableAddress"` // 应收地址
	ReceivableAmount  string `json:"receivableAmount" bson:"receivableAmount"`   // 应收金额

	Sid            string `json:"sid" bson:"sid"`
	TxID           string `json:"txID" bson:"txID"`
	WxID           string `json:"wxID" bson:"wxID"`
	FromAddress    string `json:"fromAddress" bson:"fromAddress"`
	ToAddress      string `json:"toAddress" bson:"toAddress"`
	ReceivedAmount string `json:"receivedAmount" bson:"receivedAmount"`
	BlockHash      string `json:"blockHash" bson:"blockHash"`
	BlockHeight    int64  `json:"blockHeight" bson:"blockHeight"`
}
