package authsvc

import (
	"encoding/base64"
	"errors"
	"multisigdb-svc/utils/algoutil"
	"multisigdb-svc/utils/jwtutil"
	"time"

	"github.com/algorand/go-algorand-sdk/encoding/msgpack"
	"github.com/algorand/go-algorand-sdk/types"
	"github.com/google/uuid"
	gocache "github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

type AuthService struct {
	db    *gorm.DB
	cache *gocache.Cache
}

func NewAuthService(db *gorm.DB, cache *gocache.Cache) *AuthService {
	return &AuthService{
		db:    db,
		cache: cache,
	}
}

type GenerateNonceInput struct {
	Address string `json:"address"`
}

type GenerateNoncePayload struct {
	Nonce string `json:"nonce"`
}

type AuthInput struct {
	// SignedTxBase64 is the signed transaction in base64
	SignedTxBase64 string `json:"signed_tx_base64"`

	// PubKey is the signer algorand public address
	PubKey string `json:"pub_key"`
}

type AuthPayload struct {
	Token string `json:"token"`
}

func (s *AuthService) GenerateNonce(input GenerateNonceInput) (*GenerateNoncePayload, error) {
	if input.Address == "" || !algoutil.IsValidAddress(input.Address) {
		return nil, errors.New("invalid address")
	}
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	nonce := uid.String()
	nonceKey := "AUTH_NONCE:" + input.Address
	s.cache.Set(nonceKey, nonce, time.Minute*5)
	return &GenerateNoncePayload{
		Nonce: nonce,
	}, nil
}

func (s *AuthService) Auth(input AuthInput) (*AuthPayload, error) {
	stxn := types.SignedTxn{}
	recoveredTxBytes := make([]byte, 1e3)
	base64.StdEncoding.Decode(recoveredTxBytes, []byte(input.SignedTxBase64))
	msgpack.Decode(recoveredTxBytes, &stxn)
	pubkey, err := algoutil.GetPubKey(input.PubKey)
	if err != nil {
		return nil, err
	}
	nonceKey := "AUTH_NONCE:" + input.PubKey
	nonce, nonceExists := s.cache.Get(nonceKey)
	if !nonceExists {
		return nil, errors.New("nonce not valid")
	}
	ret := algoutil.RawVerifyTransaction(pubkey, stxn.Txn, stxn.Sig[:], nonce.(string))
	if !ret {
		return nil, errors.New("unauthorized")
	}
	token, err := jwtutil.CreateAccountJWT(input.PubKey)
	if err != nil {
		return nil, err
	}
	return &AuthPayload{
		Token: token,
	}, nil
}
