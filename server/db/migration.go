package db

import (
	"github.com/FrancoMusolino/go-todoapp/db/schema"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	return db.AutoMigrate(&schema.User{})
}
