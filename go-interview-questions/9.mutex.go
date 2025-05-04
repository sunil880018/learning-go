package main

import (
	"fmt"
	"sync"
)

// mutex is  prevent multiple goroutines from accessing shared data at the same time.
// Counter with a mutex for safe concurrent access
type Counter struct {
	mu    sync.Mutex // Mutex to synchronize access
	value int
}

// Increment increments the counter's value by 1
func (c *Counter) Increment() {
	c.mu.Lock()   // Acquire the lock
	c.value++     // Modify the shared value
	c.mu.Unlock() // Release the lock
}

// Value returns the counter's current value
func (c *Counter) Value() int {
	c.mu.Lock()         // Acquire the lock
	defer c.mu.Unlock() // Ensure the lock is released after returning
	return c.value
}

func main() {
	counter := Counter{} // Initialize the counter
	var wg sync.WaitGroup

	// Start 1000 goroutines that increment the counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Final counter value:", counter.Value())
}

// Mutex (short for "mutual exclusion") is a synchronization primitive that is used to manage concurrent access to shared
// resources. When multiple goroutines need to access or modify the same piece of data, a mutex ensures that only one
// goroutine can access the data at a time, preventing race conditions and ensuring data integrity.

// How a Mutex Works
// A mutex has two main methods:

// Lock(): Acquires the lock. If the lock is already held by another goroutine, the current goroutine will wait (block) until
// the lock becomes available.
// Unlock(): Releases the lock. If other goroutines are waiting to acquire the lock, one of them will be allowed to proceed.
