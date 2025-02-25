package middleware

import (
	"errors"
	"os"

	"github.com/bootcamp-go/ExamenFinalBE3.git/pkg/web"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			c.Abort()
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			c.Abort()
			return
		}
		c.Next()
	}
}
