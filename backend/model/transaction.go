package model

type Transaction struct {
	Id                    int64               `json:"id" gorm:"column:id"`
	RawTransaction        string              `json:"raw_transaction" gorm:"column:raw_transaction"`
	TxnId                 string              `json:"txn_id" gorm:"column:txn_id"`
	NumberOfSignsRequired int64               `json:"number_of_signs_required" gorm:"column:number_of_signs_required"`
	Status                string              `json:"status" gorm:"status"`
	MultiSigAccountId     int64               `json:"multisig_account_id" gorm:"multisig_account_id"`
	MultiSigAccount       MultiSigAccount     `json:"-"`
	SignedTransactions    []SignedTransaction `json:"signed_transactions"`
}
