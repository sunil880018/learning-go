// Ensures a function runs only once (useful for initialization).
// Use when: You want a piece of code to run exactly once, even with multiple goroutines.
// ✅ Use case: Initializing singletons, setting up a config, opening DB connections.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	for i := 0; i < 3; i++ {
		once.Do(func() {
			fmt.Println("Hello executed once")
		})
	}
}
