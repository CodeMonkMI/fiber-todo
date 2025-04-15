package main

import (
	"github.com/CodeMonkMI/fiber-todo/src/todo"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Todo App",
	})

	app.Route("/todo", todo.TodoHandlers, "todo")

	app.Listen(":3000")
}
