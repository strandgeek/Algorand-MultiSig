package model

type Transaction struct {
	Model
	RawTransaction          string              `json:"raw_transaction" gorm:"column:raw_transaction"`
	TxnId                   string              `json:"txn_id" gorm:"column:txn_id"`
	Status                  string              `json:"status" gorm:"status"`
	MultiSigAccountId       int64               `json:"multisig_account_id" gorm:"multisig_account_id"`
	MultiSigAccount         MultiSigAccount     `json:"-"`
	SignedTransactions      []SignedTransaction `json:"signed_transactions"`
	SignedTransactionsCount uint8               `json:"signed_transactions_count"` // Store the count of signatures for convenience
}
