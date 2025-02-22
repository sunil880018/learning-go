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
	fmt.Println("Hello Dost")
	fmt.Println("Hello Dost")
	fmt.Println("Hello Dost")
	fmt.Println("Hello Dost")

	wg.Wait() // Wait for all goroutines to finish before exiting
}

// 1.Why Does "Hello Dost" Print First?
// In your program, "Hello Dost" prints first because the main function does not wait for the goroutine to start
// before executing the next lines of code

// 2.Goroutine is started with go printMessage("Hello, Goroutine!", &wg)

// The goroutine is scheduled asynchronously, meaning it runs in the background.
// However, it doesn't start immediately; the Go scheduler decides when to start it.
