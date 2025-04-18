package auth

import (
	"github.com/gofiber/fiber/v2"
)

func login(ctx *fiber.Ctx) error {
	var userBody UserModel

	err := ctx.BodyParser(&userBody)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Failed to parse body",
			})
	}

	var errors []string

	if userBody.Email == "" {
		errors = append(errors, "Email is required")

	}
	if userBody.Password == "" {
		errors = append(errors, "Password is required")
	}

	if len(errors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Validation failed",
				"errors":  errors,
			})
	}

	jwtToken, err2 := loginUser(userBody.Email, userBody.Password)

	if err2 != nil {

		return ctx.Status(err2.Code).JSON(
			fiber.Map{
				"message": err2.Error(),
			})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{"token": jwtToken})
}

func register(ctx *fiber.Ctx) error {

	var todoBody UserModel

	err := ctx.BodyParser(&todoBody)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Failed to parse body",
			})
	}

	var errors []string

	if todoBody.Email == "" {
		errors = append(errors, "Email is required")

	}
	if todoBody.Username == "" {
		errors = append(errors, "username is required")
	}
	if todoBody.FullName == "" {
		errors = append(errors, "Fullname is required")
	}
	if todoBody.Password == "" {
		errors = append(errors, "Password is required")
	}

	if len(errors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Validation failed",
				"errors":  errors,
				"body":    todoBody,
			})
	}

	newUser := UserModel{
		Username: todoBody.Username,
		Email:    todoBody.Email,
		FullName: todoBody.FullName,
		Password: todoBody.Password,
	}

	usr, err2 := create(newUser)

	if err2 != nil {
		return ctx.Status(err2.Code).JSON(
			fiber.Map{
				"message": err2.Error(),
			})
	}
	return ctx.Status(fiber.StatusCreated).JSON(RegisterResponse{
		Username: usr.Username,
		Email:    usr.Email,
		FullName: usr.FullName,
	})
}
