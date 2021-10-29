package model

const (
	ETH  = "ETH"
	TRUE = "TRUE"
)

// 货币信息
type OwSymbol struct {
	Id           int64  `json:"id" bson:"_id" tb:"ow_symbol" mg:"true"`
	Top          string `json:"top" bson:"top"`
	Parent       string `json:"parent" bson:"parent"`
	Ctype        int64  `json:"ctype" bson:"ctype"`
	Name         string `json:"name" bson:"name"`
	Coin         string `json:"coin" bson:"coin"`
	Curve        int64  `json:"curve" bson:"curve"`
	Resume       string `json:"resume" bson:"resume"`
	Icon         string `json:"icon" bson:"icon"`
	Orderno      int64  `json:"orderno" bson:"orderno"`
	Confirm      int64  `json:"confirm" bson:"confirm"`
	Decimals     int64  `json:"decimals" bson:"decimals"`
	BalanceMode  int64  `json:"balanceMode" bson:"balanceMode"`
	SupportMemo  int64  `json:"supportMemo" bson:"supportMemo"`
	OnlyContract int64  `json:"onlyContract" bson:"onlyContract"`
	SupportBatch int64  `json:"supportBatch" bson:"supportBatch"` // 是否支持批量交易 0.否 1.是
	AutoContract int64  `json:"autoContract" bson:"autoContract"` // 是否自动录入合约 0.否 1.是
	WithdrawStop int64  `json:"withdrawStop" bson:"withdrawStop"`
	BlockStop    int64  `json:"blockStop" bson:"blockStop"`
	Ctime        int64  `json:"ctime" bson:"ctime"`
	Utime        int64  `json:"utime" bson:"utime"`
	State        int64  `json:"state" bson:"state"`
}

type SymbolRpcReq struct {
	Symbol string
}

type SymbolRpcResp struct {
	Symbol string
}
