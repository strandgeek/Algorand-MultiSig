package transactionctrl

import (
	"fmt"
	"multisigdb-svc/model"
	"multisigdb-svc/service"
	"multisigdb-svc/service/transactionsvc"
	"multisigdb-svc/utils/apiutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	svc *service.Service
}

func NewTransactionController(svc *service.Service) *TransactionController {
	return &TransactionController{
		svc: svc,
	}
}

// Create a Transaction
func (ctrl TransactionController) Create(ctx *gin.Context) {
	var input transactionsvc.CreateInput
	validateError := ctx.BindJSON(&input)
	if validateError != nil {
		apiutil.Abort(ctx, http.StatusBadRequest)
		return
	}

	msa, err := ctrl.svc.Transaction.Create(input)
	if err != nil {
		fmt.Println(err)
		apiutil.Abort(ctx, http.StatusBadRequest)
		return
	}
	ctx.JSON(200, msa)
}

func (ctrl *TransactionController) getTransactionByTxIdParam(ctx *gin.Context) (*model.Transaction, error) {
	txId, _ := ctx.Params.Get("txId")
	return ctrl.svc.Transaction.GetTransactionByTxId(txId)
}

// Get Transaction by TxID
func (ctrl TransactionController) GetByTxId(ctx *gin.Context) {
	transaction, err := ctrl.getTransactionByTxIdParam(ctx)
	if err != nil {
		fmt.Println(err)
		apiutil.Abort(ctx, http.StatusBadRequest)
		return
	}
	ctx.JSON(200, transaction)
}
