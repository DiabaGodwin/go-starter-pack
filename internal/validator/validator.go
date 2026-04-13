package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

func ValidateStruct(s interface{}) ([]ValidationError, error) {
	err := validate.Struct(s)
	if err == nil {
		return nil, nil
	}

	var errors []ValidationError

	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, ValidationError{
			Field: err.Field(),
			Tag:   err.Tag(),
		})
	}

	return errors, fmt.Errorf("validation error")
}
