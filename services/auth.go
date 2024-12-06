package services

import (
	gin "github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jmscatena/Fatec_Sert_SGLab/dto/models/administrativo"
	"github.com/jmscatena/Fatec_Sert_SGLab/infra"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

type Token struct {
	secret  string
	refresh string
}

func NewToken() *Token {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error Loading Configuration File")
		return nil
	}
	return &Token{
		secret:  os.Getenv("TOKEN_SECRET_KEY"),
		refresh: os.Getenv("REFRESH_SECRET_KEY"),
	}
}

var validate = validator.New()

/*
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
			userId, err := VerifyToken(tokenString, "refresh")
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
*/
func Signup(conn infra.Connection) gin.HandlerFunc {
	return func(c *gin.Context) {
		//var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user administrativo.Usuario
		secrets := NewToken()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization"})
			c.Abort()
			return
		}
		foundUser, err := Get[administrativo.Usuario](&user, "email=?", user.Email, conn)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			c.Abort()
			return
		}
		if foundUser != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "User Registred"})
			c.Abort()
			return
		}

		userID, err := New[administrativo.Usuario](&user, conn)
		user.UID = userID
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization"})
			c.Abort()
			return
		}
		token, err := infra.CreateToken(user, 1440, secrets.secret)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		err = infra.StoreToken(token, userID.String(), 1440, conn)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Could Not Signup."})
			c.Abort()
			return

		}
		c.JSON(http.StatusOK, gin.H{"data": userID, "token": token})
		//defer cancel()
		//Redirect to login if successfull
		c.Next()
	}

}
func Login(conn infra.Connection) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user administrativo.Usuario
		secrets := NewToken()
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
			return
		}
		password := user.Senha
		condition := "Email=?"
		foundUser, err := Get[administrativo.Usuario](&user, condition, user.Email, conn)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			c.Abort()
			return
		}
		//foundUser := (*foundUsers)[0]
		err = administrativo.VerifyPassword(password, foundUser.Senha)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			c.Abort()
			return
		}

		token, err := infra.CreateToken(*foundUser, 1440, secrets.secret)
		err = infra.StoreToken(token, foundUser.UID.String(), 1440, conn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//refreshtoken, err := CreateToken(*foundUser, 1000, "refresh")
		//err = StoreToken(foundUser.UID.String(), refreshtoken, 10)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": foundUser.UID, "data": token})

	}
}
func Logout(conn infra.Connection) gin.HandlerFunc {
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
		infra.RevokeToken(tokenString, conn)
		infra.RevokeToken(userID, conn)
		c.Redirect(http.StatusFound, "/login")
	}
}
