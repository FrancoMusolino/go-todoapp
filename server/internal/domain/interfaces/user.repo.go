package interfaces

import (
	"github.com/FrancoMusolino/go-todoapp/internal/domain/models"
)

type IUserRepo interface {
	CreateUser(u *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	CreateVerificationCode(code *models.VerificationCode) error
	GetLastVerificationCode(userID string) (*models.VerificationCode, error)
	VerifyUser(userID string) error
}
