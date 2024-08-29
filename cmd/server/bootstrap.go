package main

import (
	"context"
	"log"
	"os"

	"github.com/erenyusufduran/wasnon/internal/config"
	"github.com/erenyusufduran/wasnon/internal/database"
	"github.com/erenyusufduran/wasnon/internal/repositories"
	"github.com/erenyusufduran/wasnon/internal/workers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type App struct {
	DB           *gorm.DB
	Server       *echo.Echo
	Repositories *repositories.Repositories
}

// Initialize bootstraps the application
func Initialize() (*App, error) {
	// Load configuration
	config.Load()

	// Initialize the database
	db := database.Init()

	// Initialize repositories
	repos := repositories.New(db)

	// Initialize the server
	e := InitializeServer(db, repos)

	// Initialize workers
	configs := workers.NewWorkerConfigs(repos)
	err := workers.Initialize(configs)
	if err != nil {
		return &App{}, err
	}

	return &App{
		DB:           db,
		Server:       e,
		Repositories: repos,
	}, nil
}

// StartServer starts the Echo server
func (app *App) StartServer() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return app.Server.Start(":" + port)
}

// Shutdown handles the graceful shutdown of the application
func (app *App) Shutdown() {
	workers.StopAll(true)

	if err := app.Server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	database.Close()
	log.Println("Application shutdown complete")
}
