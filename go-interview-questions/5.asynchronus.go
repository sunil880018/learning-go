package main

import (
	"fmt"
	"time"
)

func printMessage(message string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(message, i)
		time.Sleep(500 * time.Millisecond) // Simulate work
	}
}

func main() {
	// Start the first goroutine
	go printMessage("Hello from goroutine 1")

	// Start the second goroutine
	go printMessage("Hello from goroutine 2")

	// Keep the main function alive while the goroutines run
	time.Sleep(1 * time.Second)
	fmt.Println("Main function ends")
}
