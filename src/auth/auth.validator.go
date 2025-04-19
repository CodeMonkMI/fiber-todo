package auth

import "github.com/CodeMonkMI/fiber-todo/src/validator"

func registerValidator(data RegisterRequest) []string {

	errs := validator.ValidateStruct(data)
	if len(errs) > 0 {
		return errs
	}

	return nil

}
func loginValidator(data LoginRequest) []string {

	if errs := validator.ValidateStruct(data); len(errs) > 0 {
		return errs
	}

	return nil

}
