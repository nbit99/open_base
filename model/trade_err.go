package model

// 交易单创建广播异常
type OwTradeErr struct {
	Id        int64                  `json:"id" bson:"_id" tb:"ow_trade_err" mg:"true"`
	AppID     string                 `json:"appID" bson:"appID"`
	WalletID  string                 `json:"walletID" bson:"walletID"`
	AccountID string                 `json:"accountID" bson:"accountID"`
	Type      int64                  `json:"type" bson:"type"` // 0,创建 1.汇总 2.广播
	Sid       string                 `json:"sid" bson:"sid"`
	Request   map[string]interface{} `json:"request" bson:"request"`
	Rawtx     map[string]interface{} `json:"rawtx" bson:"rawtx"`
	ErrorMsg  string                 `json:"errorMsg" bson:"errorMsg"`
	Ctime     int64                  `json:"ctime" bson:"ctime"`
}
