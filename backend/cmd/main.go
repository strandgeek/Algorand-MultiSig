package main

import (
	"multisigdb-svc/db"
	"multisigdb-svc/router"
	"multisigdb-svc/service"
	"multisigdb-svc/utils"

	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	utils.LoadViperConfig()
	var err error
	logger := utils.GetLoggerInstance()

	logger.Info("Multi-sig go service starting ...")

	db.DbConnection, err = db.InitiateDbClient()
	if err != nil {
		logger.Error("Error in opening the connection with Error Message ", zap.Error(err))
		return
	}

	broadCastTxnJob := cron.New()
	broadCastTxnJob.AddFunc("@every 1m", service.BroadCastTheSignedTxn)
	broadCastTxnJob.Start()

	r := router.SetupRouter()
	addr := viper.GetString("server.host") + ":" + viper.GetString("server.port")
	if err = r.Run(addr); err != nil {
		logger.Error("Error while binding the port with the error message ", zap.Error(err))
	}

}
