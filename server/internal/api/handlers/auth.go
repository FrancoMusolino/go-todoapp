package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FrancoMusolino/go-todoapp/internal/api/dtos"
	"github.com/FrancoMusolino/go-todoapp/internal/services"
	"github.com/FrancoMusolino/go-todoapp/utils"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

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

	if err := validate.Struct(req); err != nil {
		res := utils.ApiResponse[any]{
			Success: false,
			Message: "Invalid body request",
			Errors:  utils.MapValidationErrors(err.(validator.ValidationErrors)),
		}

		utils.WriteJson(w, http.StatusBadRequest, &res)
		return
	}

	if !utils.PasswordMatchRegex(req.Password) {
		res := utils.ApiResponse[any]{
			Success: false,
			Message: "Password must be at least 6 characters long and include at least one uppercase letter, one lowercase letter, and one number.",
		}

		utils.WriteJson(w, http.StatusBadRequest, &res)
		return
	}

	h.userService.CreateUser(r.Context(), req)
	utils.WriteJson(w, http.StatusOK, &utils.ApiResponse[any]{})
}
