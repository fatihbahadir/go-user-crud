package service

import (
	"context"
	"user-crud/data/request"
	"user-crud/data/response"

	"github.com/google/uuid"
)

type UserService interface {
	Create(ctx context.Context, request request.UserCreateRequest) error
	Update(ctx context.Context, request request.UserUpdateRequest, userId uuid.UUID) (response.UserResponse, error)
	Delete(ctx context.Context, userId uuid.UUID) error
	FindById(ctx context.Context, userId uuid.UUID) (response.UserResponse, error)
	FindAll(ctx context.Context) ([]response.UserResponse, error)
}