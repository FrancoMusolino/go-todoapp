package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FrancoMusolino/go-todoapp/internal/api/dtos"
	"github.com/FrancoMusolino/go-todoapp/internal/services"
	"github.com/FrancoMusolino/go-todoapp/utils"
)

type AuthHandler struct {
	userService *services.UserService
}

func NewAuthHandler(userService *services.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dtos.RegisterUserDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("JSON decode error: %v", err)
		utils.WriteError(w, http.StatusBadRequest, "Invalid JSON format", nil)
		return
	}

	h.userService.CreateUser(r.Context(), req)
	utils.WriteJson(w, http.StatusOK, &utils.ApiResponse[interface{}]{})
}
