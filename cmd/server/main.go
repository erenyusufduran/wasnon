package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// go run ./cmd/server
func main() {
	// Initialize the application components
	app := Initialize()

	// Start the server in a separate goroutine to catch gracefully shutdown
	go func() {
		if err := app.StartServer(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start the server: %v", err)
		}
	}()

	// Wait for an interrupt signal to gracefully shutdown the server
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	log.Println("Shutting down server...")

	app.Shutdown()
}
