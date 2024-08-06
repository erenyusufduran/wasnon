package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/erenyusufduran/wasnon/internal/config"
	"github.com/erenyusufduran/wasnon/internal/database"
	"github.com/erenyusufduran/wasnon/internal/repositories"
	"github.com/erenyusufduran/wasnon/internal/server"
	"github.com/erenyusufduran/wasnon/internal/workers"
	"github.com/labstack/echo/v4"
)

func main() {
	config.Load()

	db := database.Init()

	productRepo := repositories.NewGormProductRepository(db)
	companyRepo := repositories.NewGormCompanyRepository(db)

	e := server.Initialize(db, productRepo, companyRepo)
	workers.Initialize(db, productRepo)

	startServer(e)
}

func startServer(e *echo.Echo) {
	// Run server in a goroutine so that it doesn't block the graceful shutdown handling
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		log.Printf("Starting server on port %s", port)
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	log.Println("Shutting down server...")

	workers.StopAll()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	database.Close()
	log.Println("Server exiting")
}
