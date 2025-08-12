package repos

import (
	"context"

	"github.com/FrancoMusolino/go-todoapp/db"
	models "github.com/FrancoMusolino/go-todoapp/db/schema"
	"gorm.io/gorm"
)

type UsersRepo struct {
	client *gorm.DB
}

func NewUsersRepo(client *gorm.DB) *UsersRepo {
	return &UsersRepo{
		client: client,
	}
}

func (r *UsersRepo) CreateUser(user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.DBOperationTiemout)
	defer cancel()

	err := gorm.G[models.User](r.client).Create(ctx, user)

	return user, err
}
