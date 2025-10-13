package models

import (
	"crypto/subtle"
	"math/rand"
	"time"

	"github.com/FrancoMusolino/go-todoapp/utils"
	"github.com/google/uuid"
)

var expirationTime = 10 * time.Minute

type VerificationCode struct {
	ID        uint      `json:"id"`
	Code      uint      `json:"code"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewVerificationCode(userID uuid.UUID) *VerificationCode {
	randCode := rand.Intn(900_000) + 100_000
	expiresAt := time.Now().Add(expirationTime)

	return &VerificationCode{
		Code:      uint(randCode),
		UserID:    userID,
		ExpiresAt: expiresAt,
	}
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
