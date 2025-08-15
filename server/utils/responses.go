package utils

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field"`
}

type PaginationMetadata struct {
	PageSize     int `json:"pageSize"`
	PageNumber   int `json:"pageNumber"`
	TotalPages   int `json:"totalPages"`
	TotalRecords int `json:"totalRecords"`
}

type ApiResponse[T any] struct {
	Success    bool                `json:"success"`
	Message    string              `json:"message"`
	Data       *T                  `json:"data,omitempty"`
	Pagination *PaginationMetadata `json:"pagination,omitempty"`
	Errors     []ApiError          `json:"errors,omitempty"`
}

func WriteJson[T any](w http.ResponseWriter, status int, v *ApiResponse[T]) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, msg string, errors []ApiError) {
	res := &ApiResponse[interface{}]{
		Success:    false,
		Message:    msg,
		Data:       nil,
		Pagination: nil,
		Errors:     errors,
	}

	WriteJson(w, status, res)
}
