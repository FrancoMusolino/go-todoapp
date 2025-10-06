package interfaces

import (
	"github.com/FrancoMusolino/go-todoapp/internal/domain/models"
	"github.com/FrancoMusolino/go-todoapp/utils/pagination"
	"github.com/google/uuid"
)

type ITodoRepo interface {
	GetUserTodos(params GetUserTodoParams) ([]*models.Todo, error)
	GetTodoLastOrder(userID uuid.UUID) uint
	CreateTodo(u *models.Todo) (*models.Todo, error)
}

type GetUserTodoParams struct {
	UserID string
	*pagination.PaginationParams
}
