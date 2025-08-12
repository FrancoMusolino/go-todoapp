package db

import (
	"github.com/FrancoMusolino/go-todoapp/db/schema"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&schema.User{})
}
