package auth

import (
	"fmt"

	"github.com/CodeMonkMI/fiber-todo/src/password"
	"github.com/CodeMonkMI/fiber-todo/src/token"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func findById(id uint) (UserModel, *fiber.Error) {

	var userData UserModel
	result := db.First(&userData, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return UserModel{}, fiber.NewError(fiber.StatusNotFound, "Record not found")
		}
		return UserModel{}, fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}
	return userData, nil
}
func findByEmail(email string) (UserModel, *fiber.Error) {

	var userData UserModel

	result := db.Where("Email = ? ", email).First(&userData)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return UserModel{}, fiber.NewError(fiber.StatusNotFound, "Record not found")
		}
		return UserModel{}, fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}
	return userData, nil
}

func findByUsername(username string) (UserModel, *fiber.Error) {

	var userData UserModel

	result := db.Where("Username = ? ", username).First(&userData)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return UserModel{}, fiber.NewError(fiber.StatusNotFound, "Record not found")
		}
		return UserModel{}, fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}
	return userData, nil
}

func loginUser(email string, pwd string) (string, *fiber.Error) {

	foundUser, findEmailErr := findByEmail(email)

	errorMsg := "Email or password is incorrect!"

	if findEmailErr != nil && findEmailErr.Code != 404 {
		return " ", fiber.NewError(fiber.StatusBadRequest, errorMsg)
	}

	if findEmailErr != nil && findEmailErr.Code == 500 {
		fmt.Println("[loginUser]", findEmailErr.Error())
		return " ", fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error! Try again Later")
	}

	isMatched := password.VerifyPassword(foundUser.Password, pwd)

	if !isMatched {
		return " ", fiber.NewError(fiber.StatusBadGateway, errorMsg)
	}

	tokenData, err := token.CreateToken(foundUser.Email)
	if err != nil {
		return " ", fiber.NewError(fiber.StatusBadRequest, errorMsg)
	}

	return tokenData, nil
}

func create(data UserModel) (UserModel, *fiber.Error) {

	errors := emailAndUserUniqueValidation(data.Email, data.Username)

	if len(errors) > 0 {
		return UserModel{}, fiber.NewError(fiber.StatusInternalServerError, errors...)
	}

	hashPassword, err := password.HashPassword(data.Password)

	if err != nil {
		return UserModel{}, fiber.NewError(fiber.StatusInternalServerError, "password hashed failed")
	}
	data.Password = hashPassword

	result := db.Create(&data)

	if result.Error != nil {
		return UserModel{}, fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}

	return data, nil
}

func emailAndUserUniqueValidation(email string, username string) []string {

	var errors []string

	_, findEmailErr := findByEmail(email)

	if findEmailErr == nil || findEmailErr.Code != 404 {
		errors = append(errors, "Email is already exist!")
	} else if findEmailErr.Code == 500 {
		errors = append(errors, "Data fetched failed!")
	}

	_, findUsernameErr := findByUsername(username)

	if findUsernameErr == nil || findUsernameErr.Code != 404 {
		errors = append(errors, "Username is already exist!")
	} else if findUsernameErr.Code == 500 {
		errors = append(errors, "Data fetched failed!")
	}

	return errors

}
