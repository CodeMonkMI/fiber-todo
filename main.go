package main

import (
	"github.com/CodeMonkMI/fiber-todo/src/auth"
	"github.com/CodeMonkMI/fiber-todo/src/database"
	"github.com/CodeMonkMI/fiber-todo/src/todo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Todo App",
	})
	database.ConnectDB()

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173/",
	}))

	app.Route("/todo", todo.TodoHandlers, "todo")
	app.Route("/auth", auth.AuthHandlers, "todo")

	app.Listen(":9000")

}
