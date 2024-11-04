package main

import (
	"fmt"
	"sync"
	"time"
)

func printMessage(message string, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done

	for i := 1; i <= 5; i++ {
		fmt.Println(message, i)
		time.Sleep(500 * time.Millisecond) // Simulate work
	}
}

func main() {
	var wg sync.WaitGroup

	// Add two goroutines to the WaitGroup
	wg.Add(2)

	// Start the first goroutine
	go printMessage("Hello from goroutine 1", &wg)

	// Start the second goroutine
	go printMessage("Hello from goroutine 2", &wg)

	// Wait for both goroutines to complete
	wg.Wait()
	fmt.Println("Main function ends")
}

// Explanation
// WaitGroup: wg.Add(2) tells wg that it should wait for two tasks to finish. Each printMessage function calls wg.Done() when it finishes, reducing the counter by one.
// wg.Wait(): This line blocks the main function until the WaitGroup counter reaches zero (i.e., until all goroutines have called Done).
