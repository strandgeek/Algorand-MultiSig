package dbutil

import (
	"multisigdb-svc/model"

	"gorm.io/gorm"
)

func addrExists(addr string, accounts []*model.Account) bool {
	for _, acc := range accounts {
		if acc.Address == addr {
			return true
		}
	}
	return false
}

func GetOrCreateAccountByAddresses(db *gorm.DB, addresses []string) ([]*model.Account, error) {
	var accounts []*model.Account
	for _, addr := range addresses {
		var acc *model.Account
		db.Where(model.Account{Address: addr}).Attrs(model.Account{Address: addr}).FirstOrCreate(&acc)
		accounts = append(accounts, acc)
	}
	return accounts, nil
}

func GetSignerIndex(accounts []*model.Account, addr string) int {
	for idx, acc := range accounts {
		if addr == acc.Address {
			return idx
		}
	}
	return -1
}
