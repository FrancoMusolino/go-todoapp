package models

import (
	"crypto/subtle"
	"time"

	"github.com/FrancoMusolino/go-todoapp/utils"
	"github.com/google/uuid"
)

type VerificationCode struct {
	ID        uint      `json:"id"`
	Code      uint      `json:"code"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (v *VerificationCode) IsExpired() bool {
	return v.ExpiresAt.Before(time.Now())
}

func (v *VerificationCode) SafeCompare(y []byte) bool {
	x := utils.UintToTextBytes(v.Code)

	if len(x) != len(y) {
		return false
	}

	if subtle.ConstantTimeCompare(x, y) == 0 {
		return false
	}

	return true
}
