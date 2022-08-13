package service

import (
	"multisigdb-svc/service/authsvc"
	"multisigdb-svc/service/multisigaccountsvc"
	"multisigdb-svc/service/signedtransactionsvc"
	"multisigdb-svc/service/transactionsvc"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

type Service struct {
	MultiSigAccount   *multisigaccountsvc.MultiSigAccountService
	Transaction       *transactionsvc.TransactionService
	SignedTransaction *signedtransactionsvc.SignedTransactionService
	Auth              *authsvc.AuthService
}

func NewService(db *gorm.DB, cache *cache.Cache) *Service {
	return &Service{
		MultiSigAccount:   multisigaccountsvc.NewMultiSigAccountService(db),
		Transaction:       transactionsvc.NewTransactionService(db),
		SignedTransaction: signedtransactionsvc.NewSignedTransactionService(db),
		Auth:              authsvc.NewAuthService(db, cache),
	}
}
