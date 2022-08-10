package service

import (
	"multisigdb-svc/service/multisigaccountsvc"
	"multisigdb-svc/service/signedtransactionsvc"
	"multisigdb-svc/service/transactionsvc"

	"gorm.io/gorm"
)

type Service struct {
	MultiSigAccount   *multisigaccountsvc.MultiSigAccountService
	Transaction       *transactionsvc.TransactionService
	SignedTransaction *signedtransactionsvc.SignedTransactionService
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		MultiSigAccount:   multisigaccountsvc.NewMultiSigAccountService(db),
		Transaction:       transactionsvc.NewTransactionService(db),
		SignedTransaction: signedtransactionsvc.NewSignedTransactionService(db),
	}
}
