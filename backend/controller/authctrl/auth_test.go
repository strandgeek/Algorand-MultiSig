package authctrl_test

import (
	"multisigdb-svc/model"
	"multisigdb-svc/service/authsvc"
	"multisigdb-svc/utils/jwtutil"
	"multisigdb-svc/utils/testutil"
	"net/http"
	"testing"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/stretchr/testify/assert"
)

func getNonceForAddress(ts *testutil.TestSuite, addr string) string {
	nonceRes := authsvc.GenerateNoncePayload{}
	ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodPost,
		Path:   "/ms-multisig/v1/auth/nonce",
		Input: &authsvc.GenerateNonceInput{
			Address: addr,
		},
		Output: &nonceRes,
	})
	nonceKey := "AUTH_NONCE:" + addr
	_, nonceExists := ts.Cache.Get(nonceKey)
	assert.True(ts.T, nonceExists)
	return nonceRes.Nonce
}

func authAccount(ts *testutil.TestSuite, acc crypto.Account) string {
	nonce := getNonceForAddress(ts, acc.Address.String())
	input, _ := testutil.GenerateAuthTransactionInput(acc, nonce)
	payload := authsvc.AuthPayload{}
	w := ts.RequestApi(testutil.RequestApiOptions{
		Method: http.MethodPost,
		Path:   "/ms-multisig/v1/auth/complete",
		Input:  input,
		Output: &payload,
	})
	assert.Equal(ts.T, 200, w.Code)
	assert.NotEmpty(ts.T, payload.Token)
	return payload.Token
}

func TestAuth(t *testing.T) {
	ts := testutil.CreateTestSuite(t)
	accounts := testutil.GetTestAccounts()
	acc := accounts[0]
	token := authAccount(ts, acc)
	address, _ := jwtutil.ParseAccountJWT(token)
	assert.Equal(t, acc.Address.String(), address)
}

func TestMe(t *testing.T) {
	ts := testutil.CreateTestSuite(t)
	accounts := testutil.GetTestAccounts()
	acc := accounts[0]
	var me model.Account
	ts.RequestApi(testutil.RequestApiOptions{
		Method: "GET",
		Path:   "/ms-multisig/v1/auth/me",
		Output: &me,
		Me:     &acc,
	})
	assert.Equal(t, acc.Address.String(), me.Address)
}
