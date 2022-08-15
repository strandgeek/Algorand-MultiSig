package multisigaccountctrl_test

import (
	"multisigdb-svc/model"
	"multisigdb-svc/service/multisigaccountsvc"
	"multisigdb-svc/utils/testutil"
	"net/http"
	"testing"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateMultiSigAccountWhenNotLogged(t *testing.T) {
	ts := testutil.CreateTestSuite(t)
	accounts := testutil.GetTestAccounts()
	acc2 := accounts[0]
	acc3 := accounts[0]
	w := ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodPost,
		Path:   "/ms-multisig/v1/multisig-accounts",
		Input: &multisigaccountsvc.CreateInput{
			Version:   1,
			Threshold: 2,
			Addresses: []string{
				acc2.Address.String(),
				acc3.Address.String(),
			},
		},
	})
	assert.Equal(t, 401, w.Code, "Only logged user should be able to create a multisig account")
}

func TestCreateMultiSigAccountWhenMeIsNotASigner(t *testing.T) {
	ts := testutil.CreateTestSuite(t)
	accounts := testutil.GetTestAccounts()
	acc1 := accounts[0]
	acc2 := accounts[1]
	acc3 := accounts[2]
	w := ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodPost,
		Path:   "/ms-multisig/v1/multisig-accounts",
		Input: &multisigaccountsvc.CreateInput{
			Version:   1,
			Threshold: 2,
			Addresses: []string{
				acc2.Address.String(),
				acc3.Address.String(),
			},
		},
		Me: &acc1,
	})
	assert.Equal(t, 400, w.Code, "The logged user should be included in the signer addresses list")
}

func TestCreateMultiSigAccountWhenMeIsASigner(t *testing.T) {
	ts := testutil.CreateTestSuite(t)
	accounts := testutil.GetTestAccounts()
	acc1 := accounts[0]
	acc2 := accounts[1]
	acc3 := accounts[2]
	payload := model.MultiSigAccount{}
	w := ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodPost,
		Path:   "/ms-multisig/v1/multisig-accounts",
		Input: &multisigaccountsvc.CreateInput{
			Version:   1,
			Threshold: 2,
			Addresses: []string{
				acc1.Address.String(),
				acc2.Address.String(),
				acc3.Address.String(),
			},
		},
		Output: &payload,
		Me:     &acc1,
	})
	algoMsa, err := crypto.MultisigAccountWithParams(1, 2, []types.Address{acc1.Address, acc2.Address, acc3.Address})
	assert.NoError(t, err)
	assert.Equal(t, 200, w.Code)
	msaPublicAddr, _ := algoMsa.Address()
	assert.Equal(t, msaPublicAddr.String(), payload.Address)
}

func TestListMultisigAccounts(t *testing.T) {
	ts := testutil.CreateTestSuite(t)
	accounts := testutil.GetTestAccounts()
	acc1 := accounts[0]
	acc2 := accounts[1]
	acc3 := accounts[2]
	// 1 - Creating the Multisig Account
	msa := model.MultiSigAccount{}
	w := ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodPost,
		Path:   "/ms-multisig/v1/multisig-accounts",
		Input: &multisigaccountsvc.CreateInput{
			Version:   1,
			Threshold: 2,
			Addresses: []string{
				acc1.Address.String(),
				acc2.Address.String(),
			},
		},
		Output: &msa,
		Me:     &acc1,
	})
	assert.Equal(t, 200, w.Code)
	// 2 - Get Multisig Account List (when the logged user is acc3)
	list := make([]model.MultiSigAccount, 0)
	w = ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodGet,
		Path:   "/ms-multisig/v1/multisig-accounts",
		Output: &list,
		Me:     &acc3,
	})
	assert.Equal(t, 200, w.Code)
	assert.Len(t, list, 0, "The account 3 should not see the multisig account since it's not a signer")
	// 3 - Get Multisig Account List (when the logged user is acc1)
	list = make([]model.MultiSigAccount, 0)
	w = ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodGet,
		Path:   "/ms-multisig/v1/multisig-accounts",
		Output: &list,
		Me:     &acc1,
	})
	assert.Equal(t, 200, w.Code)
	assert.Len(t, list, 1, "The account 1 should see the multisig account since it's a signer")
}

func TestGetMultisigAccount(t *testing.T) {
	ts := testutil.CreateTestSuite(t)
	accounts := testutil.GetTestAccounts()
	acc1 := accounts[0]
	acc2 := accounts[1]
	acc3 := accounts[2]
	// 1 - Creating the Multisig Account
	msa := model.MultiSigAccount{}
	w := ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodPost,
		Path:   "/ms-multisig/v1/multisig-accounts",
		Input: &multisigaccountsvc.CreateInput{
			Version:   1,
			Threshold: 2,
			Addresses: []string{
				acc1.Address.String(),
				acc2.Address.String(),
			},
		},
		Output: &msa,
		Me:     &acc1,
	})
	assert.Equal(t, 200, w.Code)
	// 2 - Get Multisig Account (when the logged user is acc3)
	w = ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodGet,
		Path:   "/ms-multisig/v1/multisig-accounts/" + msa.Address,
		Me:     &acc3,
	})
	assert.Equal(t, 403, w.Code, "The account 3 should NOT see the multisig account since it's not a signer")
	// 3 - Get Multisig Account (when the logged user is the acc1)
	var msaRes model.MultiSigAccount
	w = ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodGet,
		Path:   "/ms-multisig/v1/multisig-accounts/" + msa.Address,
		Output: &msaRes,
		Me:     &acc1,
	})
	assert.Equal(t, 200, w.Code, "The account 1 should see the multisig account since it's a signer")
	assert.Equal(t, msa.Address, msaRes.Address)
}

func TestListMultisigAccountTransactions(t *testing.T) {
	ts := testutil.CreateTestSuite(t)
	accounts := testutil.GetTestAccounts()
	acc1 := accounts[0]
	acc2 := accounts[1]
	acc3 := accounts[2]
	// 1 - Creating the Multisig Account
	msa := model.MultiSigAccount{}
	w := ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodPost,
		Path:   "/ms-multisig/v1/multisig-accounts",
		Input: &multisigaccountsvc.CreateInput{
			Version:   1,
			Threshold: 2,
			Addresses: []string{
				acc1.Address.String(),
				acc2.Address.String(),
			},
		},
		Output: &msa,
		Me:     &acc1,
	})
	assert.Equal(t, 200, w.Code)
	// 2 - Get Multisig Account (when the logged user is acc3)
	w = ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodGet,
		Path:   "/ms-multisig/v1/multisig-accounts/" + msa.Address + "/transactions",
		Me:     &acc3,
	})
	assert.Equal(t, 403, w.Code, "The account 3 should NOT see the multisig account transactions since it's not a signer")
	// 3 - Get Multisig Account (when the logged user is the acc1)
	var transactions []model.Transaction
	w = ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodGet,
		Path:   "/ms-multisig/v1/multisig-accounts/" + msa.Address + "/transactions",
		Output: &transactions,
		Me:     &acc1,
	})
	assert.Equal(t, 200, w.Code, "The account 1 should see the multisig account since it's a signer")
	assert.Len(t, transactions, 0)
}
