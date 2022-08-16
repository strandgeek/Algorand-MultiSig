package model

type MultiSigAccount struct {
	Model
	Version      uint8         `json:"version" gorm:"version"`
	Threshold    uint8         `json:"threshold" gorm:"threshold"`
	Accounts     []*Account    `json:"accounts" gorm:"many2many:multisig_account_accounts"`
	Address      string        `json:"address" gorm:"uniqueIndex"`
	Transactions []Transaction `json:"transactions"`
}

func (m *MultiSigAccount) HasSigner(address string) bool {
	for _, acc := range m.Accounts {
		if acc.Address == address {
			return true
		}
	}
	return false
}
