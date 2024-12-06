package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmscatena/Fatec_Sert_SGLab/config"
	"github.com/jmscatena/Fatec_Sert_SGLab/dto/models/administrativo"
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
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid": user.UID.String(),                // Subject (user identifier)
		"name": user.Nome,                        // Issuer
		"exp":  time.Now().Add(time.Hour).Unix(), // Expiration time
		//"exp":  time.Now().Add(time.Duration(expire) * time.Minute).Unix(), // Expiration time
		//"aud":  getRole(username),                // Audience (user role)
		//"iat":  time.Now().Unix(),                // Issued at
	})
	fmt.Printf("Token claims added: %+v\n", claims)
	tokenString, err := claims.SignedString([]byte(secretkey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func StoreToken(key string, value string, expire int) error {
	redisClient, err := config.database.InitDF()
	if err != nil {
		return fmt.Errorf("Error Data storing token: %w", err)
	}
	err = redisClient.Set(key, value, time.Minute*time.Duration(expire)).Err()
	if err != nil {
		return fmt.Errorf("Error storing token: %w", err)
	}
	defer redisClient.Close()
	return nil
}

/* Funcao VerifyToken ok */
func VerifyToken(tokenString string, secretKey string) (*jwt.Token, error) {
	keyfunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	}
	/*
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		println(token.Valid)
		if !token.Valid {
			return nil, fmt.Errorf("invalid token")
		}*/

	return token, nil
}

func RevokeToken(token string) error {
	redisClient, err := config.database.InitDF()
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
