package model

// 多签用户个人地址
type OwMSAddress struct {
	Id      int64  `json:"id" bson:"_id" tb:"ow_msaddress" mg:"true"`
	AppID   string `json:"appID" bson:"appID"`
	UserID  string `json:"userID" bson:"userID"`
	Type    int64  `json:"type" bson:"type"`
	Orderno int64  `json:"orderno" bson:"orderno"`
	Symbol  string `json:"symbol" bson:"symbol"`
	Address string `json:"address" bson:"address"`
	Remark  string `json:"remark" bson:"remark"`
	Ctime   int64  `json:"ctime" bson:"ctime"`
	Utime   int64  `json:"utime" bson:"utime"`
	State   int64  `json:"state" bson:"state"`
}
