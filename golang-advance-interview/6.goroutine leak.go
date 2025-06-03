// Can a goroutine leak? Explain how and when this might happen.
// Yes, a goroutine can leak if it's blocked forever (e.g., waiting for a channel that is never closed or written to).

// Example of goroutine leak:

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// This goroutine will leak because it's waiting for a value from a channel
	go func() {
		x := <-ch                   // Blocked forever: no one writes to ch
		fmt.Println("Received:", x) // This line will never execute
	}()

	fmt.Println("Main exits, but goroutine is stuck")
	time.Sleep(2 * time.Second) // Give some time to show goroutine is blocked
}

// correct make buffer channel or context.WithTimeout(context.Background(), 2*time.Second)
func main() {
	ch := make(chan int, 1)

	go func() {
		ch <- 4
		x := <-ch
		fmt.Println("Received:", x)
	}()

	fmt.Println("Main exits, but goroutine is stuck")
	time.Sleep(2 * time.Second)
}
