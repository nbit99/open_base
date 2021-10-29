package model

// 多签钱包
type OwMSWallet struct {
	Id          int64  `json:"id" bson:"_id" tb:"ow_mswallet" mg:"true"`
	AppID       string `json:"appID" bson:"appID"`
	UserID      string `json:"userID" bson:"userID"`
	WalletID    string `json:"walletID" bson:"walletID"`
	SigWalletID string `json:"sigWalletID" bson:"sigWalletID"`
	Signature   string `json:"signature" bson:"signature"`
	Ctime       int64  `json:"ctime" bson:"ctime"`
	Utime       int64  `json:"utime" bson:"utime"`
	State       int64  `json:"state" bson:"state"`
}
