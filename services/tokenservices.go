package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmscatena/Fatec_Sert_SGLab/database"
	"github.com/jmscatena/Fatec_Sert_SGLab/database/models/administrativo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

/* Funcao NewAccessToken ok */
func CreateToken(user administrativo.Usuario, expire int, keytype string) (string, error) {
	err := godotenv.Load(".env")
	var secretkey string
	if err != nil {
		log.Fatalf("Error Loading Configuration File")
	}
	if keytype == "token" {
		secretkey = os.Getenv("TOKEN_SECRET_KEY")
	} else {
		secretkey = os.Getenv("REFRESH_SECRET_KEY")
	}

	claims := jwt.MapClaims{}
	claims["uuid"] = user.UID.String()
	claims["name"] = user.Nome
	claims["exp"] = time.Now().Add(time.Duration(expire) * time.Minute).Unix() // Token valid for 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretkey))
}
func StoreToken(key string, value string) error {
	redisClient, err := database.InitDF()
	if err != nil {
		return fmt.Errorf("Error Data storing token: %w", err)
	}
	err = redisClient.Set(key, value, time.Hour*24).Err()
	if err != nil {
		return fmt.Errorf("Error storing token: %w", err)
	}
	defer redisClient.Close()
	return nil
}

/* Funcao VerifyToken ok */
func VerifyToken(tokenString string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func RevokeToken(token string) error {
	redisClient, err := database.InitDF()
	if err != nil {
		return fmt.Errorf("Error Data revoke token: %w", err)
	}
	err = redisClient.Del(token).Err()
	if err != nil {
		return fmt.Errorf("Error revoke token: %w", err)
	}
	defer redisClient.Close()
	return nil
}
