package middlewares

import (
	"multisigdb-svc/utils/apiutil"

	"github.com/gin-gonic/gin"
)

func (m *Middlewares) RequireMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		me, err := apiutil.GetMe(c)
		if err != nil || me == nil {
			apiutil.Abort(c, 401)
			return
		}
		c.Next()
	}
}
