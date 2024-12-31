package helper

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ReadRequestBody(r *http.Request, result interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() 
	err := decoder.Decode(result)
	if err != nil {
		return errors.New("invalid JSON format or unexpected fields")
	}
	return nil
}


func WriteSuceResponse(writer http.ResponseWriter, code int, status string, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)

	json.NewEncoder(writer).Encode(WebResponse{
		Code:   code,
		Status: status,
		Data:   data,
	})
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

