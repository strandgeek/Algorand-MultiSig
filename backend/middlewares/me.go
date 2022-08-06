package middlewares

import (
	"errors"
	"multisigdb-svc/db"
	"multisigdb-svc/model"
	"multisigdb-svc/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func MeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		parts := strings.Split(authorization, " ")
		if len(parts) != 2 {
			c.AbortWithStatus(401)
			return
		}
		token := parts[1]
		address, err := utils.ParseAccountJWT(token)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		var acc *model.Account
		if err = db.DbConnection.Model(model.Account{Address: address}).First(&acc).Error; err != nil {
			acc = &model.Account{
				Address: address,
			}
			if err := db.DbConnection.Model(model.Account{}).Create(acc).Error; err != nil {
				c.AbortWithError(500, errors.New("could not create address"))
				return
			}
		}
		c.Set("me", acc)
		c.Next()
	}
}
