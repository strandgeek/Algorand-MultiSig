package utils

import (
	"errors"
	"multisigdb-svc/model"

	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) (*model.Account, error) {
	me, meExists := c.Get("me")
	if !meExists {
		return nil, errors.New("not authorized")
	}
	return me.(*model.Account), nil
}
