package signedtransactionctrl

import (
	"fmt"
	"multisigdb-svc/service"
	"multisigdb-svc/service/signedtransactionsvc"
	"multisigdb-svc/utils"
	"multisigdb-svc/utils/apiutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignedTransactionController struct {
	svc *service.Service
}

func NewSignedTransactionController(svc *service.Service) *SignedTransactionController {
	return &SignedTransactionController{
		svc: svc,
	}
}

// Create a SignedTransaction
func (ctrl SignedTransactionController) Create(ctx *gin.Context) {
	var input signedtransactionsvc.CreateInput
	validateError := ctx.BindJSON(&input)
	if validateError != nil {
		apiutil.Abort(ctx, http.StatusBadRequest)
		return
	}
	me, err := utils.GetMe(ctx)
	if err != nil {
		apiutil.Abort(ctx, http.StatusUnauthorized)
	}

	input.SignerId = me.Id

	msa, err := ctrl.svc.SignedTransaction.Create(input)
	if err != nil {
		fmt.Println(err)
		apiutil.Abort(ctx, http.StatusBadRequest)
		return
	}
	ctx.JSON(200, msa)
}
