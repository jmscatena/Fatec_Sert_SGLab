package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmscatena/Fatec_Sert_SGLab/database"
	"github.com/jmscatena/Fatec_Sert_SGLab/database/models/administrativo"
	"github.com/jmscatena/Fatec_Sert_SGLab/services"
	"net/http"
	"strings"
)

var validate = validator.New()

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user administrativo.Usuario

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		foundUser, err := services.GetBy[administrativo.Usuario](&user, "email=?", user.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		if foundUser != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "User Registred"})
			return
		}

		userID, err := services.New[administrativo.Usuario](&user)
		user.UID = userID
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		token, err := services.CreateToken(user, 1440, "token")
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		err = services.StoreToken(token, userID.String())
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Could Not Signup."})
			return

		}
		c.JSON(http.StatusOK, gin.H{"data": userID, "token": token})
		//defer cancel()
	}

}
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user administrativo.Usuario
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		condition := "email=?"
		foundUsers, err := services.GetBy[administrativo.Usuario](&user, condition, user.Email)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}
		foundUser := (*foundUsers)[0]
		passwordIsValid := administrativo.VerifyPassword(foundUser.Senha, user.Senha)
		if passwordIsValid != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": passwordIsValid})
			return
		}

		token, err := services.CreateToken(foundUser, 1440, "token")
		err = services.StoreToken(token, foundUser.UID.String())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		refreshtoken, err := services.CreateToken(foundUser, 10, "refreshs")
		err = services.StoreToken(foundUser.UID.String(), refreshtoken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": foundUser.UID, "token": refreshtoken})

	}
}
func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		userID := c.Request.Header.Get("ID")
		if authHeader == "" || userID == "" {
			c.Redirect(http.StatusFound, "/login")
		}
		// Split the header value into "Bearer " and the token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			return
		}
		// Verify the JWT token
		tokenString := tokenParts[1]
		services.RevokeToken(tokenString)
		services.RevokeToken(userID)
		c.Redirect(http.StatusFound, "/login")
	}
}
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		userID := c.Request.Header.Get("ID")
		if authHeader == "" || userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			return
		}
		// Split the header value into "Bearer " and the token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			return
		}
		// Verify the JWT token
		tokenString := tokenParts[1]
		_, err := services.VerifyToken(tokenString, userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}
		condition := "UID=?"
		foundUsers, err := services.GetBy[administrativo.Usuario](new(administrativo.Usuario), condition, userID)
		foundUser := (*foundUsers)[0]
		token, err := ValidateSession(tokenString, foundUser)

		// Store the user ID in the request context
		c.JSON(http.StatusOK, gin.H{"data": foundUser.UID, "token": token})
	}
}

func ValidateSession(token string, user administrativo.Usuario) (string, error) {
	redisClient, err := database.InitDF()
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
		_, err = services.VerifyToken(token, "token")
		if err != nil {
			services.RevokeToken(user.UID.String())
			return "", fmt.Errorf("Error validate session: %w", err)
		}
		refreshtk, err := services.CreateToken(user, 10, "refresh")
		if err != nil {
			return "", fmt.Errorf("Error validate session: %w", err)
		}
		err = services.StoreToken(refreshtk, user.UID.String())
		if err != nil {
			return "", fmt.Errorf("Error validate session: %w", err)
		}
		token = refreshtk
	}
	if userId == user.UID.String() {
		tk, err := redisClient.Get(user.UID.String()).Result()
		if err != nil || tk == "" {
			services.RevokeToken(token)
			services.RevokeToken(tk)
			return "", fmt.Errorf("Error validate session: %w", err)
		}
	}
	defer redisClient.Close()
	return token, nil
}
