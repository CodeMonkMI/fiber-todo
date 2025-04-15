package todo

import (
	"github.com/CodeMonkMI/fiber-todo/src/database"
	"github.com/gofiber/fiber/v2"
)

func TodoHandlers(api fiber.Router) {

	db := database.GetDB()
	db.AutoMigrate(&TodoModel{})

	data := HandlerAttribute{DB: db}

	api.Get("/", data.getAll)
	api.Get("/:id", data.single)
	api.Post("/", data.create)
	api.Put("/:id", data.update)
	api.Delete("/:id", data.remove)

}
