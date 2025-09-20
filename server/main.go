package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/FrancoMusolino/go-todoapp/db"
	"github.com/FrancoMusolino/go-todoapp/internal/api/handlers"
	"github.com/FrancoMusolino/go-todoapp/internal/repos"
	"github.com/FrancoMusolino/go-todoapp/internal/services"
	"github.com/FrancoMusolino/go-todoapp/mailing"
	"github.com/FrancoMusolino/go-todoapp/router"
	"github.com/FrancoMusolino/go-todoapp/utils"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Cannot load .env file, using env variables")
	}

	// _ := createMailService()

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
	todoRepo := repos.NewTodoRepo(gormDB)

	// Init Services
	userService := services.NewUserService(userRepo)
	todoService := services.NewTodoService(todoRepo)

	// Init Route Handlers
	authHandler := handlers.NewAuthHandler(userService)
	todoHandler := handlers.NewTodoHandler(todoService)

	// Routes
	router := router.SetUpRouter(authHandler, todoHandler)
	port := utils.GetEnvOrDefault("port", "3000")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func createMailService() mailing.IMailService {
	port, _ := strconv.Atoi(utils.GetEnv("MAIL_PORT"))

	return mailing.NewSimpleMailService(
		&mailing.MailConfig{
			Domain:      utils.GetEnv("MAIL_DOMAIN"),
			Host:        utils.GetEnv("MAIL_DOMAIN"),
			Port:        port,
			Encryption:  utils.GetEnv("MAIL_ENCRYPTION"),
			Username:    utils.GetEnv("MAIL_USERNAME"),
			Password:    utils.GetEnv("MAIL_PASSWORD"),
			FromAddress: utils.GetEnv("FROM_ADDRESS"),
			FromName:    utils.GetEnv("FROM_NAME"),
		},
	)

}
