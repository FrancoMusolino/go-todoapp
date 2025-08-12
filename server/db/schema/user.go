package db

import "time"

type User struct {
	ID        uint
	Firstname string
	Lastname  string
	Email     string
	Birthday  *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
