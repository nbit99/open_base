package model

// 用户安全表
type OwGAuth struct {
	Id         int64             `json:"id" bson:"_id" tb:"ow_gauth" mg:"true"`
	Uid        string            `json:"uid" bson:"uid"`
	AuthSeed   string            `json:"authSeed" bson:"authSeed"`
	AuthSecret string            `json:"authSecret" bson:"authSecret"`
	Device     map[string]string `json:"device" bson:"device"`
	Ctime      int64             `json:"ctime" bson:"ctime"`
	Utime      int64             `json:"utime" bson:"utime"`
	State      int64             `json:"state" bson:"state"`
}
