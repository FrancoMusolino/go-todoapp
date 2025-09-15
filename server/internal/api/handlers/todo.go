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

type TodoHandler struct {
	todoService *services.TodoService
	logger      *logger.Logger
}

func NewTodoHandler(todoService *services.TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
		logger:      logger.NewLogger("Todo Handler"),
	}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req dtos.CreateTodoDto
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

	_, err := h.todoService.CreateTodo(r.Context(), req)
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
		Message: "Todo Created",
	})
	return
}
