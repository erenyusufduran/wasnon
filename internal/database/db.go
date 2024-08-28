package database

import (
	"fmt"
	"log"
	"os"

	"github.com/erenyusufduran/wasnon/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Init initializes the database connection
func Init() *gorm.DB {
	// Retrieve database connection details from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Construct the PostgreSQL DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	// Connect to PostgreSQL database using Gorm
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&models.Company{}, &models.Employee{}, &models.Product{}); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	log.Println("Database connection successfully established and models migrated")
	return db
}

func Close() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error retrieving SQL database instance:", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatal("Error closing database connection:", err)
	}

	log.Println("Database connection closed")
}
