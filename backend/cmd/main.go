package main

import (
	"multisigdb-svc/api"
	"multisigdb-svc/model"
	"multisigdb-svc/service/broadcastsvc"
	"multisigdb-svc/utils"

	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	utils.LoadViperConfig()
	var err error
	logger := utils.GetLoggerInstance()

	logger.Info("Multi-sig go service starting ...")

	if err != nil {
		logger.Error("Error in opening the connection with Error Message ", zap.Error(err))
		return
	}

	// Database
	db, err := gorm.Open(sqlite.Open("data/sqlite.db"), &gorm.Config{})
	if err != nil {
		logger.Error("Could not open connection with database")
		return
	}
	if err := Migrate(db); err != nil {
		logger.Error("Could not migrate database")
		return
	}

	broadcastService := broadcastsvc.NewBroadcastService(db)
	broadCastTxnJob := cron.New()
	broadCastTxnJob.AddFunc("@every 1m", broadcastService.BroadcastAllSignedTxn)
	broadCastTxnJob.Start()

	api, err := api.SetupApi(db)
	if err != nil {
		logger.Error("Error while starting the API", zap.Error(err))
	}
	addr := viper.GetString("server.host") + ":" + viper.GetString("server.port")
	if err = api.Run(addr); err != nil {
		logger.Error("Error while binding the port with the error message ", zap.Error(err))
	}
}

func Migrate(db *gorm.DB) error {
	// TODO: Use a migrations tool like golang-migrate
	return db.AutoMigrate(
		&model.Account{},
		&model.Transaction{},
		&model.SignedTransaction{},
		&model.MultiSigAccount{},
	)
}
