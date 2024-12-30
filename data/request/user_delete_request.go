package request

import "github.com/google/uuid"

type UserDeleteRequest struct {
	Id uuid.UUID `json:"id" validate:"required,uuid"` 
}