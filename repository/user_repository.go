package repository

import (
	"context"
	"user-crud/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	Save(ctx context.Context, user model.User) error 
	Update(ctx context.Context, userId uuid.UUID, user model.User) error
	Delete(ctx context.Context, userId uuid.UUID) error
	FindById(ctx context.Context, userId uuid.UUID) (model.User, error)
	FindByEmail(ctx context.Context, email string) (model.User, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (model.User, error)
	FindAll(ctx context.Context) ([]model.User, error)
}