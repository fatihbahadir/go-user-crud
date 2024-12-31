package helper

import "net/http"

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(code int, message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func WriteSuccessResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	response := NewSuccessResponse(code, message, data)
	WriteJSONResponse(w, code, response)
}