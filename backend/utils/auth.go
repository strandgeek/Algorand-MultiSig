package utils

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base32"
	"errors"

	"github.com/algorand/go-algorand-sdk/encoding/msgpack"
	"github.com/algorand/go-algorand-sdk/types"
)

func GetPubKey(address string) (ed25519.PublicKey, error) {
	checksumLenBytes := 4
	decoded, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(address)
	if err != nil {
		return nil, errors.New("could not decode algo address")
	}
	if len(decoded) != len(types.Address{})+checksumLenBytes {
		return nil, errors.New("decoded algo address wrong length")
	}
	addressBytes := decoded[:len(types.Address{})]
	return addressBytes, nil
}

func RawVerifyTransaction(pubkey ed25519.PublicKey, transaction types.Transaction, sig []byte, nonce string) bool {
	if len(nonce) == 0 {
		return false
	}
	note := transaction.Note
	nonceMessage := []byte("Authentication. Nonce: " + nonce)
	if !bytes.Equal(note, nonceMessage) {
		return false
	}
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
