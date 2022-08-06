package service

import (
	"multisigdb-svc/service/multisigaccountsvc"

	"gorm.io/gorm"
)

type Service struct {
	MultiSigAccount *multisigaccountsvc.MultiSigAccountService
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		MultiSigAccount: multisigaccountsvc.NewMultiSigAccountService(db),
	}
}
