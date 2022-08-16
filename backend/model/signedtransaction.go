package model

type SignedTransaction struct {
	Model
	RawSignedTransaction string      `json:"raw_signed_transaction" gorm:"column:raw_signed_transaction"`
	TransactionId        int64       `json:"transaction_id" gorm:"transaction_id"`
	Transaction          Transaction `json:"transaction" gorm:"transaction"`
	SignerId             int64       `json:"signer_id" gorm:"signer_id"`
	Signer               Account     `json:"signer" gorm:"signer"`
}
