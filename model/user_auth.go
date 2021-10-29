package model

// 用户谷歌认证
type OwUserAuth struct {
	Id     int64  `json:"id" bson:"_id" tb:"ow_user_auth" mg:"true"`
	Mobile string `json:"mobile" bson:"mobile"`
	Uid    string `json:"uid" bson:"uid"`
	Seed   string `json:"seed" bson:"seed"`
	Secret string `json:"secret" bson:"secret"`
	Ctime  int64  `json:"ctime" bson:"ctime"`
	Utime  int64  `json:"utime" bson:"utime"`
	State  int64  `json:"state" bson:"state"`
}
