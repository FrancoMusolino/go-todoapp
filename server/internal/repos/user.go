package repos

import (
	"context"

	"github.com/FrancoMusolino/go-todoapp/db"
	"github.com/FrancoMusolino/go-todoapp/internal/domain/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	client *gorm.DB
}

func NewUserRepo(client *gorm.DB) *UserRepo {
	return &UserRepo{
		client: client,
	}
}

func (r *UserRepo) CreateUser(user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.DBOperationTiemout)
	defer cancel()

	err := gorm.G[models.User](r.client).Create(ctx, user)

	return user, err
}
