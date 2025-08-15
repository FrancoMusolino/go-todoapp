package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func Decode[T any](w http.ResponseWriter, r *http.Request) T {
	var req T

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("JSON decode error: %v", err)
		WriteError(w, http.StatusBadRequest, "Invalid JSON format", nil)
	}

	return req
}
