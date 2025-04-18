package middleware

import (
	"strings"

	"github.com/CodeMonkMI/fiber-todo/src/auth"
	"github.com/CodeMonkMI/fiber-todo/src/token"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {

	return func(c *fiber.Ctx) error {
		// get auth header
		authHeader := c.Get("Authorization")

		// check if header is empty or doesn't start with bearer
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "You are not authorized!",
			})
		}

		// extract token from header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// validate token
		userEmail, err := token.ValidateToken(tokenString)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "You are not authorized!",
			})
		}

		// fetch userinfo form db

		user, userErr := auth.FindByEmail(userEmail)

		if userErr != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "You are not authorized!",
			})
		}

		// set user email for later user
		c.Locals("user", user)

		return c.Next()
	}
}
