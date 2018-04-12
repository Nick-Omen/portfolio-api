package server

import (
	"net/http"
	"strings"
	"encoding/json"
)

type ValidationError struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type ValidationErrorResponse struct {
	Errors []ValidationError `json:"errors"`
}

func ResponseValidationError(errors string, w http.ResponseWriter) {
	errorList := strings.Split(errors, ";")
	var respErr []ValidationError
	for _, err := range errorList {
		errKeyVal := strings.Split(err, "|")
		key, val := errKeyVal[0], errKeyVal[1]
		respErr = append(respErr, ValidationError{Key: key, Value: val})
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ValidationErrorResponse{Errors: respErr})
}
