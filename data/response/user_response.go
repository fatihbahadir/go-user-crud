package response

import (
	"time"
	"github.com/google/uuid"
)

type UserResponse struct {
	Id          uuid.UUID `json:"id"`           
	Name        string    `json:"name"`         
	Surname     string    `json:"surname"`      
	Email       string    `json:"email"`        
	PhoneNumber string    `json:"phone_number"` 
	CreatedAt   time.Time `json:"created_at"`   
}