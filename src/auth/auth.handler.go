package auth

import (
	"github.com/gofiber/fiber/v2"
)

func login(ctx *fiber.Ctx) error {

	var authBody LoginRequest

	ctx.BodyParser(&authBody)

	if errors := loginValidator(authBody); errors != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Validation failed",
				"errors":  errors,
			})
	}

	jwtToken, err := loginUser(authBody.Email, authBody.Password)

	if err != nil {
		return ctx.Status(err.Code).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{"token": jwtToken})
}

func register(ctx *fiber.Ctx) error {

	var authBody RegisterRequest

	ctx.BodyParser(&authBody)

	if errors := registerValidator(authBody); errors != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Validation failed",
				"errors":  errors,
			})
	}

	newUser := UserModel{
		Username: authBody.Username,
		Email:    authBody.Email,
		FullName: authBody.FullName,
		Password: authBody.Password,
	}

	usr, err := create(newUser)

	if err != nil {
		return ctx.Status(err.Code).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}
	return ctx.Status(fiber.StatusCreated).JSON(RegisterResponse{
		Username: usr.Username,
		Email:    usr.Email,
		FullName: usr.FullName,
	})
}
