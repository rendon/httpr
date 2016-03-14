package httpr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// DataResponse describes a general JSON response.
type DataResponse struct {
	Status int         `json:"status"`
	Errors []string    `json:"errors,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

// Error replies with error message and code.
func Error(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	fmt.Fprintf(w, `{"code": %d, "message": "%s"}`, code, message)
}

// Created replies with HTTP CREATED code (201).
func Created(w http.ResponseWriter, message string) {
	Error(w, message, http.StatusCreated)
}

// Data replies with code and data as a JSON document.
func Data(w http.ResponseWriter, data interface{}, code int) {
	resp := DataResponse{
		Status: code,
		Data:   data,
	}
	body, err := json.Marshal(resp)
	if err != nil {
		Error(w, "Failed to marshal response", http.StatusInternalServerError)
	} else {
		w.WriteHeader(code)
		fmt.Fprintf(w, "%s", body)
	}
}
