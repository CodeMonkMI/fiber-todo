package todo

import (
	"github.com/CodeMonkMI/fiber-todo/src/auth"
	"github.com/gofiber/fiber/v2"
)

func getAll(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(auth.UserModel)
	todos, err := Find(user.ID)
	if err != nil {
		return ctx.Status(err.Code).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	return ctx.JSON(todos)
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

	user := ctx.Locals("user").(auth.UserModel)
	if todoData.CreatedBy != user.ID {
		return ctx.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"message": "You are not authorized!",
				"todo":    todoData,
				"user":    user,
			})

	}
	return ctx.JSON(todoData)
}
func createTodo(ctx *fiber.Ctx) error {

	var todoBody TodoCreateRequest
	ctx.BodyParser(&todoBody)

	if errors := todoCreateValidator(todoBody); errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	user := ctx.Locals("user").(auth.UserModel)
	todoData := TodoModel{
		Title:     todoBody.Title,
		Completed: false,
		CreatedBy: user.ID,
	}

	newTodData, err2 := create(todoData)

	if err2 != nil {
		return ctx.Status(err2.Code).JSON(
			fiber.Map{
				"message": err2.Error(),
			})
	}

	return ctx.Status(fiber.StatusCreated).JSON(newTodData)
}
func updateTodo(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	var todoBody TodoUpdateRequest
	ctx.BodyParser(&todoBody)

	user := ctx.Locals("user").(auth.UserModel)
	if errors := todoUpdateValidator(todoBody, int(id), user.ID); errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
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

	return ctx.Status(fiber.StatusAccepted).JSON(todo)
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
