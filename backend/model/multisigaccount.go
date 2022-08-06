package model

type MultiSigAccount struct {
	Id        int64      `json:"id" gorm:"column:id"`
	Version   uint8      `json:"version" gorm:"version"`
	Threshold uint8      `json:"threshold" gorm:"threshold"`
	Accounts  []*Account `json:"accounts" gorm:"many2many:multisig_account_accounts"`
	Address   string     `json:"address" gorm:"uniqueIndex"`
}
