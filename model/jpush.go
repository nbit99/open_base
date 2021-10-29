package model

// 用户设备推送
type OwJPush struct {
	Id      int64    `json:"id" bson:"_id" tb:"ow_jpush" mg:"true"`
	Uid     string   `json:"uid" bson:"uid"`
	JPushId []string `json:"jpushId" bson:"jpushId"`
	Alias   []string `json:"alias" bson:"alias"`
	Tag     []string `json:"tag" bson:"tag"`
	Ctime   int64    `json:"ctime" bson:"ctime"`
	Utime   int64    `json:"utime" bson:"utime"`
	State   int64    `json:"state" bson:"state"`
}
