package middleware

import (
	gin "github.com/gin-gonic/gin"
	"github.com/jmscatena/Fatec_Sert_SGLab/services"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the JWT token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.String(http.StatusUnauthorized, "Authorization header is missing")
			c.Abort()
			return
		}

		// Split the header value into "Bearer " and the token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.String(http.StatusUnauthorized, "Invalid authorization header format")
			c.Abort()
			return
		}

		// Verify the JWT token
		tokenString := tokenParts[1]
		userId, err := services.VerifyToken(tokenString, "refresh")
		if err != nil {
			c.String(http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
			return
		}

		// Store the user ID in the request context
		c.Set("userId", userId)

		// Proceed to the next handler
		c.Next()
	}
}
