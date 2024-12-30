package service

import (
	"context"
	"fmt"
	"user-crud/data/request"
	"user-crud/data/response"
	"user-crud/helper"
	"user-crud/model"
	"user-crud/repository"

	"github.com/google/uuid"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImplementation(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (service *UserServiceImpl) Create(ctx context.Context, request request.UserCreateRequest) error {
	err := helper.ValidateStruct(request)
	if err != nil {
		return err
	}

	existingUser, err := service.UserRepository.FindByEmail(ctx, request.Email)

	if err != nil {
		return err
	}

	if existingUser.Email != "" {
		return helper.NewErrorResponse(400, "User with email already exists", nil)
	}

	existingUserByPhoneNumber, err := service.UserRepository.FindByPhoneNumber(ctx, request.PhoneNumber)
	if err != nil {
		return err
	}

	if existingUserByPhoneNumber.PhoneNumber != "" {
		return helper.NewErrorResponse(400, "User with this phone nubmer already exists", nil)
	}

	user := model.User{
		Name:        request.Name,
		Surname:     request.Surname,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	}

	if err := service.UserRepository.Save(ctx, user); err != nil {
		return helper.NewErrorResponse(500, "Failed to save user", nil)
	}

	return nil
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId uuid.UUID) error {
	user, err := service.UserRepository.FindById(ctx, userId)
	if err != nil {
		return helper.NewErrorResponse(404, fmt.Sprintf("User with id %s not found", userId), nil)
	}

	err = service.UserRepository.Delete(ctx, user.Id)
	if err != nil {
		return helper.NewErrorResponse(500, "Failed to delete user", nil)
	}

	return nil
}

func (service *UserServiceImpl) FindAll(ctx context.Context) ([]response.UserResponse, error) {
	users, err := service.UserRepository.FindAll(ctx)
	if err != nil {
		return nil, helper.NewErrorResponse(500, "Failed to retrieve users", nil)
	}

	if len(users) == 0 {
		return nil, helper.NewErrorResponse(404, "No users found", nil)
	}

	var userResponses []response.UserResponse
	for _, user := range users {
		userResponse := response.UserResponse{
			Id:          user.Id,
			Name:        user.Name,
			Surname:     user.Surname,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			CreatedAt:   user.CreatedAt,
		}
		userResponses = append(userResponses, userResponse)
	}

	return userResponses, nil
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId uuid.UUID) (response.UserResponse, error) {
	user, err := service.UserRepository.FindById(ctx, userId)
	if err != nil {
		return response.UserResponse{}, helper.NewErrorResponse(404, "User not found", nil)
	}

	userResponse := response.UserResponse{
		Id:          user.Id,
		Name:        user.Name,
		Surname:     user.Surname,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
	}

	return userResponse, nil
}

func (service *UserServiceImpl) Update(ctx context.Context, request request.UserUpdateRequest, userId uuid.UUID) (response.UserResponse, error) {
	err := helper.ValidateStruct(request)
	if err != nil {
		return response.UserResponse{}, err
	}

	user, err := service.UserRepository.FindById(ctx, request.Id)
	if err != nil {
		return response.UserResponse{}, helper.NewErrorResponse(404, "User with given id not found", nil)
	}

	if request.Name == "" && request.Surname == "" && request.Email == "" && request.PhoneNumber == "" {
		return response.UserResponse{}, helper.NewErrorResponse(400, "No fields to update", nil) 
	}

	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Surname != "" {
		user.Surname = request.Surname
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.PhoneNumber != "" {
		user.PhoneNumber = request.PhoneNumber
	}

	err = service.UserRepository.Update(ctx, request.Id, user)

	if err != nil {
		return response.UserResponse{}, helper.NewErrorResponse(500, "Failed to update user", nil)
	}

	userResponse := response.UserResponse{
		Id:          user.Id,
		Name:        user.Name,
		Surname:     user.Surname,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
	}

	return userResponse, nil
}