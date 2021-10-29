package model

// 客户端订阅信息
type OwSubscribe struct {
	Id              int64        `json:"id" bson:"_id" tb:"ow_subscribe" mg:"true"`
	AppID           string       `json:"appID" bson:"appID"`
	CallbackMode    int64        `json:"callbackMode" bson:"callbackMode"` // 回调模式，1：当前连接模式，长连接可以用当前连接，接收推送，2：新建连接模式，短连接可以采用建立回调服务接口接收推送
	CallbackNode    CallbackNode `json:"callbackNode" bson:"callbackNode"`
	SubscribeMethod []string     `json:"subscribeMethod" bson:"subscribeMethod"` // 订阅方法列表
	SubscribeToken  string       `json:"subscribeToken" bson:"subscribeToken"`   // 客户端令牌
	ClientIP        string       `json:"clientIP" bson:"clientIP"`
	ConnectState    int64        `json:"connectState" bson:"connectState"`   // 0.离线 1.在线
	ConnectDate     int64        `json:"connectDate" bson:"connectDate"`     // 最后连接时间
	UnConnectDate   int64        `json:"unConnectDate" bson:"unConnectDate"` // 断开连接时间
	Usestate        int64        `json:"usestate" bson:"usestate"`
	Applytime       int64        `json:"applytime" bson:"applytime"` // 创建时间
}

type CallbackNode struct {
	NodeID             string `json:"nodeID" bson:"nodeID"`
	Address            string `json:"address" bson:"address"`
	ConnectType        string `json:"connectType" bson:"connectType"`
	EnableKeyAgreement bool   `json:"enableKeyAgreement" bson:"enableKeyAgreement"`
	EnableSignature    bool   `json:"enableSignature" bson:"enableSignature"`
	EnableSSL          bool   `json:"enableSSL" bson:"enableSSL"`
}
