package router

import (
	"multisigdb-svc/controller"
	"multisigdb-svc/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Global Middlewares
	r.Use(middlewares.Cors)

	ms := r.Group("ms-multisig")
	{
		v1 := ms.Group("v1")
		{
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
