package password

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (string, error) {
	log.Println("receiving pwd", pwd)
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func VerifyPassword(hash string, password string) bool {
	log.Println("hash password", hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println(err.Error())
	}
	return err == nil

}
