package model

// 短信接口
type OwSms struct {
	Id        int64  `json:"id" bson:"_id" tb:"ow_sms" mg:"true"`
	Area      string `json:"area" bson:"area"`
	Mobile    string `json:"mobile" bson:"mobile"`
	Access    string `json:"access" bson:"access"`
	Code      string `json:"code" bson:"code"`
	Content   string `json:"content" bson:"content"`
	Usestate  int64  `json:"usestate" bson:"usestate"`
	Usecount  int64  `json:"usecount" bson:"usecount"`
	Type      int64  `json:"type" bson:"type"`
	Applytime int64  `json:"applytime" bson:"applytime"`
	Exptime   int64  `json:"exptime" bson:"exptime"`
	Succtime  int64  `json:"succtime" bson:"succtime"`
	Ctime     int64  `json:"ctime" bson:"ctime"`
	Utime     int64  `json:"utime" bson:"utime"`
	State     int64  `json:"state" bson:"state"`
}
