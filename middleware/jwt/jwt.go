package jwt

import (
	"github.com/gin-gonic/gin"
	"gravitation/pkg/err_msg"
	"gravitation/pkg/util"
	"net/http"
	"time"
)

func JWT() (gin.HandlerFunc) {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = err_msg.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = err_msg.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken()
			if err != nil {
				code = err_msg.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = err_msg.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != err_msg.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg": err_msg.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}