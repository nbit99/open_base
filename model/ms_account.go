package model

// 多签账户
type OwMSAccount struct {
	Id        int64  `json:"id" bson:"_id" tb:"ow_msaccount" mg:"true"`
	AppID     string `json:"appID" bson:"appID"`
	WalletID  string `json:"walletID" bson:"walletID"`
	AccountID string `json:"accountID" bson:"accountID"`
	PreResult string `json:"preResult" bson:"preResult"`
	Ctime     int64  `json:"ctime" bson:"ctime"`
	Utime     int64  `json:"utime" bson:"utime"`
	State     int64  `json:"state" bson:"state"`
}
