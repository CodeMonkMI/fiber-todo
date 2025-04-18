package todo

import (
	"github.com/CodeMonkMI/fiber-todo/src/database"
	"github.com/CodeMonkMI/fiber-todo/src/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func TodoHandlers(api fiber.Router) {

	db = database.GetDB()
	db.AutoMigrate(&TodoModel{})

	api.Use(middleware.AuthMiddleware())
	api.Get("/", getAll)
	api.Get("/:id", single)
	api.Post("/", createTodo)
	api.Put("/:id", updateTodo)
	api.Delete("/:id", removeTodo)

}
