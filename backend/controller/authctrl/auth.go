package authctrl

import (
	"multisigdb-svc/service"
	"multisigdb-svc/service/authsvc"
	"multisigdb-svc/utils/apiutil"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	svc *service.Service
}

func NewAuthController(svc *service.Service) *AuthController {
	return &AuthController{
		svc: svc,
	}
}

// GenerateNonce generates a unique nonce to be used on authentication
func (h AuthController) GenerateNonce(c *gin.Context) {
	input := authsvc.GenerateNonceInput{}
	c.BindJSON(&input)
	payload, err := h.svc.Auth.GenerateNonce(input)
	if err != nil {
		apiutil.Abort(c, 400)
		return
	}
	c.JSON(200, payload)
}

// Auth use the signed tx to authenticate user account with a JWT token
func (h AuthController) Auth(c *gin.Context) {
	input := authsvc.AuthInput{}
	c.BindJSON(&input)
	payload, err := h.svc.Auth.Auth(input)
	if err != nil {
		apiutil.Abort(c, 400)
		return
	}
	c.JSON(200, payload)
}

// Me retrieve the logged user account
func (h AuthController) Me(c *gin.Context) {
	me, err := apiutil.GetMe(c)
	if err != nil {
		c.AbortWithError(401, err)
		return
	}
	c.JSON(200, me)
}
