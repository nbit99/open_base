package model

// 服务定价，配置收款钱包
type OwSubscriptionPrice struct {
	Id          int64  `json:"id" bson:"_id" tb:"ow_subscription_price" mg:"true"`
	Level       int64  `json:"level" bson:"level"` // 套餐级别 SubscriptionLevel 1：试用版，2：开发版，3：企业版
	Description string `json:"desc" bson:"desc"`

	AppID     string `json:"appID" bson:"appID"` // 收费钱包
	WalletID  string `json:"walletID" bson:"walletID"`
	AccountID string `json:"accountID" bson:"accountID"`

	Price        string `json:"price" bson:"price"`               // 显示定价，每月 "10USDT"
	Amount       string `json:"amount" bson:"amount"`             // 定价，每月
	AnnualPrice  string `json:"annualPrice" bson:"annualPrice"`   // 显示定价，每年 "100USDT"
	AnnualAmount string `json:"annualAmount" bson:"annualAmount"` // 定价，每年
	Symbol       string `json:"symbol" bson:"symbol"`
	ContractID   string `json:"contractID" bson:"contractID"`

	Confirm int64 `json:"confirm" bson:"confirm"`
	Ctime   int64 `json:"ctime" bson:"ctime"`
	Utime   int64 `json:"utime" bson:"utime"`
	State   int64 `json:"state" bson:"state"`
}
