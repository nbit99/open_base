package model

// 用户服务订阅
type OwUserSubscription struct {
	Id         int64    `json:"id" bson:"_id" tb:"ow_user_subscription" mg:"true"`
	Uid        string   `json:"uid" bson:"uid"`
	Type       int64    `json:"type" bson:"type"`       // SubscriptionType 1：试用版，2：开发版，3：企业版
	Symbols    []string `json:"symbols" bson:"symbols"` // 可空，限制symbol
	Expiration int64    `json:"expiration" bson:"expiration"`
	Ctime      int64    `json:"ctime" bson:"ctime"`
	Utime      int64    `json:"utime" bson:"utime"`
	State      int64    `json:"state" bson:"state"` // 1有效，2过期
}
