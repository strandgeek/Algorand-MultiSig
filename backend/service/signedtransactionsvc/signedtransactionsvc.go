package signedtransactionsvc

import (
	"encoding/base64"
	"errors"
	"multisigdb-svc/model"
	"multisigdb-svc/utils"
	"multisigdb-svc/utils/algoutil"
	"multisigdb-svc/utils/dbutil"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/encoding/msgpack"
	"github.com/algorand/go-algorand-sdk/types"
	"gorm.io/gorm"
)

type SignedTransactionService struct {
	db *gorm.DB
}

func NewSignedTransactionService(db *gorm.DB) *SignedTransactionService {
	return &SignedTransactionService{
		db: db,
	}
}

type ListFilter struct {
	MultiSigAccountId *int64
}

type CreateInput struct {
	SignerId                   int64  `json:"-"`
	TransactionTxnId           string `json:"transaction_txn_id"`
	RawSignedTransactionBase64 string `json:"raw_signed_transaction_base_64"`
}

func (s *SignedTransactionService) Create(input CreateInput) (*model.SignedTransaction, error) {
	decodedSignedTxn := types.SignedTxn{}
	decodedSignedTxnBytes := make([]byte, 1e3)

	if _, err := base64.StdEncoding.Decode(decodedSignedTxnBytes, []byte(input.RawSignedTransactionBase64)); err != nil {
		return nil, err
	}

	if err := msgpack.Decode(decodedSignedTxnBytes, &decodedSignedTxn); err != nil {
		return nil, err
	}

	var txn model.Transaction
	if err := s.db.Where("txn_id = ?", input.TransactionTxnId).Preload("MultiSigAccount.Accounts").First(&txn).Error; err != nil {
		return nil, errors.New("could not get transaction")
	}

	if txn.TxnId != crypto.GetTxID(decodedSignedTxn.Txn) {
		return nil, errors.New("signed txn is not equal selected txn")
	}

	var signer model.Account
	if err := s.db.Where("id = ?", input.SignerId).First(&signer).Error; err != nil {
		return nil, errors.New("could not get signer")
	}

	stx := model.SignedTransaction{
		RawSignedTransaction: input.RawSignedTransactionBase64,
		TransactionId:        txn.Id,
		SignerId:             input.SignerId,
	}

	signerIndex := dbutil.GetSignerIndex(txn.MultiSigAccount.Accounts, signer.Address)
	if signerIndex == -1 {
		return nil, errors.New("signer is not valid for this multisig account")
	}

	pubkey, _ := utils.GetPubKey(signer.Address)
	decodedTxn := decodedSignedTxn.Txn
	subsigSignature := decodedSignedTxn.Msig.Subsigs[signerIndex].Sig[:]
	isSignatureValid := algoutil.VerifySignedTransaction(pubkey, decodedTxn, subsigSignature)

	if !isSignatureValid {
		return nil, errors.New("could not validate signature for signer")
	}

	if err := s.db.Create(&stx).Error; err != nil {
		return nil, err
	}

	return &stx, nil
}
