package service

import (
	"context"
	"testing"
	"user-crud/data/request"
	"user-crud/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(ctx context.Context, user model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(ctx context.Context, userId uuid.UUID) error {
	args := m.Called(ctx, userId)
	return args.Error(0)
}

func (m *MockUserRepository) FindById(ctx context.Context, userId uuid.UUID) (model.User, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserRepository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (model.User, error) {
	args := m.Called(ctx, phoneNumber)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserRepository) FindAll(ctx context.Context) ([]model.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, userId uuid.UUID, user model.User) error {
	args := m.Called(ctx, userId, user)
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	userRequest := request.UserCreateRequest{
		Name:        "John",
		Surname:     "Doe",
		Email:       "john.doe@example.com",
		PhoneNumber: "1234567890",
	}

	mockRepo.On("FindByEmail", mock.Anything, userRequest.Email).Return(model.User{}, nil)
	mockRepo.On("FindByPhoneNumber", mock.Anything, userRequest.PhoneNumber).Return(model.User{}, nil)
	mockRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

	err := service.Create(context.Background(), userRequest)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	userId := uuid.New()
	user := model.User{Id: userId}

	mockRepo.On("FindById", mock.Anything, userId).Return(user, nil)
	mockRepo.On("Delete", mock.Anything, userId).Return(nil)

	err := service.Delete(context.Background(), userId)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestFindAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	users := []model.User{
		{Id: uuid.New(), Name: "John", Surname: "Doe", Email: "john.doe@example.com", PhoneNumber: "1234567890"},
		{Id: uuid.New(), Name: "Jane", Surname: "Doe", Email: "jane.doe@example.com", PhoneNumber: "0987654321"},
	}


	mockRepo.On("FindAll", mock.Anything).Return(users, nil)


	result, err := service.FindAll(context.Background())


	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestFindByIdUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	userId := uuid.New()
	user := model.User{Id: userId, Name: "John", Surname: "Doe", Email: "john.doe@example.com", PhoneNumber: "1234567890"}


	mockRepo.On("FindById", mock.Anything, userId).Return(user, nil)


	result, err := service.FindById(context.Background(), userId)


	assert.NoError(t, err)
	assert.Equal(t, userId, result.Id)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	userId := uuid.New()
	userRequest := request.UserUpdateRequest{
		Id:          userId,
		Name:        "Updated Name",
		Surname:     "Updated Surname",
		Email:       "updated.email@example.com",
		PhoneNumber: "1231231234",
	}


	user := model.User{Id: userId, Name: "John", Surname: "Doe", Email: "john.doe@example.com", PhoneNumber: "1234567890"}
	mockRepo.On("FindById", mock.Anything, userId).Return(user, nil)
	mockRepo.On("Update", mock.Anything, userId, mock.Anything).Return(nil)


	result, err := service.Update(context.Background(), userRequest, userId)


	assert.NoError(t, err)
	assert.Equal(t, userRequest.Name, result.Name)
	assert.Equal(t, userRequest.Email, result.Email)
	mockRepo.AssertExpectations(t)
}