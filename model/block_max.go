package model

// 最大高度区块头
type OwBlockMax struct {
	Id     int64  `json:"id" bson:"_id" tb:"ow_block_max" mg:"true"`
	Height int64  `json:"height" bson:"height"`
	MaxID  int64  `json:"maxID" bson:"maxID"`
	Symbol string `json:"symbol" bson:"symbol"`
	Utime  int64  `json:"utime" bson:"utime"`
}
