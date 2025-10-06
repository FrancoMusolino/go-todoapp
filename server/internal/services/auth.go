package services

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/FrancoMusolino/go-todoapp/internal/api/dtos"
	"github.com/FrancoMusolino/go-todoapp/internal/domain/interfaces"
	"github.com/FrancoMusolino/go-todoapp/internal/domain/models"
	"github.com/FrancoMusolino/go-todoapp/utils"
	"github.com/FrancoMusolino/go-todoapp/utils/logger"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type AuthService struct {
	userService *UserService
	userRepo    interfaces.IUserRepo
	logger      *logger.Logger
}

func NewAuthService(userService *UserService, userRepo interfaces.IUserRepo) *AuthService {
	return &AuthService{
		userService: userService,
		userRepo:    userRepo,
		logger:      logger.NewLogger("User Service"),
	}
}

func (as *AuthService) Register(ctx context.Context, req dtos.RegisterUserDto) (*models.User, error) {
	if !utils.PasswordMatchRegex(req.Password) {
		return nil, errors.New("Password must be at least 6 characters long and include at least one uppercase letter, one lowercase letter, and one number.")
	}

	user, err := as.userService.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	randCode := rand.Intn(900_000) + 100_000
	expiresAt := time.Now().Add(10 * time.Minute)
	code := models.VerificationCode{
		Code:      uint(randCode),
		UserID:    user.ID,
		ExpiresAt: expiresAt,
	}

	as.userRepo.CreateVerificationCode(&code)

	return user, nil
}

func (s *AuthService) Login(ctx context.Context, req dtos.LoginDto) (*dtos.LoginResponseDto, error) {
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
