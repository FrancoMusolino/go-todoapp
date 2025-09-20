package services

import (
	"context"
	"errors"

	"github.com/FrancoMusolino/go-todoapp/internal/api/dtos"
	"github.com/FrancoMusolino/go-todoapp/internal/domain/interfaces"
	"github.com/FrancoMusolino/go-todoapp/internal/domain/models"
	"github.com/FrancoMusolino/go-todoapp/utils/logger"
	"github.com/FrancoMusolino/go-todoapp/utils/pagination"
	"github.com/google/uuid"
)

type TodoService struct {
	todoRepo interfaces.ITodoRepo
	logger   logger.Logger
}

func NewTodoService(todoRepo interfaces.ITodoRepo) *TodoService {
	return &TodoService{
		todoRepo: todoRepo,
		logger:   *logger.NewLogger("Todo Service"),
	}
}

func (s *TodoService) GetUserTodos(ctx context.Context) ([]*models.Todo, error) {
	userID := ctx.Value("userID").(string)

	todos, err := s.todoRepo.GetUserTodos(interfaces.GetUserTodoParams{UserID: userID, PaginationParams: pagination.NewPaginationParams(1, 1)})
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *TodoService) CreateTodo(ctx context.Context, req dtos.CreateTodoDto) (*models.Todo, error) {
	userID := ctx.Value("userID").(string)

	lastTodoOrder := s.todoRepo.GetTodoLastOrder(userID)
	todo := models.Todo{
		ID:          uuid.New(),
		Title:       req.Title,
		Description: req.Description,
		UserID:      userID,
		Order:       lastTodoOrder + 1,
	}

	_, err := s.todoRepo.CreateTodo(&todo)
	if err != nil {
		s.logger.Error(ctx, "CreateTodo", err.Error())
		return nil, errors.New("Cannot create todo")
	}

	return &todo, nil
}
