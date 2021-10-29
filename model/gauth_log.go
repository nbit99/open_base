package model

// 用户安全日志表
type OwGAuthLog struct {
	Id        int64  `json:"id" bson:"_id" tb:"ow_gauth_log" mg:"true"`
	Type      int64  `json:"type" bson:"type"`
	Uid       string `json:"uid" bson:"uid"`
	LastIP    string `json:"lastIP" bson:"lastIP"`
	LastArea  string `json:"lastArea" bson:"lastArea"`
	LastTime  int64  `json:"lastTime" bson:"lastTime"`
	JPushId   string `json:"jpushId" bson:"jpushId"`
	DeviceId  string `json:"deviceId" bson:"deviceId"`
	DeviceMsg string `json:"deviceMsg" bson:"deviceMsg"`
	OsMsg     string `json:"osMsg" bson:"osMsg"`
	Ctime     int64  `json:"ctime" bson:"ctime"`
	Utime     int64  `json:"utime" bson:"utime"`
	State     int64  `json:"state" bson:"state"`
}
