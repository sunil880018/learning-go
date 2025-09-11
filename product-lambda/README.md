<!-- Connect to PostgreSQL -->
<!-- Command: psql -U "product-lambda" -d products -->
<!-- Password: product123 -->

# Product Lambda Project

A lightweight Go application to manage product data with PostgreSQL and Redis integration.

### Key Features

1.  **Upload Product Data** from a CSV file into PostgreSQL and Redis.
2.  **Retrieve All Products** from Redis (fallback to PostgreSQL if cache miss).

---

## Prerequisites

Ensure the following tools are installed and running locally:

- Go 1.18+
- PostgreSQL (port `5432`)
- Redis (port `6379`)
- `products` PostgreSQL database and user access

---

## PostgreSQL Setup

### Step 1: Create Database & User

Run these SQL commands in `psql`:

```sql
CREATE DATABASE products;

CREATE USER "product-lambda" WITH PASSWORD 'product123';

GRANT CONNECT ON DATABASE products TO "product-lambda";
GRANT USAGE ON SCHEMA public TO "product-lambda";
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO "product-lambda";
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO "product-lambda";
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO "product-lambda";
GRANT CREATE ON SCHEMA public TO "product-lambda";


CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUE NOT NULL,
  image TEXT,
  price NUMERIC(10, 2) NOT NULL,
  quantity INT NOT NULL,
  out_of_stock BOOLEAN DEFAULT FALSE
);

POSTGRES_URL=postgres://product-lambda:product123@localhost:5432/products?sslmode=disable
REDIS_ADDR=localhost:6379


go run cmd/uploadProduct/main_local.go
go run cmd/getAllProducts/main_local.go


type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Image      string  `json:"image"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"qty"`
	OutOfStock bool    `json:"out_of_stock"`
}
```
