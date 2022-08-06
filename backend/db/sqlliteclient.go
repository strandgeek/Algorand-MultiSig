package db

import (
	"multisigdb-svc/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DbConnection *gorm.DB

func InitiateDbClient() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("data/sqlite.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := Migrate(db); err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	// TODO: Use a migrations tool like golang-migrate
	return db.AutoMigrate(
		&model.Account{},
		&model.RawTxn{},
		&model.SignedTxn{},
	)
}
