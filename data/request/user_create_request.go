package request

import "time"

type UserCreateRequest struct {
	Name        string    `json:"name" validate:"required,min=2,max=100"`
	Surname     string    `json:"surname" validate:"required,min=2,max=100"`
	Email       string    `json:"email" validate:"required,email"`
	PhoneNumber string    `json:"phone_number" validate:"required,min=10,max=15"`
	CreatedAt   time.Time `json:"created_at"`
}