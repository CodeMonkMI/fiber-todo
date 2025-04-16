package todo

import (
	"github.com/gofiber/fiber/v2"
)

func getAll(ctx *fiber.Ctx) error {

	todos, err := Find()
	if err != nil {
		return ctx.Status(err.Code).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	return ctx.JSON(fiber.Map{"message": "Success", "data": todos})
}
func single(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Invalid user id",
			})
	}

	todoData, err2 := findById(uint(id))
	if err2 != nil {
		return ctx.Status(err2.Code).JSON(
			fiber.Map{
				"message": err2.Error(),
			})
	}

	return ctx.JSON(fiber.Map{
		"message": "Success",
		"data":    todoData,
	})
}
func createTodo(ctx *fiber.Ctx) error {

	var todoBody TodoModel

	err := ctx.BodyParser(&todoBody)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Failed to parse body",
			})
	}

	if todoBody.Title == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Title is required",
			})
	}
	todoData := TodoModel{
		Title:     todoBody.Title,
		Completed: false,
	}

	newTodData, err2 := create(todoData)

	if err2 != nil {
		return ctx.Status(err2.Code).JSON(
			fiber.Map{
				"message": err2.Error(),
			})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Success", "data": newTodData})
}
func updateTodo(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Invalid user id",
			})
	}

	var todoBody TodoModel
	err2 := ctx.BodyParser(&todoBody)

	if err2 != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Failed to parse body",
			})
	}

	if todoBody.Title == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Title is required",
			})
	}

	todoData := TodoModel{
		Title:     todoBody.Title,
		Completed: todoBody.Completed,
	}

	todo, err3 := update(uint(id), todoData)
	if err3 != nil {
		return ctx.Status(err3.Code).JSON(
			fiber.Map{
				"message": err3.Error(),
			})
	}

	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Success", "data": todo})
}
func removeTodo(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Invalid todo id",
			})
	}

	err2 := remove(uint(id))
	if err2 != nil {
		return ctx.Status(err2.Code).JSON(
			fiber.Map{
				"message": err2.Error(),
			})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(
		fiber.Map{
			"message": "Data deleted successfully",
		})
}
