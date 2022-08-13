package apiutil

import (
	"errors"
	"multisigdb-svc/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiError struct {
	Status int
	Error  string `json:"error"`
}

func Abort(ctx *gin.Context, code int) {
	ctx.JSON(code, ApiError{
		Status: code,
		Error:  http.StatusText(code),
	})
}

func GetMe(c *gin.Context) (*model.Account, error) {
	me, meExists := c.Get("me")
	if !meExists {
		return nil, errors.New("not authorized")
	}
	return me.(*model.Account), nil
}
