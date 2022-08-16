package signedtransactionctrl_test

import (
	"crypto/ed25519"
	"multisigdb-svc/model"
	"multisigdb-svc/service/multisigaccountsvc"
	"multisigdb-svc/service/signedtransactionsvc"
	"multisigdb-svc/service/transactionsvc"
	"multisigdb-svc/utils/testutil"
	"net/http"
	"testing"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/stretchr/testify/assert"
)

func createTransaction(ts *testutil.TestSuite) model.Transaction {
	accounts := testutil.GetTestAccounts()
	acc1 := accounts[0]
	acc2 := accounts[1]
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
	assert.Equal(ts.T, 200, w.Code)
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
	return txn
}

// func TestCreateSignedTransaction_When_Raw_Signed_Is_Empty(t *testing.T) {
// 	ts := testutil.CreateTestSuite(t)
// 	accounts := testutil.GetTestAccounts()
// 	acc1 := accounts[0]
// 	acc2 := accounts[1]
// 	txn := createTransaction(ts)
// 	msa := crypto.MultisigAccount{
// 		Version:   1,
// 		Threshold: 2,
// 		Pks: []ed25519.PublicKey{
// 			acc1.PublicKey,
// 			acc2.PublicKey,
// 		},
// 	}
// 	w := ts.RequestApi(testutil.RequestApiOptions{
// 		Method: http.MethodPost,
// 		Path:   "/ms-multisig/v1/signed-transactions",
// 		Input: &signedtransactionsvc.CreateInput{
// 			TransactionTxnId:           txn.TxnId,
// 			RawSignedTransactionBase64: testutil.SignMultisigTransaction(acc1, msa, txn.RawTransaction),
// 		},
// 		Me: &acc1,
// 	})
// 	assert.Equal(t, 400, w.Code)
// }

func TestCreateSignedTransaction_With_Unauthorized_Signer(t *testing.T) {
	ts := testutil.CreateTestSuite(t)
	accounts := testutil.GetTestAccounts()
	acc1 := accounts[0]
	acc2 := accounts[1]
	acc3 := accounts[2]
	msa := crypto.MultisigAccount{
		Version:   1,
		Threshold: 2,
		Pks: []ed25519.PublicKey{
			acc1.PublicKey,
			acc2.PublicKey,
		},
	}
	txn := createTransaction(ts)
	// 1 - Create Signed Transaction from an unauthorized signer
	w := ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodPost,
		Path:   "/ms-multisig/v1/signed-transactions",
		Input: &signedtransactionsvc.CreateInput{
			TransactionTxnId:           txn.TxnId,
			RawSignedTransactionBase64: testutil.SignMultisigTransaction(acc3, msa, txn.RawTransaction),
		},
		Me: &acc3,
	})
	assert.Equal(t, 400, w.Code)
	// 1 - Create Signed Transaction from account 1
	w = ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodPost,
		Path:   "/ms-multisig/v1/signed-transactions",
		Input: &signedtransactionsvc.CreateInput{
			TransactionTxnId:           txn.TxnId,
			RawSignedTransactionBase64: testutil.SignMultisigTransaction(acc1, msa, txn.RawTransaction),
		},
		Me: &acc1,
	})
	assert.Equal(t, 200, w.Code)
	// 1 - Create Signed Transaction from account 2
	w = ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodPost,
		Path:   "/ms-multisig/v1/signed-transactions",
		Input: &signedtransactionsvc.CreateInput{
			TransactionTxnId:           txn.TxnId,
			RawSignedTransactionBase64: testutil.SignMultisigTransaction(acc2, msa, txn.RawTransaction),
		},
		Me: &acc2,
	})
	assert.Equal(t, 200, w.Code)
	// 2 - Check if the transaction status changed after all required signatures
	var txnRes model.Transaction
	w = ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodGet,
		Path:   "/ms-multisig/v1/transactions/" + txn.TxnId,
		Output: &txnRes,
		Me:     &acc1,
	})
	assert.Equal(t, "READY", txnRes.Status)
}
