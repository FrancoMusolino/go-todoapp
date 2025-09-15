package interfaces

import (
	"github.com/FrancoMusolino/go-todoapp/internal/domain/models"
)

type ITodoRepo interface {
	GetTodoLastOrder(userId string) uint
	CreateTodo(u *models.Todo) (*models.Todo, error)
}
