package models

import (
	"time"

	"github.com/google/uuid"
)

type VerificationCode struct {
	ID        uint      `json:"id"`
	Code      uint      `json:"code"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
