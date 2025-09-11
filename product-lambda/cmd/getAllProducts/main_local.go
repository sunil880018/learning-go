package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"product-lambda/internal/config"
	"product-lambda/internal/db"
	"product-lambda/internal/service"
)

func main() {
	cfg := config.Load()
	redis := db.InitRedis(cfg.RedisAddr)

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		products, err := service.GetCachedProducts(redis)
		if err != nil {
			http.Error(w, "Redis error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	})

	fmt.Println("🚀 Server running at http://localhost:8080/products")
	http.ListenAndServe(":8080", nil)
}
