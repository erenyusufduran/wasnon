package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

// Init initializes the database connection
func Init() *pgxpool.Pool {
	// Retrieve database connection details from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Construct the PostgreSQL DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Unable to parse config: %v", err)
	}

	config.MaxConns = 10
	config.MinConns = 5
	config.MaxConnLifetime = time.Hour

	db, err = pgxpool.New(context.Background(), config.ConnString())
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	err = runMigrations(db)
	if err != nil {
		log.Fatalf("Unable to do migrations: %v", err)
	}

	log.Println("Database connection successfully established")
	return db
}

func Close() {
	if db != nil {
		db.Close()
		log.Println("Database connection closed")
	}
}
