package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"godemo.com/demo/internal/token"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		userId, _ := token.ExtractTokenID(c)
		c.Set("user_id", userId)
		c.Next()
	}
}
