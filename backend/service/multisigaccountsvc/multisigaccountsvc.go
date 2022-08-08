package multisigaccountsvc

import (
	"multisigdb-svc/model"
	"multisigdb-svc/utils/algoutil"
	"multisigdb-svc/utils/dbutil"
	"multisigdb-svc/utils/paginateutil"

	"github.com/algorand/go-algorand-sdk/crypto"
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
	Version   uint8    `json:"version"`
	Threshold uint8    `json:"threshold"`
	Addresses []string `json:"addresses"`
}

type ListFilter struct {}

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

	algoMsa, err := crypto.MultisigAccountWithParams(msa.Version, msa.Threshold, algoutil.AccountsToAlgoAddresses(accounts))

	if err != nil {
		return nil, err
	}

	msaAddress, err := algoMsa.Address()

	if err != nil {
		return nil, err
	}

	msa.Address = msaAddress.String()

	if err := s.db.Create(&msa).Error; err != nil {
		return nil, err
	}

	return &msa, nil
}

func (s *MultiSigAccountService) List(filter *ListFilter, paginate *paginateutil.Paginate) ([]model.MultiSigAccount, error) {
	var msaccounts []model.MultiSigAccount

	tx := paginateutil.ApplyGormPaginate(s.db, paginate)
	err := tx.Preload("Accounts").Preload("Transactions").Find(&msaccounts).Error

	return msaccounts, err
}

func (s *MultiSigAccountService) GetByAddress(address string) (*model.MultiSigAccount, error) {
	var msaccount model.MultiSigAccount

	err := s.db.Preload("Accounts").Where("address = ?", address).First(&msaccount).Error

	return &msaccount, err
}
