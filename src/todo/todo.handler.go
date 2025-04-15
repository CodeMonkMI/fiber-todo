package todo

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HandlerAttribute struct {
	DB *gorm.DB
}

func (ha HandlerAttribute) getAll(ctx *fiber.Ctx) error {
	var todos []TodoModel
	result := ha.DB.Find(&todos)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": "Failed to get todos",
			})
	}

	return ctx.JSON(fiber.Map{"message": "Success", "data": todos})
}
func (ha HandlerAttribute) single(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Invalid user id",
			})
	}
	var todoData TodoModel
	result := ha.DB.First(&todoData, uint(id))
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		if result.Error == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(
				fiber.Map{
					"message": "Record not found",
				})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": result.Error.Error(),
			})
	}

	return ctx.JSON(fiber.Map{
		"message": "Success",
		"data":    todoData,
	})
}
func (ha HandlerAttribute) create(ctx *fiber.Ctx) error {

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

	result := ha.DB.Create(&todoData)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": "Failed to create todo",
			})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Success", "data": todoData})
}
func (ha HandlerAttribute) update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Invalid user id",
			})
	}

	var todoBody TodoModel

	ctx.BodyParser(&todoBody)

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
		Completed: todoBody.Completed,
	}

	result := ha.DB.Model(&TodoModel{}).Where("id = ?", id).Updates(todoData)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": "Failed to update todo",
			})
	}

	var findTodo TodoModel
	ha.DB.First(&findTodo, id)

	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "Success", "data": findTodo})
}
func (ha HandlerAttribute) remove(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Invalid todo id",
			})
	}

	ha.DB.Delete(&TodoModel{}, uint(id))

	return ctx.Status(fiber.StatusNoContent).JSON(
		fiber.Map{
			"message": "Data deleted successfully",
		})
}
