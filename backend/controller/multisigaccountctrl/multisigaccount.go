package multisigaccountctrl

import (
	"fmt"
	"multisigdb-svc/service"
	"multisigdb-svc/service/multisigaccountsvc"
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
func (ctrl MultiSigAccountController) List(ctx *gin.Context) {
	msa, err := ctrl.svc.MultiSigAccount.List(&multisigaccountsvc.ListFilter{}, paginateutil.NewPaginateFromApi(ctx))
	fmt.Println(err)
	if err != nil {
		apiutil.Abort(ctx, http.StatusBadRequest)
		return
	}
	ctx.JSON(200, msa)
}

// Get a MultiSig Account by Address
func (ctrl MultiSigAccountController) Get(ctx *gin.Context) {
	address, exists := ctx.Params.Get("msAddress")
	if !exists {
		apiutil.Abort(ctx, http.StatusBadRequest)
		return
	}
	msa, err := ctrl.svc.MultiSigAccount.GetByAddress(address)
	fmt.Println(err)
	if err != nil {
		apiutil.Abort(ctx, http.StatusBadRequest)
		return
	}
	ctx.JSON(200, msa)
}
