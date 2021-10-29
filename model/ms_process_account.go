package model

// 多签账户流程
type OwProcessAccount struct {
	Id             int64  `json:"id" bson:"_id" tb:"ow_process_account" mg:"true"`
	AppID          string `json:"appID" bson:"appID"`
	UserID         string `json:"userID" bson:"userID"`
	ProcessID      string `json:"processID" bson:"processID"`
	WalletID       string `json:"walletID" bson:"walletID"`
	AccountID      string `json:"accountID" bson:"accountID"`
	StepIndex      int64  `json:"stepIndex" bson:"stepIndex"`
	Symbol         string `json:"symbol" bson:"symbol"`
	Operate        int64  `json:"operate" bson:"operate"`
	RoundPoint     string `json:"roundPoint" bson:"roundPoint"`
	PublicKeyPiece string `json:"publicKeyPiece" bson:"publicKeyPiece"`
	Reason         string `json:"reason" bson:"reason"`
	Request        string `json:"request" bson:"request"`
	Ctime          int64  `json:"ctime" bson:"ctime"`
	Utime          int64  `json:"utime" bson:"utime"`
	State          int64  `json:"state" bson:"state"`
}
