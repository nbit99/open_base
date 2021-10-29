package model

// 应用节点证书
type OwAppNode struct {
	Id     int64  `json:"id" bson:"_id" tb:"ow_app_node" mg:"true"`
	AppID  string `json:"appID" bson:"appID"`
	PriKey string `json:"priKey" bson:"priKey"`
	PubKey string `json:"pubKey" bson:"pubKey"`
	NodeId string `json:"nodeId" bson:"nodeId"`
	Ctime  int64  `json:"ctime" bson:"ctime"`
	Utime  int64  `json:"utime" bson:"utime"`
	State  int64  `json:"state" bson:"state"`
}
