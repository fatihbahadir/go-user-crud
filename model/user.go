package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id          uuid.UUID
	Name        string
	Surname     string
	Email       string
	PhoneNumber string
	CreatedAt   time.Time
}
