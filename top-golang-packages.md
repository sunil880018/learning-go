# 📚 Essential Go Packages

This guide covers the **most commonly used Go standard library packages** and the **top external libraries** you should master for real-world development.

---

## 🔹 Core Utilities

- **`fmt`** → Printing & formatting (`fmt.Println`, `fmt.Sprintf`)
- **`errors`** → Error handling (`errors.New`, `errors.Is`)
- **`log`** → Basic logging

---

## 🔹 Strings, Numbers, Time

- **`strings`** → String operations (`strings.Split`, `strings.Contains`)
- **`strconv`** → Convert string ↔ numbers (`strconv.Atoi`, `strconv.Itoa`)
- **`time`** → Working with time, timers, and sleep (`time.Now`, `time.Sleep`)

---

## 🔹 Concurrency

- **`sync`** → Concurrency control (`sync.Mutex`, `sync.WaitGroup`, `sync.Cond`)
- **`sync/atomic`** → Lock-free atomic counters & flags
- **`context`** → Cancellation & timeouts in goroutines

---

## 🔹 Collections / Data

- **`sort`** → Sorting slices (`sort.Ints`, custom sort)
- **`math` / `math/rand`** → Math functions, random numbers

---

## 🔹 I/O, Files, Networking

- **`io` / `io/ioutil`** _(deprecated → use `os` + `io`)_ → Reading/writing streams
- **`os`** → File, process, environment variables
- **`path/filepath`** → File path handling
- **`net/http`** → Building web servers & clients
- **`encoding/json`** → JSON encoding/decoding
- **`bufio`** → Buffered I/O (faster reading/writing)
- **`bytes`** → Working with byte slices

---

## 🔹 For Testing

- **`testing`** → Go’s built-in test framework

---

# 🚀 Top External Packages

These are **industry favorites** you’ll see in production-ready Go applications.

- **`github.com/gin-gonic/gin`** → High-performance web framework
- **`github.com/gorilla/mux`** → Powerful HTTP router
- **`github.com/stretchr/testify`** → Unit testing & assertions
- **`go.uber.org/zap`** → Fast structured logging
- **`go.uber.org/ratelimit`** → Production-grade rate limiter
- **`github.com/go-redis/redis/v9`** → Redis client
- **`gorm.io/gorm`** → ORM for databases
- **`github.com/Shopify/sarama`** → Kafka client
- **`github.com/spf13/viper`** → Config management
- **`github.com/spf13/cobra`** → CLI applications

---

## 📌 How to Use This Guide

1. Start with the **standard library** packages — they cover 80% of Go development.
2. Learn the **external packages** gradually as you build real-world apps.
3. Apply them in small projects like:
   - REST API with Gin
   - Redis-backed caching system
   - Kafka message producer/consumer
   - CLI tool using Cobra & Viper

---

✨ Mastering these packages will prepare you for **90% of Go jobs and interviews**.
