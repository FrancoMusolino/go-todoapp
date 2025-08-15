package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FrancoMusolino/go-todoapp/db"
	"github.com/FrancoMusolino/go-todoapp/internal/api/handlers"
	"github.com/FrancoMusolino/go-todoapp/internal/repos"
	"github.com/FrancoMusolino/go-todoapp/internal/services"
	"github.com/FrancoMusolino/go-todoapp/router"
	"github.com/FrancoMusolino/go-todoapp/utils"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Cannot load .env file, using env variables")
	}

	// DB
	gormDB, dbConnection := db.NewDatabase()
	defer dbConnection.Close()

	// Migrations (Auto)
	err := db.RunMigrations(gormDB)
	if err != nil {
		fmt.Println("Cannot run migrations", err)
	}

	// Init Repos
	userRepo := repos.NewUserRepo(gormDB)

	// Init Services
	userService := services.NewUserService(userRepo)

	// Init Route Handlers
	authHandler := handlers.NewAuthHandler(userService)

	// Routes
	router := router.SetUpRouter(authHandler)
	port := utils.GetEnvOrDefault("port", "3000")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
