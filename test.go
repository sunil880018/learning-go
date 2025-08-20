package main

import (
	"fmt"
	"sync"
)

func printMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		fmt.Println(msg)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // We have 2 goroutines to wait for

	go printMessage("Hello, Goroutine 1!", &wg)
	go printMessage("Hello, Goroutine 2!", &wg)

	wg.Wait() // Wait for both goroutines to finish
	fmt.Println("All goroutines completed!")
}

// If we run multiple goroutines, we just increment wg.Add() accordingly:

// wg.Add(2) ensures main() waits for both goroutines.
// The program doesn't exit until both complete.

// Key Takeaways
// ✅ sync.WaitGroup ensures all goroutines complete before exiting.
// ✅ wg.Add(n) tells how many goroutines we are waiting for.
// ✅ wg.Done() must be called inside each goroutine to decrement the counter.
// ✅ wg.Wait() blocks until all goroutines finish.
