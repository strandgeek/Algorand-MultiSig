package model

type Account struct {
	Address string `gorm:"address,primaryKey"`
}
