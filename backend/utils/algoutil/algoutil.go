package algoutil

import (
	"bytes"
	"crypto/ed25519"
	"multisigdb-svc/model"

	"github.com/algorand/go-algorand-sdk/encoding/msgpack"
	"github.com/algorand/go-algorand-sdk/types"
)

func AccountsToAlgoAddresses(accounts []*model.Account) []types.Address {
	addrs := make([]types.Address, 0)
	for _, acc := range accounts {
		addr, _ := types.DecodeAddress(acc.Address)
		addrs = append(addrs, addr)
	}
	return addrs
}

func VerifySignedTransaction(pubkey ed25519.PublicKey, transaction types.Transaction, sig []byte) bool {
	domainSeparator := []byte("TX")
	encodedTxn := msgpack.Encode(transaction)
	msgParts := [][]byte{domainSeparator, encodedTxn}
	toVerify := bytes.Join(msgParts, nil)
	ret := ed25519.Verify(pubkey, toVerify, sig)
	if ret {
		return true
	}
	return false
}
