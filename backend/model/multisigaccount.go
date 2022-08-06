package model

type MultiSigAccount struct {
	Id        int64      `json:"id" gorm:"column:id"`
	Version   int        `json:"version" gorm:"version"`
	Threshold int        `json:"threshold" gorm:"threshold"`
	Accounts  []*Account `json:"accounts" gorm:"many2many:multisig_account_accounts"`
}
