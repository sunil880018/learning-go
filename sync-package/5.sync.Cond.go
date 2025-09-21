// A condition variable: allows goroutines to wait until they’re signaled.
// Use when: You have a producer-consumer scenario or need goroutines to wait until a condition is met.
// ✅ Use case: Queue workers waiting for items, signaling state changes between goroutines.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	done := false

	// Waiter goroutine
	go func() {
		cond.L.Lock()
		for !done {
			cond.Wait() // wait until condition is true
		}
		fmt.Println("Condition met!")
		cond.L.Unlock()
	}()

	time.Sleep(1 * time.Second)
	cond.L.Lock()
	done = true
	cond.Signal() // wake up one waiter
	cond.L.Unlock()
}
