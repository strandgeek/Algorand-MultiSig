package main

import (
	"multisigdb-svc/api"
	"multisigdb-svc/utils"

	"github.com/spf13/viper"
	"go.uber.org/zap"
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

	// broadCastTxnJob := cron.New()
	// broadCastTxnJob.AddFunc("@every 1m", service.BroadCastTheSignedTxn)
	// broadCastTxnJob.Start()

	api, err := api.SetupApi()
	if err != nil {
		logger.Error("Error while starting the API", zap.Error(err))
	}
	addr := viper.GetString("server.host") + ":" + viper.GetString("server.port")
	if err = api.Run(addr); err != nil {
		logger.Error("Error while binding the port with the error message ", zap.Error(err))
	}
}
