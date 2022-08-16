package model

type Account struct {
	Model
	Address string `json:"address" gorm:"uniqueIndex"`
}
