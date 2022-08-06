package multisigaccountsvc

import (
	"multisigdb-svc/model"
	"multisigdb-svc/utils/dbutil"
	"multisigdb-svc/utils/paginateutil"

	"gorm.io/gorm"
)

type MultiSigAccountService struct {
	db *gorm.DB
}

func NewMultiSigAccountService(db *gorm.DB) *MultiSigAccountService {
	return &MultiSigAccountService{
		db: db,
	}
}

type CreateInput struct {
	Version   int      `json:"version"`
	Threshold int      `json:"threshold"`
	Addresses []string `json:"addresses"`
}

type ListFilter struct {
	Limit  *int `json:"limit"`
	Offset *int `json:"offset"`
}

func (s *MultiSigAccountService) Create(input CreateInput) (*model.MultiSigAccount, error) {
	accounts, err := dbutil.GetOrCreateAccountByAddresses(s.db, input.Addresses)
	if err != nil {
		return nil, err
	}

	msa := model.MultiSigAccount{
		Version:   input.Version,
		Threshold: input.Threshold,
		Accounts:  accounts,
	}

	if err != nil {
		return nil, err
	}

	if err := s.db.Create(&msa).Error; err != nil {
		return nil, err
	}

	return &msa, nil
}

func (s *MultiSigAccountService) List(filter *ListFilter, paginate *paginateutil.Paginate) ([]model.MultiSigAccount, error) {
	var msaccounts []model.MultiSigAccount

	tx := paginateutil.ApplyGormPaginate(s.db, paginate)
	err := tx.Preload("Accounts").Find(&msaccounts).Error

	return msaccounts, err
}
