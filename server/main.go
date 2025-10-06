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

	// Mailing
	mailingJobQueue := make(chan mailing.Message)
	defer close(mailingJobQueue)
	go startMailing(mailingJobQueue)

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
	authService := services.NewAuthService(userService, userRepo)

	// Init Route Handlers
	authHandler := handlers.NewAuthHandler(authService)
	todoHandler := handlers.NewTodoHandler(todoService)

	// Routes
	router := router.SetUpRouter(authHandler, todoHandler, userRepo)
	port := utils.GetEnvOrDefault("port", "3000")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func startMailing(mailingJobQueue chan mailing.Message) {
	numWorkers := 3
	mailingResultsChan := make(chan error)

	port, _ := strconv.Atoi(utils.GetEnv("MAIL_PORT"))
	mailingService := mailing.NewSimpleMailService(
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

	mailingDispatcher := mailing.NewMailingDispatcher(
		mailingJobQueue, mailingService, numWorkers, mailingResultsChan,
	)
	mailingDispatcher.Run()

	go func() {
		defer close(mailingResultsChan)

		for {
			err := <-mailingResultsChan
			if err != nil {
				fmt.Println("Cannot send email", err.Error())
			}
		}
	}()
}
