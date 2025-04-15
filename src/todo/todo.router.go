package todo

import (
	"github.com/gofiber/fiber/v2"
)

func TodoHandlers(api fiber.Router) {
	api.Get("/", getAll)
	api.Get("/:id", single)
	api.Post("/", create)
	api.Patch("/:id", update)
	api.Delete("/:id", remove)
}
