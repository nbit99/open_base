package model

// 声明流程
type OwProcess struct {
	Id          int64    `json:"id" bson:"_id" tb:"ow_process" mg:"true"`
	Title       string   `json:"title" bson:"title"`
	AppID       string   `json:"appID" bson:"appID"`
	CurrID      string   `json:"currID" bson:"currID"`
	UserID      []string `json:"userID" bson:"userID"`
	WalletIDs   []string `json:"walletIDs" bson:"walletIDs"`
	ProcessID   string   `json:"processID" bson:"processID"`
	BatchNo     int64    `json:"batchNo" bson:"batchNo"`
	WalletID    string   `json:"walletID" bson:"walletID"`
	AccountID   string   `json:"accountID" bson:"accountID"`
	Symbol      string   `json:"symbol" bson:"symbol"`
	StepNumber  int64    `json:"stepNumber" bson:"stepNumber"`
	StepIndex   int64    `json:"stepIndex" bson:"stepIndex"`
	ProType     int64    `json:"proType" bson:"proType"`
	Applytime   int64    `json:"applytime" bson:"applytime"`
	Succtime    int64    `json:"succtime" bson:"succtime"`
	Expiretime  int64    `json:"expiretime" bson:"expiretime"`
	Reason      string   `json:"reason" bson:"reason"`
	Applystate  int64    `json:"applystate" bson:"applystate"`
	Dealstate   int64    `json:"dealstate" bson:"dealstate"`
	Process     string   `json:"process" bson:"process"`
	Transaction string   `json:"transaction" bson:"transaction"`
	Showme      int64    `json:"showme" bson:"showme"`
	Ctime       int64    `json:"ctime" bson:"ctime"`
	Utime       int64    `json:"utime" bson:"utime"`
	State       int64    `json:"state" bson:"state"`
}
