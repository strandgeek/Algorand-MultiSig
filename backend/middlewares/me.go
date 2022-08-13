package middlewares

import (
	"errors"
	"multisigdb-svc/model"
	"multisigdb-svc/utils/jwtutil"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Middlewares) Me() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		parts := strings.Split(authorization, " ")
		if len(parts) == 2 {
			token := parts[1]
			address, err := jwtutil.ParseAccountJWT(token)
			if err != nil {
				c.Next()
				return
			}
			var acc *model.Account
			if err = m.db.Where("address = ?", address).First(&acc).Error; err != nil {
				acc = &model.Account{
					Address: address,
				}
				if err := m.db.Create(acc).Error; err != nil {
					c.AbortWithError(500, errors.New("could not create address"))
					return
				}
			}
			c.Set("me", acc)
		}
		c.Next()
	}
}
