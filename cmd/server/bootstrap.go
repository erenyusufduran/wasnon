package main

import (
	"context"
	"log"
	"os"

	"github.com/erenyusufduran/wasnon/internal/config"
	"github.com/erenyusufduran/wasnon/internal/database"
	"github.com/erenyusufduran/wasnon/pkg/worker"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type App struct {
	DB           *pgxpool.Pool
	Server       *echo.Echo
	Repositories *Repositories
}

// Initialize bootstraps the application
func Initialize() (*App, error) {
	// Load configuration
	config.Load()

	// Initialize the database
	db := database.Init()

	// Initialize repositories
	repos := NewRepositories(db)

	// Initialize the server
	e := InitializeServer(repos)

	// Initialize workers
	configs := NewWorkerConfigs(repos)
	err := worker.Initialize(configs)

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
	worker.StopAll(true)

	if err := app.Server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	database.Close()
	log.Println("Application shutdown complete")
}
