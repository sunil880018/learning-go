package main

import (
	"context"
	"product-lambda/internal/config"
	"product-lambda/internal/db"
	"product-lambda/internal/service"
	"product-lambda/internal/utils"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	FilePath string `json:"file_path"`
}

func handler(ctx context.Context, req Request) (string, error) {
	cfg := config.Load()
	pg := db.InitPostgres(cfg.PostgresURL)
	redis := db.InitRedis(cfg.RedisAddr)

	products, err := utils.ParseCSV(req.FilePath)
	if err != nil {
		return "Failed to parse CSV", err
	}

	if err := service.UpsertProducts(pg, redis, products); err != nil {
		return "DB or Redis update failed", err
	}

	return "Products uploaded successfully", nil
}

func main() {
	lambda.Start(handler)
}
