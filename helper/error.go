package helper

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Errors  []ValidationError `json:"errors,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Errors: %v", e.Code, e.Message, e.Errors)
}

func NewErrorResponse(code int, message string, validationErrors []ValidationError) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
		Errors:  validationErrors,
	}
}

func FormatValidationError(err error) *ErrorResponse {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		var validationErrors []ValidationError
		for _, fieldError := range ve {
			validationErrors = append(validationErrors, ValidationError{
				Field:   fieldError.Field(),
				Tag:     fieldError.Tag(),
				Message: fmt.Sprintf("Field '%s' failed validation on the '%s' tag", fieldError.Field(), fieldError.Tag()),
			})
		}
		return NewErrorResponse(400, "Validation failed", validationErrors)
	}

	return NewErrorResponse(500, "Internal server error", nil)
}

func HandleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}