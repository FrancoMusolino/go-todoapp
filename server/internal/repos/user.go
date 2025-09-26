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

func (r *UserRepo) CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.DBOperationTiemout)
	defer cancel()

	err := gorm.G[models.User](r.client).Create(ctx, user)
	return err
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
		Username:     existing.Username,
		Email:        existing.Email,
		PasswordHash: existing.PasswordHash,
		Birthday:     existing.Birthday,
		Verified:     existing.Verified,
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
		Username:     existing.Username,
		Email:        existing.Email,
		PasswordHash: existing.PasswordHash,
		Birthday:     existing.Birthday,
		Verified:     existing.Verified,
		CreatedAt:    existing.CreatedAt,
		UpdatedAt:    existing.UpdatedAt,
	}

	return user, nil
}

func (r *UserRepo) CreateVerificationCode(code *models.VerificationCode) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.DBOperationTiemout)
	defer cancel()

	err := gorm.G[models.VerificationCode](r.client).Create(ctx, code)
	return err
}

func (r *UserRepo) GetLastVerificationCode(userID string) (*models.VerificationCode, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.DBOperationTiemout)
	defer cancel()

	existing, err := gorm.G[schema.VerificationCode](r.client).Where("user_id = ?", userID).Order("created_at desc").Take(ctx)
	if err != nil {
		return nil, err
	}

	code := &models.VerificationCode{
		ID:        existing.ID,
		UserID:    existing.UserID,
		Code:      existing.Code,
		CreatedAt: existing.CreatedAt,
		ExpiresAt: existing.ExpiresAt,
	}

	return code, nil
}

func (r *UserRepo) VerifyUser(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.DBOperationTiemout)
	defer cancel()

	_, err := gorm.G[schema.VerificationCode](r.client).Where("user_id = ?", userID).Update(ctx, "verified", true)
	return err
}
