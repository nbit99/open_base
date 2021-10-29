package model

// 合约回执
type OwReceipt struct {
	Id           int64  `json:"id" bson:"_id" tb:"ow_receipt" mg:"true"`
	AppID        string `json:"appID" bson:"appID"`
	WalletID     string `json:"walletID" bson:"walletID"`
	AccountID    string `json:"accountID" bson:"accountID"`
	Sid          string `json:"sid" bson:"sid"`
	Symbol       string `json:"symbol" bson:"symbol"`
	ContractID   string `json:"contractID" bson:"contractID"`
	ContractName string `json:"contractName" bson:"contractName"`
	ContractAddr string `json:"contractAddr" bson:"contractAddr"`
	TxID         string `json:"txID" bson:"txID"`
	WxID         string `json:"wxID" bson:"wxID"`
	From         string `json:"from" bson:"from"`
	To           string `json:"to" bson:"to"`
	Value        string `json:"value" bson:"value"`
	Fees         string `json:"fees" bson:"fees"`
	RawReceipt   string `json:"rawReceipt" bson:"rawReceipt"`
	BlockHash    string `json:"blockHash" bson:"blockHash"`
	BlockHeight  int64  `json:"blockHeight" bson:"blockHeight"`
	ConfirmTime  int64  `json:"confirmTime" bson:"confirmTime"`
	Status       string `json:"status" bson:"status"`
	Reason       string `json:"reason" bson:"reason"`
	ExtParam     string `json:"extParam" bson:"extParam"`
	EventValue   string `json:"eventValue" bson:"eventValue"`
	Ctime        int64  `json:"ctime" bson:"ctime"`
	Utime        int64  `json:"utime" bson:"utime"`
	State        int64  `json:"state" bson:"state"`
}
