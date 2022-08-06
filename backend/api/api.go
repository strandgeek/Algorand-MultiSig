package api

import (
	"multisigdb-svc/controller"
	"multisigdb-svc/controller/authctrl"
	"multisigdb-svc/controller/multisigaccountctrl"
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

	// Global Middlewares
	api.Use(middlewares.Cors)
	// api.Use(middlewares.MeMiddleware())

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

			v1.POST("/transactions", controller.AddRawTxn)
			v1.GET("/transactions/:txId", controller.GetRawTxn)

			v1.GET("/transactions/:txId/signatures", controller.GetAllSignedTxn)
			v1.POST("/transactions/:txId/signatures", controller.AddSingedTxn)

			// TODO: Verify the real use-case for this route on frontend
			v1.GET("/transaction-signatures/:signedTxId", controller.GetSignedTxn)
		}
	}

	return api, nil
}

func Migrate(db *gorm.DB) error {
	// TODO: Use a migrations tool like golang-migrate
	return db.AutoMigrate(
		&model.Account{},
		&model.RawTxn{},
		&model.SignedTxn{},
		&model.MultiSigAccount{},
	)
}
