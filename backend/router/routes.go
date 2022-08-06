package router

import (
	"multisigdb-svc/controller"
	"multisigdb-svc/controller/authctrl"
	"multisigdb-svc/middlewares"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Global Middlewares
	r.Use(middlewares.Cors)
	r.Use(middlewares.MeMiddleware())

	// General Cache
	c := cache.New(5*time.Minute, 10*time.Minute)

	ms := r.Group("ms-multisig")
	{
		v1 := ms.Group("v1")
		{
			authCtrl := authctrl.AuthController{
				Cache: c,
			}
			v1.POST("/auth/nonce", authCtrl.GenerateNonce)
			v1.POST("/auth/complete", authCtrl.Auth)
			v1.GET("/auth/me", authCtrl.Me)

			v1.POST("/transactions", controller.AddRawTxn)
			v1.GET("/transactions/:txId", controller.GetRawTxn)

			v1.GET("/transactions/:txId/signatures", controller.GetAllSignedTxn)
			v1.POST("/transactions/:txId/signatures", controller.AddSingedTxn)

			// TODO: Verify the real use-case for this route on frontend
			v1.GET("/transaction-signatures/:signedTxId", controller.GetSignedTxn)
		}
	}

	return r
}
