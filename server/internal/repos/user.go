package repos

import (
	"context"

	"github.com/FrancoMusolino/go-todoapp/db"
	"github.com/FrancoMusolino/go-todoapp/db/schema"
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

func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.DBOperationTiemout)
	defer cancel()

	existing, err := gorm.G[schema.User](r.client).Where("email = ?", email).First(ctx)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:           existing.ID,
		Username:     existing.Email,
		Email:        existing.Email,
		PasswordHash: existing.PasswordHash,
		Birthday:     existing.Birthday,
		CreatedAt:    existing.CreatedAt,
		UpdatedAt:    existing.UpdatedAt,
	}

	return user, nil
}

func (r *UserRepo) GetByUsername(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.DBOperationTiemout)
	defer cancel()

	existing, err := gorm.G[schema.User](r.client).Where("username = ?", username).First(ctx)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:           existing.ID,
		Username:     existing.Email,
		Email:        existing.Email,
		PasswordHash: existing.PasswordHash,
		Birthday:     existing.Birthday,
		CreatedAt:    existing.CreatedAt,
		UpdatedAt:    existing.UpdatedAt,
	}

	return user, nil
}
