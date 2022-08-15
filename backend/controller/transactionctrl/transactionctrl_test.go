package transactionctrl_test

import (
	"multisigdb-svc/model"
	"multisigdb-svc/service/multisigaccountsvc"
	"multisigdb-svc/service/transactionsvc"
	"multisigdb-svc/utils/testutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
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
	// 2 - Create a transaction from an unauthorized account
	w = ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodPost,
		Path:   "/ms-multisig/v1/transactions",
		Input: &transactionsvc.CreateInput{
			MultiSigAccountAddress: msa.Address,
			RawTransactionBase64: testutil.GenerateRawPaymentTransactionBase64(
				msa.Address,
				acc2.Address.String(),
				1000000,
			),
		},
		Output: &msa,
		Me:     &acc3,
	})
	assert.Equal(t, 403, w.Code)
	// 2 - Create a transaction from an authorized account (signer)
	w = ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodPost,
		Path:   "/ms-multisig/v1/transactions",
		Input: &transactionsvc.CreateInput{
			MultiSigAccountAddress: msa.Address,
			RawTransactionBase64: testutil.GenerateRawPaymentTransactionBase64(
				msa.Address,
				acc2.Address.String(),
				1000000,
			),
		},
		Output: &msa,
		Me:     &acc1,
	})
	assert.Equal(t, 200, w.Code)
}

func TestGetTransactionByAddress(t *testing.T) {
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
	// 2 - Create a transaction from an authorized account (signer)
	var txn model.Transaction
	w = ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodPost,
		Path:   "/ms-multisig/v1/transactions",
		Input: &transactionsvc.CreateInput{
			MultiSigAccountAddress: msa.Address,
			RawTransactionBase64: testutil.GenerateRawPaymentTransactionBase64(
				msa.Address,
				acc2.Address.String(),
				1000000,
			),
		},
		Output: &txn,
		Me:     &acc1,
	})
	// 3 - Get transaction from an unauthorized account
	w = ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodGet,
		Path:   "/ms-multisig/v1/transactions/" + txn.TxnId,
		Me:     &acc3,
	})
	assert.Equal(t, 403, w.Code)
	// 3 - Get transaction from an unauthorized account
	var txnRes model.Transaction
	w = ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodGet,
		Path:   "/ms-multisig/v1/transactions/" + txn.TxnId,
		Output: &txnRes,
		Me:     &acc1,
	})
	assert.Equal(t, 200, w.Code)
}
