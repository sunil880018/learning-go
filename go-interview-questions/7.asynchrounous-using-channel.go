package main

import (
	"fmt"
	"time"
)

func printMessage(message string, done chan bool) {
	for i := 1; i <= 5; i++ {
		fmt.Println(message, i)
		time.Sleep(500 * time.Millisecond) // Simulate work
	}
	done <- true // Signal that this goroutine is done
}

func main() {
	done := make(chan bool, 2) // Buffered channel to handle two signals

	// Start the first goroutine
	go printMessage("Hello from goroutine 1", done)

	// Start the second goroutine
	go printMessage("Hello from goroutine 2", done)

	// Wait for both goroutines to signal completion
	<-done
	<-done
	fmt.Println("Main function ends")
}

// Explanation
// Channel: done is a buffered channel that can hold two values, allowing each goroutine to send a completion signal without blocking.
// <-done: This waits to receive a value from each goroutine, effectively blocking until both signals are received, indicating both goroutines have completed.
// Both approaches ensure main doesn’t exit until all goroutines have finished. sync.WaitGroup is generally preferred for managing multiple concurrent tasks but channels are also effective for communication between goroutines.
