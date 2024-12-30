package helper

import (
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		return FormatValidationError(err)
	}
	return nil
}