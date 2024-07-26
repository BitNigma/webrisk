package router

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateNewToken(id string) (string, error) {

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := jwtToken.SignedString(os.Getenv("SECRET"))
	if err != nil {
		log.Println("can't format jwt token to string", err)
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) error {

	secretKey := os.Getenv("SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
