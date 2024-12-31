package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-crud/data/request"
	"user-crud/data/response"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock implementation of the UserService interface
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Create(ctx context.Context, req request.UserCreateRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockUserService) Update(ctx context.Context, req request.UserUpdateRequest, userId uuid.UUID) (response.UserResponse, error) {
	args := m.Called(ctx, req, userId)
	return args.Get(0).(response.UserResponse), args.Error(1)
}

func (m *MockUserService) Delete(ctx context.Context, userId uuid.UUID) error {
	args := m.Called(ctx, userId)
	return args.Error(0)
}

func (m *MockUserService) FindAll(ctx context.Context) ([]response.UserResponse, error) {
	args := m.Called(ctx)
	return args.Get(0).([]response.UserResponse), args.Error(1)
}

func (m *MockUserService) FindById(ctx context.Context, userId uuid.UUID) (response.UserResponse, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).(response.UserResponse), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockService := new(MockUserService)
	controller := NewUserController(mockService)

	reqBody := request.UserCreateRequest{
		Name:        "John",
		Surname:     "Doe",
		Email:       "john.doe@example.com",
		PhoneNumber: "123456789",
	}
	body, _ := json.Marshal(reqBody)

	mockService.On("Create", mock.Anything, reqBody).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	controller.Create(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	mockService := new(MockUserService)
	controller := NewUserController(mockService)

	userId := uuid.New()
	reqBody := request.UserUpdateRequest{
		Id:          userId,
		Name:        "John Updated",
		Surname:     "Doe Updated",
		Email:       "john.updated@example.com",
		PhoneNumber: "987654321",
	}
	body, _ := json.Marshal(reqBody)

	updatedUser := response.UserResponse{
		Id:          userId,
		Name:        "John Updated",
		Surname:     "Doe Updated",
		Email:       "john.updated@example.com",
		PhoneNumber: "987654321",
	}

	mockService.On("Update", mock.Anything, reqBody, userId).Return(updatedUser, nil)

	req := httptest.NewRequest(http.MethodPut, "/users/"+userId.String(), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	vars := map[string]string{"userId": userId.String()}
	req = mux.SetURLVars(req, vars)

	controller.Update(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockService := new(MockUserService)
	controller := NewUserController(mockService)

	userId := uuid.New()

	mockService.On("Delete", mock.Anything, userId).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/users/"+userId.String(), nil)
	rec := httptest.NewRecorder()

	vars := map[string]string{"userId": userId.String()}
	req = mux.SetURLVars(req, vars)

	controller.Delete(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestFindAllUsers(t *testing.T) {
	mockService := new(MockUserService)
	controller := NewUserController(mockService)

	users := []response.UserResponse{
		{
			Id:          uuid.New(),
			Name:        "John",
			Surname:     "Doe",
			Email:       "john.doe@example.com",
			PhoneNumber: "123456789",
		},
		{
			Id:          uuid.New(),
			Name:        "Jane",
			Surname:     "Smith",
			Email:       "jane.smith@example.com",
			PhoneNumber: "987654321",
		},
	}

	mockService.On("FindAll", mock.Anything).Return(users, nil)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()

	controller.FindAll(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestFindUserById(t *testing.T) {
	mockService := new(MockUserService)
	controller := NewUserController(mockService)

	userId := uuid.New()
	user := response.UserResponse{
		Id:          userId,
		Name:        "John",
		Surname:     "Doe",
		Email:       "john.doe@example.com",
		PhoneNumber: "123456789",
	}

	mockService.On("FindById", mock.Anything, userId).Return(user, nil)

	req := httptest.NewRequest(http.MethodGet, "/users/"+userId.String(), nil)
	rec := httptest.NewRecorder()

	vars := map[string]string{"userId": userId.String()}
	req = mux.SetURLVars(req, vars)

	controller.FindById(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}
