package request

import "github.com/google/uuid"

type UserUpdateRequest struct {
	Id          uuid.UUID `json:"id" validate:"required,uuid"` 
	Name        string    `json:"name,omitempty" validate:"min=2,max=100"`
	Surname     string    `json:"surname,omitempty" validate:"min=2,max=100"`
	Email       string    `json:"email,omitempty" validate:"email"`
	PhoneNumber string    `json:"phone_number,omitempty" validate:"min=10,max=15"`
}