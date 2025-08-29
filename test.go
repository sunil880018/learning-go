package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ctx = context.Background()
)

type User struct {
	ID   int
	Name string
}

func main() {
	// ✅ Connect to Postgres (only once!)
	pool, err := pgxpool.New(ctx, "postgres://myuser:mypassword@localhost:5445/mydb?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to Postgres: %v\n", err)
	}
	defer pool.Close()
	fmt.Println("Connected to Postgres ✅")

	// ✅ Create table if not exists
	_, err = pool.Exec(ctx, `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
	);`)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	// ✅ Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password
		DB:       0,  // default DB
	})
	defer rdb.Close()
	fmt.Println("Connected to Redis ✅")

	// ✅ Insert a sample user
	var userID int
	err = pool.QueryRow(ctx, "INSERT INTO users (name) VALUES ($1) RETURNING id", "Sunil").Scan(&userID)
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
	}
	fmt.Printf("Inserted user with ID %d ✅\n", userID)

	// ✅ Fetch user with caching
	user, err := getUser(pool, rdb, userID)
	if err != nil {
		log.Fatalf("Error fetching user: %v", err)
	}
	fmt.Printf("Fetched User: %+v ✅\n", user)
}

// getUser tries Redis first, then Postgres if not found
func getUser(pool *pgxpool.Pool, rdb *redis.Client, id int) (*User, error) {
	// Try from Redis
	key := fmt.Sprintf("user:%d", id)
	name, err := rdb.Get(ctx, key).Result()
	if err == nil {
		fmt.Println("Cache hit from Redis ✅")
		return &User{ID: id, Name: name}, nil
	}

	// If not in Redis, fetch from Postgres
	fmt.Println("Cache miss → fetching from Postgres ❌")
	row := pool.QueryRow(ctx, "SELECT id, name FROM users WHERE id=$1", id)

	var u User
	err = row.Scan(&u.ID, &u.Name)
	if err != nil {
		return nil, err
	}

	// Save to Redis with TTL
	err = rdb.Set(ctx, key, u.Name, 10*time.Minute).Err()
	if err != nil {
		log.Printf("Error caching user: %v", err)
	}

	return &u, nil
}
