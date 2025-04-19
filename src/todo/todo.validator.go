package todo

import (
	"github.com/CodeMonkMI/fiber-todo/src/validator"
)

func todoCreateValidator(data TodoCreateRequest) []string {
	if errs := validator.ValidateStruct(data); len(errs) > 0 {
		return errs
	}
	return nil
}
func todoUpdateValidator(data TodoUpdateRequest, id int, userId uint) []string {

	var errorData []string

	// find todo with id
	todo, err := findById(uint(id))

	if err != nil {
		errorData = append(errorData, "Invalid params id")
		return errorData
	}

	if todo.CreatedBy != userId {
		errorData = append(errorData, "You are not authorize!")
		return errorData
	}

	if errs := validator.ValidateStruct(data); len(errs) > 0 {
		return errs
	}
	return nil
}
