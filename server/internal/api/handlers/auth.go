package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FrancoMusolino/go-todoapp/internal/api/dtos"
	"github.com/FrancoMusolino/go-todoapp/internal/logger"
	"github.com/FrancoMusolino/go-todoapp/internal/services"
	"github.com/FrancoMusolino/go-todoapp/utils"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type AuthHandler struct {
	userService *services.UserService
	logger      *logger.Logger
}

func NewAuthHandler(userService *services.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
		logger:      logger.NewLogger("Auth Handler"),
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

	_, err := h.userService.CreateUser(r.Context(), req)
	if err != nil {
		res := utils.ApiResponse[any]{
			Success: false,
			Message: err.Error(),
		}

		utils.WriteJson(w, http.StatusBadRequest, &res)
		return
	}

	utils.WriteJson(w, http.StatusOK, &utils.ApiResponse[any]{
		Success: true,
		Message: "User registered",
	})
	return
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dtos.LoginDto
	h.logger.IncomingRequest(r, r.Context())

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

	data, err := h.userService.GetToken(r.Context(), req)
	if err != nil {
		res := utils.ApiResponse[any]{
			Success: false,
			Message: err.Error(),
		}

		utils.WriteJson(w, http.StatusBadRequest, &res)
		return
	}

	utils.WriteJson(w, http.StatusOK, &utils.ApiResponse[dtos.LoginResponseDto]{
		Success: true,
		Message: "Login success",
		Data:    data,
	})
	return
}
