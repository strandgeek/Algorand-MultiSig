package signedtransactionctrl

import (
	"multisigdb-svc/service"
	"multisigdb-svc/service/signedtransactionsvc"
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
	me, err := apiutil.GetMe(ctx)
	if err != nil {
		apiutil.Abort(ctx, http.StatusUnauthorized)
	}

	input.SignerId = me.Id

	msa, err := ctrl.svc.SignedTransaction.Create(input)
	if err != nil {
		status := http.StatusBadRequest
		if err == signedtransactionsvc.ErrAlreadyExists {
			status = http.StatusConflict
		}
		apiutil.Abort(ctx, status)
		return
	}
	ctx.JSON(200, msa)
}
