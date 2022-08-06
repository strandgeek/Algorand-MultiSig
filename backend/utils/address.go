package utils

import (
	"github.com/algorand/go-algorand-sdk/types"
)

func IsValidAddress(addr string) bool {
	_, err := types.DecodeAddress(addr)
	return err == nil
}
