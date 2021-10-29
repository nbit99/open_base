package model

// 多签钱包绑定用户
type OwMSWalletBind struct {
	Id       int64  `json:"id" bson:"_id" tb:"ow_mswallet_bind" mg:"true"`
	AppID    string `json:"appID" bson:"appID"`
	UserID   string `json:"userID" bson:"userID"`
	WalletID string `json:"walletID" bson:"walletID"`
	Defop    int64  `json:"defop" bson:"defop"`
	Ctime    int64  `json:"ctime" bson:"ctime"`
	Utime    int64  `json:"utime" bson:"utime"`
	State    int64  `json:"state" bson:"state"`
}
