package todo

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Find() ([]TodoModel, *fiber.Error) {

	var todos []TodoModel
	db.Find(&todos)

	return todos, nil
}

func findById(id uint) (TodoModel, *fiber.Error) {

	var todoData TodoModel
	result := db.First(&todoData, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return TodoModel{}, fiber.NewError(fiber.StatusNotFound, "Record not found")
		}
		return TodoModel{}, fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}
	return todoData, nil
}

func create(data TodoModel) (TodoModel, *fiber.Error) {
	result := db.Create(&data)

	if result.Error != nil {
		return TodoModel{}, fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}

	return data, nil
}

func update(id uint, data TodoModel) (TodoModel, *fiber.Error) {
	result := db.Model(&TodoModel{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return TodoModel{}, fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}

	var findTodo TodoModel
	db.First(&findTodo, id)

	return findTodo, nil

}

func remove(id uint) *fiber.Error {
	result := db.Delete(&TodoModel{}, id)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}
	return nil
}
