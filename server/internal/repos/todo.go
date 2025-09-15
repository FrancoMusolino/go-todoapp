package repos

import (
	"context"

	"github.com/FrancoMusolino/go-todoapp/db"
	"github.com/FrancoMusolino/go-todoapp/db/schema"
	"github.com/FrancoMusolino/go-todoapp/internal/domain/models"
	"gorm.io/gorm"
)

type TodoRepo struct {
	client *gorm.DB
}

func NewTodoRepo(client *gorm.DB) *TodoRepo {
	return &TodoRepo{
		client: client,
	}
}

func (r *TodoRepo) GetTodoLastOrder(userId string) uint {
	ctx, cancel := context.WithTimeout(context.Background(), db.DBOperationTiemout)
	defer cancel()

	todo, err := gorm.G[schema.Todo](r.client).Where("user_id = ?", userId).Order("\"order\"").Take(ctx)
	if err != nil {
		return 0
	}

	return todo.Order
}

func (r *TodoRepo) CreateTodo(todo *models.Todo) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.DBOperationTiemout)
	defer cancel()

	err := gorm.G[models.Todo](r.client).Create(ctx, todo)

	return todo, err
}
