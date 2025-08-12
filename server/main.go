package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FrancoMusolino/go-todoapp/db"
	"github.com/FrancoMusolino/go-todoapp/router"
	"github.com/FrancoMusolino/go-todoapp/utils"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Cannot load .env file, using env variables")
	}

	// DB
	_, dbConnection := db.NewDatabase()
	defer dbConnection.Close()

	// Init Repos
	// usersRepo := repos.NewUsersRepo(gormClient)

	// Routes
	router := router.SetUpRouter()
	port := utils.GetEnvOrDefault("port", "8080")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
