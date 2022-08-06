package authctrl

import (
	"encoding/base64"
	"multisigdb-svc/utils"
	"time"

	"github.com/algorand/go-algorand-sdk/encoding/msgpack"
	"github.com/algorand/go-algorand-sdk/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
)

type AuthController struct {
	Cache *cache.Cache
}

type GenerateNonceRequest struct {
	Address string `json:"address"`
}

type GenerateNoncePayload struct {
	Nonce string `json:"nonce"`
}

// GenerateNonce generates a unique nonce to be used on authentication
func (h AuthController) GenerateNonce(c *gin.Context) {
	req := GenerateNonceRequest{}
	c.BindJSON(&req)
	if req.Address == "" || !utils.IsValidAddress(req.Address) {
		c.AbortWithStatusJSON(500, map[string]string{
			"error": "invalid address",
		})
		return
	}
	uid, err := uuid.NewRandom()
	if err != nil {
		c.AbortWithStatusJSON(500, map[string]string{
			"error": "could not generate nonce",
		})
		return
	}
	nonce := uid.String()
	nonceKey := "AUTH_NONCE:" + req.Address
	h.Cache.Set(nonceKey, nonce, time.Minute*5)
	c.JSON(200, GenerateNoncePayload{
		Nonce: nonce,
	})
}

type AuthRequest struct {
	// SignedTxBase64 is the signed transaction in base64
	SignedTxBase64 string `json:"signed_tx_base64"`

	// PubKey is the signer algorand public address
	PubKey string `json:"pub_key"`
}

type AuthPayload struct {
	Token string `json:"token"`
}

// Auth use the signed tx to authenticate user account with a JWT token
func (h AuthController) Auth(c *gin.Context) {
	req := AuthRequest{}
	c.BindJSON(&req)
	if req.SignedTxBase64 == "" {
		c.AbortWithStatusJSON(500, map[string]string{
			"error": "invalid signed_tx_base64",
		})
		return
	}
	stxn := types.SignedTxn{}
	recoveredTxBytes := make([]byte, 1e3)
	base64.StdEncoding.Decode(recoveredTxBytes, []byte(req.SignedTxBase64))
	msgpack.Decode(recoveredTxBytes, &stxn)
	pubkey, err := utils.GetPubKey(req.PubKey)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	nonceKey := "AUTH_NONCE:" + req.PubKey
	nonce, nonceExists := h.Cache.Get(nonceKey)
	if !nonceExists {
		c.AbortWithStatusJSON(401, map[string]string{
			"error": "could not retrieve nonce for pub_key",
		})
		return
	}
	ret := utils.RawVerifyTransaction(pubkey, stxn.Txn, stxn.Sig[:], nonce.(string))
	if !ret {
		c.AbortWithStatusJSON(401, map[string]string{
			"error": "unauthorized",
		})
		return
	}
	token, err := utils.CreateAccountJWT(req.PubKey)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, AuthPayload{
		Token: token,
	})
}

// Me retrieve the logged user account
func (h AuthController) Me(c *gin.Context) {
	me, err := utils.GetMe(c)
	if err != nil {
		c.AbortWithError(401, err)
		return
	}
	c.JSON(200, me)
}
