// A mutex ensures that only one goroutine can hold the lock and access the shared resource at any given time.
// This is crucial when multiple goroutines are reading from and writing to shared memory or variables.

// A mutex has two main methods:
// Lock(): Locks the mutex, preventing other goroutines from accessing the critical section.
// Unlock(): Unlocks the mutex, allowing other goroutines to acquire the lock.

// Types of Mutexes:

// sync.Mutex: This is the standard mutex in Go, providing simple mutual exclusion. Once locked,
// any other goroutine trying to acquire the lock will block until the mutex is unlocked.

// sync.RWMutex: This is a read-write mutex. It allows multiple readers to hold the lock simultaneously, but only
// one writer can hold the lock, and no readers are allowed when a writer holds the lock.

package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	// Lock the mutex to ensure exclusive access to the counter
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	fmt.Println("World")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
	fmt.Println("Hello")
}

// When to Use a Mutex:
// Use a mutex when you need to safely read from or write to shared resources from multiple goroutines.
// If you're only reading shared data from multiple goroutines, sync.RWMutex can be used to allow concurrent reads
// but still synchronize writes.
