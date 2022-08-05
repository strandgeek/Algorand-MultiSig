package client

import (
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/client/v2/common"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var logger = zap.L()

func AlgoRandClient() *algod.Client {
	address := viper.GetString("algorand.address")
	apiHeader := viper.GetString("algorand.api_header")
	apiToken := viper.GetString("algorand.api_token")
	commonClient, err := common.MakeClient(address, apiHeader, apiToken)
	if err != nil {
		logger.Error("failed to make common client: with the message ", zap.Error(err))
		return nil
	}
	return (*algod.Client)(commonClient)
}
