package config

import (
	"log"

	"github.com/joho/godotenv"
)

// Load loads environment variables from a .env file
func Load() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}
