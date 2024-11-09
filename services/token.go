package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmscatena/Fatec_Sert_SGLab/database/models/administrativo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

/* Funcao NewAccessToken ok */
func createToken(user administrativo.Usuario) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error Loading Configuration File")
	}
	secretkey := os.Getenv("TOKEN_SECRET_KEY")

	claims := jwt.MapClaims{}
	claims["uuid"] = user.ID
	claims["name"] = user.Nome
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix() // Token valid for 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretkey)
}

/* Funcao VerifyToken ok */
func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("TOKEN_SECRET_KEY"), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token, nil
}

func revokeToken(tokenString string) error {
	/*ctx := context.Background()
	err := s.redisClient.Del(ctx, tokenString).Err()
	if err != nil {
		return fmt.Errorf("error deleting token from Redis: %w", err)
	}*/
	return nil
}
