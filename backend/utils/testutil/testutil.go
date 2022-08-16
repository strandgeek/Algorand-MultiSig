package testutil

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"multisigdb-svc/api"
	"multisigdb-svc/model"
	"multisigdb-svc/service"
	"multisigdb-svc/service/authsvc"
	"multisigdb-svc/utils/jwtutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/encoding/msgpack"
	"github.com/algorand/go-algorand-sdk/transaction"
	"github.com/algorand/go-algorand-sdk/types"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TestSuite struct {
	T     *testing.T
	DB    *gorm.DB
	Cache *cache.Cache
	Svc   *service.Service
	Api   *gin.Engine
}

func GetTestAccounts() []crypto.Account {
	_, b, _, _ := runtime.Caller(0)
	// Root folder of this project
	root := filepath.Join(filepath.Dir(b), "../..")
	var accounts []crypto.Account
	accountsFile, err := os.ReadFile(root + "/utils/testutil/testdata/accounts.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(accountsFile, &accounts)
	return accounts
}

func CreateTestSuite(t *testing.T) *TestSuite {
	gin.SetMode(gin.TestMode)
	db, err := gorm.Open(sqlite.Open(":memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&model.Account{},
		&model.Transaction{},
		&model.SignedTransaction{},
		&model.MultiSigAccount{},
	)
	c := cache.New(time.Minute*5, time.Minute)
	logger, _ := zap.NewDevelopment()
	apiEngine := gin.New()
	_ = api.SetupApi(apiEngine, db, logger, c)
	ts := &TestSuite{
		DB:    db,
		Cache: c,
		Api:   apiEngine,
		T:     t,
	}
	return ts
}

func BodyFromInput(input interface{}) io.Reader {
	if input == nil {
		return nil
	}
	inputJson, _ := json.Marshal(input)
	return bytes.NewReader(inputJson)
}

type RequestApiOptions struct {
	// HTTP Method
	Method string

	// HTTP Path
	Path string

	// Input to be marshaled as JSON
	Input interface{}

	// Output pointer to be unmarshalled from JSON response
	Output interface{}

	// Additional headers for the request
	Headers *map[string]string

	// Use a token for an account (authorized routes)
	Me *crypto.Account
}

func (ts *TestSuite) RequestApi(opts RequestApiOptions) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	body := BodyFromInput(opts.Input)
	req, _ := http.NewRequest(opts.Method, opts.Path, body)
	if opts.Headers != nil {
		for key, value := range *opts.Headers {
			req.Header.Set(key, value)
		}
	}
	if opts.Me != nil {
		token, _ := jwtutil.CreateAccountJWT(string(opts.Me.Address.String()))
		req.Header.Set("Authorization", "Bearer "+token)
	}
	ts.Api.ServeHTTP(w, req)
	if opts.Output != nil {
		err := json.Unmarshal(w.Body.Bytes(), &opts.Output)
		if err != nil {
			ts.T.Error(err)
		}
	}
	return w
}

func GenerateAuthTransactionInput(acc crypto.Account, nonce string) (*authsvc.AuthInput, error) {
	addr := acc.Address.String()
	txn, err := transaction.MakePaymentTxnWithFlatFee(
		addr,
		addr,
		0,
		0,
		0,
		0,
		[]byte("Authentication. Nonce: "+nonce),
		"",
		"testnet-v1.0",
		[]byte("SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI="),
	)
	if err != nil {
		return nil, err
	}
	_, stxn, _ := crypto.SignTransaction(acc.PrivateKey, txn)
	var stxnBytes = make([]byte, 1e3)
	base64.StdEncoding.Encode(stxnBytes, stxn)
	stxnBytes = bytes.Trim(stxnBytes, "\x00")
	input := &authsvc.AuthInput{
		SignedTxBase64: string(stxnBytes),
		PubKey:         addr,
	}
	return input, nil
}

func GenerateRawPaymentTransactionBase64(from string, to string, amount uint64) string {
	txn, err := transaction.MakePaymentTxnWithFlatFee(
		from,
		to,
		0,
		amount,
		0,
		0,
		[]byte(""),
		"",
		"testnet-v1.0",
		[]byte("SGO1GKSzyE7IEPItTxCByw9x8FmnrCDexi9/cOUJOiI="),
	)
	if err != nil {
		return ""
	}
	var txnBytes = make([]byte, 1e3)
	base64.StdEncoding.Encode(txnBytes, msgpack.Encode(txn))
	txnBytes = bytes.Trim(txnBytes, "\x00")
	return string(txnBytes)
}

func SignMultisigTransaction(acc crypto.Account, msa crypto.MultisigAccount, rawTransaction string) string {
	b64TxnBytes := []byte(rawTransaction)
	recoveredTxn := types.Transaction{}
	recoveredTxnBytes := make([]byte, 1e3)
	recoveredTxn = types.Transaction{}
	base64.StdEncoding.Decode(recoveredTxnBytes, b64TxnBytes)
	msgpack.Decode(recoveredTxnBytes, &recoveredTxn)
	_, signedTxn, _ := crypto.SignMultisigTransaction(acc.PrivateKey, msa, recoveredTxn)
	var signedTxnBytes = make([]byte, 1e3)
	base64.StdEncoding.Encode(signedTxnBytes, signedTxn)
	signedTxnBytes = bytes.Trim(signedTxnBytes, "\x00")
	return string(signedTxnBytes)
}
