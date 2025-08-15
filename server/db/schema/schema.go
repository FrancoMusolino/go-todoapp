package schema

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username     string    `gorm:"unique"`
	Email        string    `gorm:"unique"`
	PasswordHash string
	Birthday     *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
