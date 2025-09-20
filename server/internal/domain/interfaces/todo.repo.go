package interfaces

import (
	"github.com/FrancoMusolino/go-todoapp/internal/domain/models"
	"github.com/FrancoMusolino/go-todoapp/utils/pagination"
)

type ITodoRepo interface {
	GetUserTodos(params GetUserTodoParams) ([]*models.Todo, error)
	GetTodoLastOrder(userID string) uint
	CreateTodo(u *models.Todo) (*models.Todo, error)
}

type GetUserTodoParams struct {
	UserID string
	*pagination.PaginationParams
}
