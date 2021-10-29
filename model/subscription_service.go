package model

// 服务
type OwSubscriptionService struct {
	Id          int64  `json:"id" bson:"_id" tb:"ow_subscription_service" mg:"true"`
	Type        int64  `json:"type" bson:"type"` // SubscriptionType 1：试用版，2：开发版，3：企业版
	Name        string `json:"name" bson:"name"`
	Description string `json:"desc" bson:"desc"`
	Limit       int64  `json:"limit" bson:"limit"` // 限额，如限1个应用、限2个钱包、限5个账户、限1000个地址、限1000个交易通知每月
	Ctime       int64  `json:"ctime" bson:"ctime"`
	Utime       int64  `json:"utime" bson:"utime"`
	State       int64  `json:"state" bson:"state"`
}
