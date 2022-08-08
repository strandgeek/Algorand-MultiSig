package api

import (
	"multisigdb-svc/controller/authctrl"
	"multisigdb-svc/controller/multisigaccountctrl"
	"multisigdb-svc/controller/transactionctrl"
	"multisigdb-svc/middlewares"
	"multisigdb-svc/model"
	"multisigdb-svc/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupApi() (*gin.Engine, error) {
	api := gin.Default()

	// Database
	db, err := gorm.Open(sqlite.Open("data/sqlite.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := Migrate(db); err != nil {
		return nil, err
	}

	// Service
	svc := service.NewService(db)

	// General Cache
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Global Middlewares
	m := middlewares.NewMiddlewares(db)
	api.Use(m.Cors())
	api.Use(m.Me())

	ms := api.Group("ms-multisig")
	{
		v1 := ms.Group("v1")
		{
			// Auth routes
			authCtrl := authctrl.AuthController{
				Cache: c,
			}
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
		}
	}

	return api, nil
}

func Migrate(db *gorm.DB) error {
	// TODO: Use a migrations tool like golang-migrate
	return db.AutoMigrate(
		&model.Account{},
		&model.Transaction{},
		&model.SignedTxn{},
		&model.MultiSigAccount{},
	)
}
