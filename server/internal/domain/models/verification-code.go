package models

import "time"

type VerificationCode struct {
	ID        uint      `json:"id"`
	Code      uint      `json:"code"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
