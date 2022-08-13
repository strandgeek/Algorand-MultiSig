package api

import (
	"multisigdb-svc/controller/authctrl"
	"multisigdb-svc/controller/multisigaccountctrl"
	"multisigdb-svc/controller/signedtransactionctrl"
	"multisigdb-svc/controller/transactionctrl"
	"multisigdb-svc/middlewares"
	"multisigdb-svc/service"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetupApi(db *gorm.DB, logger *zap.Logger, cache *cache.Cache) (*gin.Engine, error) {
	api := gin.Default()

	// Service
	svc := service.NewService(db, cache)

	// Global Middlewares
	m := middlewares.NewMiddlewares(db)
	api.Use(m.Cors())
	api.Use(m.Me())

	ms := api.Group("ms-multisig")
	{
		v1 := ms.Group("v1")
		{
			// Auth routes
			authCtrl := authctrl.NewAuthController(svc)
			v1.POST("/auth/nonce", authCtrl.GenerateNonce)
			v1.POST("/auth/complete", authCtrl.Auth)
			v1.GET("/auth/me", authCtrl.Me)

			// MultisigAccount routes
			msaCtrl := multisigaccountctrl.NewMultiSigAccountController(svc)
			v1.POST("/multisig-accounts", msaCtrl.Create)
			v1.GET("/multisig-accounts", msaCtrl.List)
			v1.GET("/multisig-accounts/:msAddress", msaCtrl.Get)
			v1.GET("/multisig-accounts/:msAddress/transactions", msaCtrl.GetTransactions)

			// Transaction routes
			txnCtrl := transactionctrl.NewTransactionController(svc)
			v1.POST("/transactions", txnCtrl.Create)
			v1.GET("/transactions/:txId", txnCtrl.GetByTxId)

			// Signed Transaction routes
			signedTxnCtrl := signedtransactionctrl.NewSignedTransactionController(svc)
			v1.POST("/signed-transactions", signedTxnCtrl.Create)
		}
	}

	return api, nil
}
