package main

import (
	"fmt"
	"product-lambda/internal/config"
	"product-lambda/internal/db"
	"product-lambda/internal/service"
	"product-lambda/internal/utils"
)

func main() {
	cfg := config.Load()
	pg := db.InitPostgres(cfg.PostgresURL)
	redis := db.InitRedis(cfg.RedisAddr)

	products, err := utils.ParseCSV("sample_products.csv")
	fmt.Println(products)
	if err != nil {
		panic(err)
	}

	if err := service.UpsertProducts(pg, redis, products); err != nil {
		panic(err)
	}

	fmt.Println("✅ Products uploaded and cache updated.")
}
// psql -U "product-lambda" -d products