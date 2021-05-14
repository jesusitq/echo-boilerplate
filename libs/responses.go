package libs

import (
	"api-echo-template/models"
	"encoding/json"
	"net/http"
	"strings"

	"gopkg.in/go-validator/validator.v2"
)

// CreateJSONResponse :
func CreateJSONResponse(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(response)
}

// AppendError :
func AppendError(errors []models.ErrorResponse, code, message string) []models.ErrorResponse {
	return append(errors, models.ErrorResponse{Code: code, Message: message})
}

// AppendErrors :
func AppendErrors(errValidator error) (errors []models.ErrorResponse) {
	errs := errValidator.(validator.ErrorMap)
	keys := make([]string, 0, len(errs))
	for k := range errs {
		keys = append(keys, k)
	}

	for i := 0; i < len(errs); i++ {
		errors = AppendError(errors, strings.ToUpper("ERROR."+keys[i]), errs[keys[i]].Error())
	}
	return errors
}
