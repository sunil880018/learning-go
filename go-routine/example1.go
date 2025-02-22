package main

import (
	"fmt"
	"sync"
	"time"
)

func printMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when function finishes
	for i := 1; i <= 3; i++ {
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // Tell WaitGroup to wait for 1 goroutine

	go printMessage("Hello, Goroutine!", &wg) // // Start a goroutine , The &wg (pointer) is passed to the function so it can call wg.Done() when finished.

	wg.Wait() // Wait for all goroutines to finish before exiting
}

// The Go runtime does not wait for goroutines to finish unless explicitly told to.
// The main goroutine (main function) terminates, causing all other child goroutines to be stopped immediately.

// sync.WaitGroup is used in Go to wait for multiple goroutines to complete before the main function exits.
