package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"product-lambda/internal/models"

	"github.com/go-redis/redis/v8"
)

func UpsertProducts(dbConn *sql.DB, rdb *redis.Client, products []models.Product) error {
	tx, err := dbConn.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO products (name, image, price, quantity, out_of_stock)
    VALUES ($1, $2, $3, $4, $5)
    ON CONFLICT(name) DO UPDATE
    SET image=$2, price=$3, quantity=$4, out_of_stock=$5`)
	if err != nil {
		return err
	}

	for _, p := range products {
		_, err := stmt.Exec(p.Name, p.Image, p.Price, p.Quantity, p.OutOfStock)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	// Update Redis cache
	jsonData, _ := json.Marshal(products)
	return rdb.Set(context.Background(), "all_products", jsonData, 0).Err()
}

func GetCachedProducts(rdb *redis.Client) ([]models.Product, error) {
	val, err := rdb.Get(context.Background(), "all_products").Result()
	if err != nil {
		return nil, err
	}

	var products []models.Product
	err = json.Unmarshal([]byte(val), &products)
	return products, err
}
