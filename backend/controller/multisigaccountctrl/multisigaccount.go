package multisigaccountctrl

import (
	"multisigdb-svc/model"
	"multisigdb-svc/service"
	"multisigdb-svc/service/multisigaccountsvc"
	"multisigdb-svc/service/transactionsvc"
	"multisigdb-svc/utils/apiutil"
	"multisigdb-svc/utils/paginateutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MultiSigAccountController struct {
	svc *service.Service
}

func NewMultiSigAccountController(svc *service.Service) *MultiSigAccountController {
	return &MultiSigAccountController{
		svc: svc,
	}
}

// Create a MultiSig Account
func (ctrl MultiSigAccountController) Create(ctx *gin.Context) {
	var input multisigaccountsvc.CreateInput
	validateError := ctx.BindJSON(&input)
	if validateError != nil {
		apiutil.Abort(ctx, http.StatusBadRequest)
		return
	}

	msa, err := ctrl.svc.MultiSigAccount.Create(input)
	if err != nil {
		apiutil.Abort(ctx, http.StatusBadRequest)
		return
	}
	ctx.JSON(200, msa)
}

// List all MultiSig Account
func (ctrl *MultiSigAccountController) List(ctx *gin.Context) {
	msa, err := ctrl.svc.MultiSigAccount.List(&multisigaccountsvc.ListFilter{}, paginateutil.NewPaginateFromApi(ctx))
	if err != nil {
		apiutil.Abort(ctx, http.StatusBadRequest)
		return
	}
	ctx.JSON(200, msa)
}

func (ctrl *MultiSigAccountController) getMultiSigAccountByAddressParam(ctx *gin.Context) (*model.MultiSigAccount, error) {
	address, _ := ctx.Params.Get("msAddress")
	return ctrl.svc.MultiSigAccount.GetByAddress(address)
}

// Get a MultiSig Account by Address
func (ctrl *MultiSigAccountController) Get(ctx *gin.Context) {
	msa, err := ctrl.getMultiSigAccountByAddressParam(ctx)
	if err != nil {
		apiutil.Abort(ctx, http.StatusNotFound)
		return
	}
	ctx.JSON(200, msa)
}

// Get a MultiSig Account by Address
func (ctrl *MultiSigAccountController) GetTransactions(ctx *gin.Context) {
	msa, err := ctrl.getMultiSigAccountByAddressParam(ctx)
	if err != nil {
		apiutil.Abort(ctx, http.StatusNotFound)
		return
	}
	filter := &transactionsvc.ListFilter{
		MultiSigAccountId: &msa.Id,
	}
	paginate := paginateutil.NewPaginateFromApi(ctx)
	txns, err := ctrl.svc.Transaction.List(filter, paginate)
	ctx.JSON(200, txns)
}
