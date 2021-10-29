package model

// 订阅合约回执
type OwReceiptFollow struct {
	Id              int64    `json:"id" bson:"_id" tb:"ow_receipt_follow" mg:"true"`
	AppID           string   `json:"appID" bson:"appID"`
	FollowContracts []string `json:"followContracts" bson:"followContracts"` // contractID ...1
	Ctime           int64    `json:"ctime" bson:"ctime"`
	Utime           int64    `json:"utime" bson:"utime"`
	State           int64    `json:"state" bson:"state"`
}
