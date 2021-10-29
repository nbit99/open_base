package model

// 多签授权设备
type OwMSDevice struct {
	Id         int64  `json:"id" bson:"_id" tb:"ow_msdevice" mg:"true"`
	UserID     string `json:"userID" bson:"userID"`
	DeviceID   string `json:"deviceID" bson:"deviceID"`
	DeviceName string `json:"deviceName" bson:"deviceName"`
	OS         string `json:"os" bson:"os"`
	AccessTime int64  `json:"accessTime" bson:"accessTime"`
	Ctime      int64  `json:"ctime" bson:"ctime"`
	Utime      int64  `json:"utime" bson:"utime"`
	State      int64  `json:"state" bson:"state"`
}
