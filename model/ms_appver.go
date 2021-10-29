package model

// 多签应用版本
type OwMSAppVer struct {
	Id          int64  `json:"id" bson:"_id" tb:"ow_app_version" mg:"true"`
	Platform    string `json:"platform" bson:"platform"`
	IsRefresh   int64  `json:"isRefresh" bson:"isRefresh"`
	VersionName string `json:"versionName" bson:"versionName"`
	VersionCode int64  `json:"versionCode" bson:"versionCode"`
	Appmark     string `json:"appmark" bson:"appmark"`
	Android     string `json:"android" bson:"android"`
	Iphone      string `json:"iphone" bson:"iphone"`
	Ctime       int64  `json:"ctime" bson:"ctime"`
	Utime       int64  `json:"utime" bson:"utime"`
	State       int64  `json:"state" bson:"state"`
}
