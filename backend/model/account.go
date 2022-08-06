package model

type Account struct {
	Id      int64  `json:"id" gorm:"column:id"`
	Address string `json:"address" gorm:"uniqueIndex"`
}
