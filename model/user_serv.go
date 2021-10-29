package model

// 用户服务配置
type OwUserServ struct {
	Id           int64    `json:"id" bson:"_id" tb:"ow_user_serv" mg:"true"`
	Uid          string   `json:"uid" bson:"uid"`                   // 用户ID
	OrderNo      string   `json:"orderNo" bson:"orderNo"`           // 订单号
	Level        int64    `json:"level" bson:"level"`               // 套餐级别 1.开发者模式 2.应用级 3.企业级
	ApplyAt      int64    `json:"applyAt" bson:"applyAt"`           // 服务开通时间
	ExpireAt     int64    `json:"expireAt" bson:"expireAt"`         // 服务到期时间
	OpenSymbol   []string `json:"openSymbol" bson:"openSymbol"`     // 开通的币种,默认[BTC,ETH,LTC,XRP,BCH]
	AllowLimit   int64    `json:"allowLmit" bson:"allowLmit"`       // API访问频率 默认60/分钟
	AllowSymbol  int64    `json:"allowSymbol" bson:"allowSymbol"`   // 允许开通币种数量 开发者5, 应用级10 , 企业级15
	AllowApp     int64    `json:"allowApp" bson:"allowApp"`         // 允许开通应用数量
	AllowWallet  int64    `json:"allowWallet" bson:"allowWallet"`   // 允许开通钱包数量 0.无限
	AllowAccount int64    `json:"allowAccount" bson:"allowAccount"` // 允许开通账号数量 0.无限
	AllowAddress int64    `json:"allowAddress" bson:"allowAddress"` // 允许开通地址数量
	AllowNotify  int64    `json:"allowNotify" bson:"allowNotify"`   // 允许每月通知交易单数量
	Ctime        int64    `json:"ctime" bson:"ctime"`
	Utime        int64    `json:"utime" bson:"utime"`
	State        int64    `json:"state" bson:"state"`
}

// 用户服务统计
type OwUserServReport struct {
	Id         int64  `json:"id" bson:"_id" tb:"ow_user_serv_report" mg:"true"`
	Uid        string `json:"uid" bson:"uid"`               // 用户ID
	AppID      string `json:"appID" bson:"appID"`           // 应用ID
	WalletSum  int64  `json:"walletSum" bson:"walletSum"`   // 创建钱包数量
	AccountSum int64  `json:"accountSum" bson:"accountSum"` // 创建账号数量
	AddressSum int64  `json:"addressSum" bson:"addressSum"` // 创建地址数量
	NotifySum  int64  `json:"notifySum" bson:"notifySum"`   // 通知交易单数量
	MonthDate  int64  `json:"monthDate" bson:"monthDate"`   // 通知统计月份时间
	Ctime      int64  `json:"ctime" bson:"ctime"`
	Utime      int64  `json:"utime" bson:"utime"`
	State      int64  `json:"state" bson:"state"`
}

// 用户发送通知交易统计
type OwUserServReportLog struct {
	Id         int64  `json:"id" bson:"_id" tb:"ow_user_serv_reportlog" mg:"true"`
	Uid        string `json:"uid" bson:"uid"`               // 用户ID
	AppID      string `json:"appID" bson:"appID"`           // 应用ID
	WalletSum  int64  `json:"walletSum" bson:"walletSum"`   // 创建钱包数量
	AccountSum int64  `json:"accountSum" bson:"accountSum"` // 创建账号数量
	AddressSum int64  `json:"addressSum" bson:"addressSum"` // 创建地址数量
	NotifySum  int64  `json:"notifySum" bson:"notifySum"`   // 通知交易单数量
	MonthDate  int64  `json:"monthDate" bson:"monthDate"`   // 通知统计月份时间
	Ctime      int64  `json:"ctime" bson:"ctime"`
	Utime      int64  `json:"utime" bson:"utime"`
	State      int64  `json:"state" bson:"state"`
}
