package controller

import (
	"net/http"
	"user-crud/data/request"
	"user-crud/helper"
	"user-crud/service"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (controller *UserController) Create(writer http.ResponseWriter, requests *http.Request) {
	userCreateRequest := request.UserCreateRequest{}
	err := helper.ReadRequestBody(requests, &userCreateRequest)

	if err != nil {
		response := helper.NewErrorResponse(400, "Invalid Request Body", nil)
		helper.WriteJSONResponse(writer, http.StatusBadRequest, response)
		return
	}

	if err := controller.UserService.Create(requests.Context(), userCreateRequest); err != nil {
		if errorResponse, ok := err.(*helper.ErrorResponse); ok {
			helper.WriteJSONResponse(writer, errorResponse.Code, errorResponse)
			return
		}

		response := helper.NewErrorResponse(http.StatusInternalServerError, "Internal server error", nil)
		helper.WriteJSONResponse(writer, http.StatusInternalServerError, response)
		return
	}

	helper.WriteSuccessResponse(writer, http.StatusCreated, "User created successfully", nil)
}

func (controller *UserController) Update(writer http.ResponseWriter, requests *http.Request) {
	userUpdateRequest := request.UserUpdateRequest{}
	err := helper.ReadRequestBody(requests, &userUpdateRequest)

	if err != nil {
		response := helper.NewErrorResponse(400, "Invalid Request Body", nil)
		helper.WriteJSONResponse(writer, http.StatusBadRequest, response)
		return
	}

	userId := mux.Vars(requests)["userId"]
	id, err := uuid.Parse(userId)

	if err != nil {
		response := helper.NewErrorResponse(http.StatusBadRequest, "Invalid user ID", nil)
		helper.WriteJSONResponse(writer, http.StatusBadRequest, response)
		return
	}

	userUpdateRequest.Id = id

	updatedUser, err := controller.UserService.Update(requests.Context(), userUpdateRequest, id)

	if err != nil {
		if errorResponse, ok := err.(*helper.ErrorResponse); ok {
			helper.WriteJSONResponse(writer, errorResponse.Code, errorResponse)
			return
		}
		response := helper.NewErrorResponse(http.StatusInternalServerError, "Internal server error", nil)
		helper.WriteJSONResponse(writer, http.StatusInternalServerError, response)
		return
	}

	helper.WriteSuccessResponse(writer, http.StatusOK, "User updated successfully", updatedUser)
}

func (controller *UserController) Delete(writer http.ResponseWriter, requests *http.Request) {
	userId := mux.Vars(requests)["userId"]
	id, err := uuid.Parse(userId)

	if err != nil {
		response := helper.NewErrorResponse(http.StatusBadRequest, "Invalid user ID", nil)
		helper.WriteJSONResponse(writer, http.StatusBadRequest, response)
		return
	}

	err = controller.UserService.Delete(requests.Context(), id)
	if err != nil {
		if errorResponse, ok := err.(*helper.ErrorResponse); ok {
			helper.WriteJSONResponse(writer, errorResponse.Code, errorResponse)
			return
		}
		helper.WriteJSONResponse(writer, http.StatusInternalServerError, helper.NewErrorResponse(500, "Internal server error", nil))
		return
	}

	helper.WriteSuccessResponse(w, http.StatusOK, "User deleted successfully", nil)
}

func (controller *UserController) FindAll(writer http.ResponseWriter, requests *http.Request) {
	users, err := controller.UserService.FindAll(requests.Context())
	if err != nil {
		if errorResponse, ok := err.(*helper.ErrorResponse); ok {
			helper.WriteJSONResponse(writer, errorResponse.Code, errorResponse)
			return
		}
		helper.WriteJSONResponse(writer, http.StatusInternalServerError, helper.NewErrorResponse(500, "Internal server error", nil))
		return
	}
	helper.WriteJSONResponse(writer, http.StatusOK, users)

}

func (controller *UserController) FindById(writer http.ResponseWriter, requests *http.Request) {
	userId := mux.Vars(requests)["userId"]
	id, err := uuid.Parse(userId)

	if err != nil {
		response := helper.NewErrorResponse(http.StatusBadRequest, "Invalid user ID", nil)
		helper.WriteJSONResponse(writer, http.StatusBadRequest, response)
		return
	}

	userResponse, err := controller.UserService.FindById(r.Context(), id)
	if err != nil {
		if errorResponse, ok := err.(*helper.ErrorResponse); ok {
			helper.WriteJSONResponse(writer, errorResponse.Code, errorResponse)
			return
		}
		response := helper.NewErrorResponse(http.StatusInternalServerError, "Internal server error", nil)
		helper.WriteJSONResponse(writer, http.StatusInternalServerError, response)
		return
	}

	helper.WriteJSONResponse(writer, http.StatusOK, userResponse)
}
