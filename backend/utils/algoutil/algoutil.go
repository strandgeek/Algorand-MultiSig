package algoutil

import (
	"multisigdb-svc/model"

	"github.com/algorand/go-algorand-sdk/types"
)

func AccountsToAlgoAddresses(accounts []*model.Account) []types.Address {
	addrs := make([]types.Address, 0)
	for _, acc := range accounts {
		var addr types.Address
		copy(addr[:], []byte(acc.Address))
		addrs = append(addrs, addr)
	}
	return addrs
}
