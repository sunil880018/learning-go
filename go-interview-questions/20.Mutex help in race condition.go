package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

// Yes, sync.Mutex in Go helps prevent race conditions.
// A race condition happens when two or more goroutines access shared data
// (like a variable, slice, or map) at the same time,
func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go increment(&wg)
	wg.Wait()
	fmt.Println(counter)

	wg.Add(1)
	go increment(&wg)
	wg.Wait()
	fmt.Println(counter)

	wg.Add(1)
	go increment(&wg)
	wg.Wait()
	fmt.Println(counter)
}

// output
// 1
// 2
// 3

// Use Mutex:

// 1.When multiple goroutines access shared variables/data.

// 2.To prevent race conditions, data corruption, or inconsistent state.

// 3.For synchronizing access to critical sections (balance, inventory, config, etc.).
