package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	Validator = NewValidator()
}

func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func WriteJsonResponse(w http.ResponseWriter, statusCode int, data any) error {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json")
	}

	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func WriteJsonSuccessResponse(w http.ResponseWriter, statusCode int, message string, data any) error {
	response := map[string]any{}

	response["message"] = message
	response["data"] = data
	response["status"] = "success"
	return WriteJsonResponse(w, statusCode, response)
}

func WriteJsonErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) error {
	response := map[string]any{}

	response["message"] = message
	response["status"] = "error"
	response["error"] = err.Error()
	return WriteJsonResponse(w, statusCode, response)
}
func ReadJsonRequest(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}
