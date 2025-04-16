package token

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("secret_key")

func CreateToken(email string) (string, error) {
	tokenData := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
		"isa":   time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := tokenData.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
