package main

import (
	"fmt"
	"log"
	"net/http"

	"time"

	"go.uber.org/ratelimit"
)

// --- Middleware ---
func rateLimitMiddleware(rl ratelimit.Limiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rl.Take() // blocks until allowed
		waited := time.Since(start)

		log.Printf("Request %s %s → waited %v before serving", r.Method, r.URL.Path, waited)

		next.ServeHTTP(w, r)
	})
}

// --- API Handlers ---
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home API → OK")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User API → OK")
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Order API → OK")
}

func main() {
	// --- One global limiter for ALL APIs ---
	globalLimiter := ratelimit.New(1) // 200 req/sec total

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/user", userHandler)
	mux.HandleFunc("/order", orderHandler)

	// Wrap the entire mux with the global limiter
	handler := rateLimitMiddleware(globalLimiter, mux)
	// --- Start server ---
	log.Println("server listening on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
