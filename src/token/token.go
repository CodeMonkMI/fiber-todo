package token

import (
	"errors"
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

func ValidateToken(tokenString string) (string, error) {

	// parse the token

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}

	// check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check the expiration
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return "", errors.New("token expired")
			}
		}
		return claims["email"].(string), nil
	}

	return "", errors.New("invalid token")
}
