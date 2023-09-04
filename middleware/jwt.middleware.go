package middleware

import (
	"blutzerz/sawerya/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get("token")
		if accessToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "access token required",
			})
			return
		}

		claims, err := helpers.ValidateToken(accessToken)
		if err == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			return
		}
		c.Set("ID", claims.ID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
