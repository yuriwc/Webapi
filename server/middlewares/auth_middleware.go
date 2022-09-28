package middlewares

import (
	"net/http"
	"webApi/services"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		token := header[len(Bearer_schema):]

		if !services.NewJWTService("secret-key", "web-api").ValidateToken(token) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
