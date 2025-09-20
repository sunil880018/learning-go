// single.go
package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}
	fmt.Scanln() // wait for input to exit ,basically use to hold screen
}

// ✅ Concept

// 	Ensure only one instance of a resource exists across the entire application.
// 	Useful for shared resources (like DB connection, config, logger).

// ✅ Real-World Go Use Cases

// 	Database connection pool (sql.DB, gorm.DB)

// 	You don’t want 100 different DB connections being created everywhere → you keep a single shared instance.

// 	Logger (zap.Logger, log.Logger)

// 	Consistency in logging format and central control.

// 	Configuration manager (loading .env, YAML config once).

// 	Load config once at startup and reuse everywhere.

// 	Cache clients (Redis, Memcached).

// 	Single connection object shared by all goroutines.
