package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/FrancoMusolino/go-todoapp/internal/api/dtos"
	"github.com/FrancoMusolino/go-todoapp/internal/domain/interfaces"
	"github.com/FrancoMusolino/go-todoapp/internal/domain/models"
	"github.com/FrancoMusolino/go-todoapp/utils"
	"github.com/FrancoMusolino/go-todoapp/utils/logger"
	"github.com/google/uuid"
)

type UserService struct {
	userRepo interfaces.IUserRepo
	logger   *logger.Logger
}

func NewUserService(userRepo interfaces.IUserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
		logger:   logger.NewLogger("User Service"),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req dtos.RegisterUserDto) (*models.User, error) {
	_, err := s.userRepo.GetByEmail(req.Email)
	if err == nil {
		return nil, errors.New("Email already registered")
	}

	_, err = s.userRepo.GetByUsername(req.Username)
	if err == nil {
		return nil, errors.New("Username already in use")
	}

	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		fmt.Println("Error while hashing password", err)
		return nil, errors.New("Cannot register user")
	}

	user := models.User{
		ID:           uuid.New(),
		Username:     req.Username,
		Email:        req.Email,
		Verified:     false,
		PasswordHash: passwordHash,
	}

	err = s.userRepo.CreateUser(&user)
	if err != nil {
		return nil, errors.New("Cannot register user")
	}

	return &user, nil
}
