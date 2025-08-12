package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/FrancoMusolino/go-todoapp/utils"
	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBOperationTiemout = time.Second * 10

func NewDatabase() (client *gorm.DB, conn *sql.DB) {
	dsn := utils.GetEnv("DSN")
	log.Printf("=== DATABASE CONNECTION ===")
	log.Printf("Full DSN: %s", dsn)
	log.Printf("========================================")

	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	sqlDB, _ := client.DB()
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Cannot ping DB", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(35)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return client, sqlDB
}
