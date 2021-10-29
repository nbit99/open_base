package model

// 系统控制
type OwSetting struct {
	Id       int64  `json:"id" bson:"_id" tb:"ow_setting" mg:"true"`
	Platform string `json:"platform" bson:"platform"`
	IsOff    int64  `json:"isOff" bson:"isOff"`
	Ctime    int64  `json:"ctime" bson:"ctime"`
	Utime    int64  `json:"utime" bson:"utime"`
	State    int64  `json:"state" bson:"state"`
}
