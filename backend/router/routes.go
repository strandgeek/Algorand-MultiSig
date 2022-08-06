package router

import (
	"multisigdb-svc/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORS(c *gin.Context) {

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {

		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(CORS)

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
