-- Create companies table
CREATE TABLE IF NOT EXISTS companies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS branches (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) UNIQUE NOT NULL,
	address VARCHAR(255) NOT NULL,
	city VARCHAR(20) NOT NULL,
	county VARCHAR(20) NOT NULL,
	company_id INT NOT NULL,
	FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE
)

-- Create products table with foreign key to companies and enum status
CREATE TYPE product_status AS ENUM ('active', 'past', 'deleted', 'waiting');

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    expiration DATE NOT NULL,
    company_id INT NOT NULL,
    status product_status NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE
);