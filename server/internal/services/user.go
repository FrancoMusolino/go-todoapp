package services

import (
	"context"
	"errors"

	"github.com/FrancoMusolino/go-todoapp/internal/api/dtos"
	"github.com/FrancoMusolino/go-todoapp/internal/domain/interfaces"
	"github.com/FrancoMusolino/go-todoapp/internal/domain/models"
	"github.com/FrancoMusolino/go-todoapp/utils"
	"github.com/google/uuid"
)

type UserService struct {
	userRepo interfaces.IUsersRepo
}

func NewUserService(userRepo interfaces.IUsersRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req dtos.RegisterUserDto) (*models.User, error) {

	if !utils.PasswordMatchRegex(req.Password) {
		return nil, errors.New("Password must be at least 6 characters long and include at least one uppercase letter, one lowercase letter, and one number.")
	}

	_, err := s.userRepo.GetByEmail(req.Email)
	if err == nil {
		return nil, errors.New("Email already registered")
	}

	_, err = s.userRepo.GetByUsername(req.Username)
	if err == nil {
		return nil, errors.New("Username already in use")
	}

	user := models.User{
		ID:           uuid.New(),
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: req.Password,
	}

	_, err = s.userRepo.CreateUser(&user)
	if err != nil {
		return nil, errors.New("Cannot register user")
	}

	return &user, nil
}
