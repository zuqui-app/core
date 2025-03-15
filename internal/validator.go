package internal

import (
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New(validator.WithRequiredStructEnabled())

func ValidationErrorsToMap(errors validator.ValidationErrors) map[string]string {
	result := make(map[string]string)
	for _, err := range errors {
		result[err.Field()] = err.Tag()
	}

	return result
}
