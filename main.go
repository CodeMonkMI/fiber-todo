package main

import (
	"github.com/CodeMonkMI/fiber-todo/src/database"
	"github.com/CodeMonkMI/fiber-todo/src/todo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Todo App",
	})
	database.ConnectDB()

	app.Use(logger.New())

	app.Route("/todo", todo.TodoHandlers, "todo")

	app.Listen(":3000")
}
