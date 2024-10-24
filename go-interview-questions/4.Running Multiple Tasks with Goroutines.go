package main

import (
	"fmt"
	"time"
)

func task(name string, delay time.Duration) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Task %s is running iteration %d\n", name, i)
		time.Sleep(delay)
	}
}

func main() {
	// Running multiple tasks concurrently using goroutines
	go task("A", 1*time.Second)        // Task A runs every 1 second
	go task("B", 500*time.Millisecond) // Task B runs every 500 milliseconds
	go task("C", 200*time.Millisecond) // Task C runs every 200 milliseconds

	// Let the main function wait before it finishes
	time.Sleep(6 * time.Second) // Wait for enough time for the goroutines to finish
	fmt.Println("All tasks completed")
}
