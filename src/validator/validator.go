package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(data interface{}) []string {
	var errors []string
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, formateValidationError(err))
		}
	}
	return errors
}

func formateValidationError(err validator.FieldError) string {
	field := err.Field()

	switch err.Tag() {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be valid email"
	case "min":
		return field + " must be at least" + err.Param() + "chars"
	case "max":
		return field + " must not exceed" + err.Param() + "chars"
	default:
		return field + " is invalid"
	}

}
