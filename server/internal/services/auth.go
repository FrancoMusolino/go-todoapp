package services

import (
	"context"
	"errors"

	"github.com/FrancoMusolino/go-todoapp/internal/api/dtos"
	"github.com/FrancoMusolino/go-todoapp/internal/domain/models"
	"github.com/FrancoMusolino/go-todoapp/utils"
	"github.com/FrancoMusolino/go-todoapp/utils/logger"
)

type AuthService struct {
	userService *UserService
	logger      *logger.Logger
}

func NewAuthService(userService *UserService) *AuthService {
	return &AuthService{
		userService: userService,
		logger:      logger.NewLogger("User Service"),
	}
}

func (as *AuthService) Register(ctx context.Context, req dtos.RegisterUserDto) (*models.User, error) {
	if !utils.PasswordMatchRegex(req.Password) {
		return nil, errors.New("Password must be at least 6 characters long and include at least one uppercase letter, one lowercase letter, and one number.")
	}

	_, err := as.userService.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return nil, nil

}
