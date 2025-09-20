package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/FrancoMusolino/go-todoapp/internal/api/dtos"
	"github.com/FrancoMusolino/go-todoapp/internal/domain/interfaces"
	"github.com/FrancoMusolino/go-todoapp/internal/domain/models"
	"github.com/FrancoMusolino/go-todoapp/utils"
	"github.com/FrancoMusolino/go-todoapp/utils/logger"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
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

	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		fmt.Println("Error while hashing password", err)
		return nil, errors.New("Cannot register user")
	}

	user := models.User{
		ID:           uuid.New(),
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: passwordHash,
	}

	_, err = s.userRepo.CreateUser(&user)
	if err != nil {
		return nil, errors.New("Cannot register user")
	}

	return &user, nil
}

func (s *UserService) GetToken(ctx context.Context, req dtos.LoginDto) (*dtos.LoginResponseDto, error) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("Invalid email or password")
	}

	err = utils.ComparePasswords(user.PasswordHash, req.Password)
	if err != nil {
		log.Printf("Invalid pass %s", err)
		return nil, errors.New("Invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		ID:       user.ID.String(),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	secret := utils.GetEnv("JWT_SECRET")
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Printf("UserService")
		return nil, errors.New("Cannot authenticate the user")
	}

	return &dtos.LoginResponseDto{
		ID:       user.ID.String(),
		Username: user.Username,
		Token:    signed,
	}, nil
}
