package request

import "github.com/google/uuid"

type UserUpdateRequest struct {
	Id          uuid.UUID `json:"id" validate:"required,uuid"`
	Name        string    `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Surname     string    `json:"surname,omitempty" validate:"omitempty,min=2,max=100"`
	Email       string    `json:"email,omitempty" validate:"omitempty,email"`
	PhoneNumber string    `json:"phone_number,omitempty" validate:"omitempty,min=10,max=15"`
}
