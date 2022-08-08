package transactionsvc

import (
	"encoding/base64"
	"errors"
	"multisigdb-svc/model"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/encoding/msgpack"
	"github.com/algorand/go-algorand-sdk/types"
	"gorm.io/gorm"
)

type TransactionService struct {
	db *gorm.DB
}

func NewTransactionService(db *gorm.DB) *TransactionService {
	return &TransactionService{
		db: db,
	}
}

type CreateInput struct {
	MultiSigAccountAddress string `json:"multisig_account_address"`
	RawTransactionBase64   string `json:"raw_transaction_base_64"`
}

func (s *TransactionService) Create(input CreateInput) (*model.Transaction, error) {
	decodedTxn := types.Transaction{}
	decodedTxnBytes := make([]byte, 1e3)

	if _, err := base64.StdEncoding.Decode(decodedTxnBytes, []byte(input.RawTransactionBase64)); err != nil {
		return nil, err
	}

	if err := msgpack.Decode(decodedTxnBytes, &decodedTxn); err != nil {
		return nil, err
	}

	var msa model.MultiSigAccount
	if err := s.db.Where("address = ?", input.MultiSigAccountAddress).First(&msa).Error; err != nil {
		return nil, errors.New("could not get multisig account")
	}

	txId := crypto.TransactionIDString(decodedTxn)

	tx := model.Transaction{
		RawTransaction:    input.RawTransactionBase64,
		TxnId:             txId,
		MultiSigAccountId: msa.Id,
	}

	if err := s.db.Create(&tx).Error; err != nil {
		return nil, err
	}

	return &tx, nil
}
