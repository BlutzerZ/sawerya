package middleware

import (
	"blutzerz/sawerya/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "access token required",
			})
			c.Abort()
			return
		}
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) >= 2 && authHeaderParts[0] != "Bearer" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid type access token",
			})
			c.Abort()
			return
		}
		accessToken := authHeaderParts[1]
		claims, err := helpers.ValidateToken(accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   err,
				"message": "failed to validate",
			})
			c.Abort()
			return
		}
		c.Set("ID", claims.ID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
