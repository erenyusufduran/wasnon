package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// migrateCompanies creates the companies table
func migrateCompanies() string {
	query := `
		CREATE TABLE IF NOT EXISTS companies (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL
		)`
	return query
}

// migrateBranches creates the branches table with a foreign key to companies
func migrateBranches() string {
	query := `
		CREATE TABLE IF NOT EXISTS branches (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			address VARCHAR(255) NOT NULL,
			city VARCHAR(20) NOT NULL,
			county VARCHAR(20) NOT NULL,
			company_id INT NOT NULL,
			FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE
		)`
	return query
}

// migrateProductStatusEnum creates the product_status enum
func migrateProductStatusEnum() string {
	query := `
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'product_status') THEN
				CREATE TYPE product_status AS ENUM ('active', 'past', 'deleted', 'waiting');
			END IF;
		END $$;`
	return query
}

// migrateProducts creates the products table with a foreign key to companies
func migrateProducts() string {
	query := `
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description TEXT NOT NULL,
			price DECIMAL(10, 2) NOT NULL,
			expiration DATE NOT NULL,
			company_id INT NOT NULL,
			status product_status NOT NULL,
			FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE
		)`
	return query
}

// runMigrations creates the necessary tables and types
func runMigrations(db *pgxpool.Pool) error {
	queries := []string{
		migrateCompanies(),
		migrateBranches(),
		migrateProductStatusEnum(),
		migrateProducts(),
	}

	// Execute each query
	for _, query := range queries {
		_, err := db.Exec(context.Background(), query)
		if err != nil {
			return fmt.Errorf("failed to run migration: %w", err)
		}
	}

	return nil
}
