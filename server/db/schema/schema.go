package schema

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username          string    `gorm:"unique"`
	Email             string    `gorm:"unique"`
	PasswordHash      string
	Birthday          *time.Time
	Verified          bool `gorm:"default:false"`
	VerificationCodes []VerificationCode
	Todos             []Todo
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type VerificationCode struct {
	ID        uint `gorm:"primaryKey"`
	Code      uint
	UserID    string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type Todo struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title       string
	Description string
	Order       uint
	UserID      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
