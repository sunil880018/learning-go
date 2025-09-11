package main

import (
	"context"
	"encoding/json"
	"product-lambda/internal/config"
	"product-lambda/internal/db"
	"product-lambda/internal/service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cfg := config.Load()
	redis := db.InitRedis(cfg.RedisAddr)

	products, err := service.GetCachedProducts(redis)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       `{"error": "Redis unavailable"}`,
		}, nil
	}

	body, _ := json.Marshal(products)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Cache-Control": "max-age=60"},
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(handler)
}
