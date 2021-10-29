package model

import (
	"github.com/blocktree/openwallet/v2/common"
	"github.com/blocktree/openwallet/v2/openwallet"
	"github.com/blocktree/openwallet/v2/crypto"
	owkeychain "github.com/blocktree/go-owcdrivers/owkeychain"
)

// 资产账户
type OwAccount struct {
	Id               int64    `json:"id" bson:"_id" tb:"ow_account" mg:"true"`
	AppID            string   `json:"appID" bson:"appID"`
	WalletID         string   `json:"walletID" bson:"walletID"`
	AccountID        string   `json:"accountID" bson:"accountID"`
	Alias            string   `json:"alias" bson:"alias"`
	Symbol           string   `json:"symbol" bson:"symbol"`
	OtherOwnerKeys   []string `json:"otherOwnerKeys" bson:"otherOwnerKeys"`
	ReqSigs          int64    `json:"reqSigs" bson:"reqSigs"`
	IsTrust          int64    `json:"isTrust" bson:"isTrust"`
	Password         string   `json:"password" bson:"password"`
	PublicKey        string   `json:"publicKey" bson:"publicKey"`
	HdPath           string   `json:"hdPath" bson:"hdPath"`
	ContractAddress  string   `json:"contractAddress" bson:"contractAddress"`
	AccountIndex     int64    `json:"accountIndex" bson:"accountIndex"`
	Balance          string   `json:"balance" bson:"balance"`
	ConfirmBalance   string   `json:"confirmBalance" bson:"confirmBalance"`
	UnconfirmBalance string   `json:"unconfirmBalance" bson:"unconfirmBalance"`
	FastChange       int64    `json:"fastChange" bson:"fastChange"`
	ExtInfo          string   `json:"extInfo" bson:"extInfo"`
	AddressIndex     int64    `json:"addressIndex" bson:"addressIndex"`
	IsImport         int64    `json:"isImport" bson:"isImport"`
	Applytime        int64    `json:"applytime" bson:"applytime"`
	Succtime         int64    `json:"succtime" bson:"succtime"`
	Dealstate        int64    `json:"dealstate" bson:"dealstate"`
	Ctime            int64    `json:"ctime" bson:"ctime"`
	Utime            int64    `json:"utime" bson:"utime"`
	State            int64    `json:"state" bson:"state"`
}

// 中心数据库account转成openwallet.account
func (self OwAccount) ToAssetsAccount() *openwallet.AssetsAccount {
	return &openwallet.AssetsAccount{
		WalletID:        self.WalletID,
		Alias:           self.Alias,
		PublicKey:       self.PublicKey,
		AccountID:       self.AccountID,
		Index:           uint64(self.AccountIndex),
		AddressIndex:    int(self.AddressIndex),
		HDPath:          self.HdPath,
		OwnerKeys:       self.OtherOwnerKeys,
		Symbol:          self.Symbol,
		ContractAddress: self.ContractAddress,
		Required:        uint64(self.ReqSigs),
		IsTrust:         common.UIntToBool(uint64(self.IsTrust)),
	}
}

//computeKeyID 计算HDKey的KeyID
func (a *OwAccount) GetAccountID() string {

	if len(a.AccountID) > 0 {
		return a.AccountID
	}

	pub, err := owkeychain.OWDecode(a.PublicKey)
	if err != nil {
		return ""
	}

	//seed Keccak256 两次得到keyID
	hash := crypto.Keccak256(pub.GetPublicKeyBytes())
	hash = crypto.Keccak256(hash)

	a.AccountID = owkeychain.Encode(hash, owkeychain.BitcoinAlphabet)

	return a.AccountID
}
