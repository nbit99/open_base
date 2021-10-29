package model

// 用户设备推送日志
type OwJPushLog struct {
	Id        int64  `json:"id" bson:"_id" tb:"ow_jpush_log" mg:"true"`
	Uid       string `json:"uid" bson:"uid"`
	ProcessID string `json:"processID" bson:"processID"`
	Owners    string `json:"owners" bson:"owners"`
	Request   string `json:"request" bson:"request"`
	Response  string `json:"response" bson:"response"`
	Ctime     int64  `json:"ctime" bson:"ctime"`
	Utime     int64  `json:"utime" bson:"utime"`
	State     int64  `json:"state" bson:"state"`
}
