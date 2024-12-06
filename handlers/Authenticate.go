package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmscatena/Fatec_Sert_SGLab/dto/models/administrativo"
	"net/http"
	"strings"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		userID := c.Request.Header.Get("ID")

		if authHeader == "" || userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization"})
			c.Abort()
			return
		}
		// Split the header value into "Bearer " and the token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization"})
			c.Abort()
			return
		}
		// Verify the JWT token
		tokenString := tokenParts[1]
		_, err := VerifyToken(tokenString, "refresh")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or Expired Access"})
			c.Abort()
			return
		}
		condition := "UID=?"
		foundUser, err := Get[administrativo.Usuario](new(administrativo.Usuario), condition, userID)
		//		foundUser := (*foundUsers)[0]
		token, err := ValidateSession(tokenString, *foundUser)

		// Store the user ID in the request context
		//c.JSON(http.StatusOK, gin.H{"data": foundUser.UID, "token": token})
		c.Set("id", foundUser.UID)
		c.Set("data", token)

		// Proceed to the next handler
		c.Next()

	}
}

func ValidateSession(token string, user administrativo.Usuario) (string, error) {
	redisClient, err := infra.database.InitDF()
	if err != nil {
		return "", fmt.Errorf("Error Data revoke token: %w", err)
	}
	userId, err := redisClient.Get(token).Result()
	if err != nil {
		return "", fmt.Errorf("Error validate session: %w", err)
	}
	if userId == "" {
		token, err := redisClient.Get(user.UID.String()).Result()
		if err != nil || token == "" {
			return "", fmt.Errorf("Error validate session: %w", err)
		}
		_, err = VerifyToken(token, "token")
		if err != nil {
			RevokeToken(user.UID.String())
			return "", fmt.Errorf("Error validate session: %w", err)
		}
		refreshtk, err := CreateToken(user, 10, "refresh")
		if err != nil {
			return "", fmt.Errorf("Error validate session: %w", err)
		}
		err = StoreToken(refreshtk, user.UID.String(), 10)
		if err != nil {
			return "", fmt.Errorf("Error validate session: %w", err)
		}
		token = refreshtk
	}
	if userId == user.UID.String() {
		tk, err := redisClient.Get(user.UID.String()).Result()
		if err != nil || tk == "" {
			RevokeToken(token)
			RevokeToken(tk)
			return "", fmt.Errorf("Error validate session: %w", err)
		}
	}
	defer redisClient.Close()
	return token, nil
}
