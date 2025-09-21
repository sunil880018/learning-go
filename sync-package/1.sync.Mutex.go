package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	counter := 0

	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1
	go func() {
		defer wg.Done()
		mu.Lock()
		counter++
		mu.Unlock()
	}()

	// Goroutine 2
	go func() {
		defer wg.Done()
		mu.Lock()
		counter++
		mu.Unlock()
	}()

	wg.Wait()
	fmt.Println("Counter:", counter) // always 2
}

// A mutual exclusion lock to protect shared resources.
// Only one goroutine can lock the mutex at a time.

// Use when: Multiple goroutines need to modify a shared variable or resource,
// and you want to prevent race conditions.

// ✅ Use case: Incrementing a shared counter, writing to a shared file, updating shared data structures.
