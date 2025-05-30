// Can a goroutine leak? Explain how and when this might happen.
// Yes, a goroutine can leak if it's blocked forever (e.g., waiting for a channel that is never closed or written to).

// Example of goroutine leak:

package main

import "fmt"

func main() {
	ch := make(chan int)
	// This goroutine waits forever, leaking memory
	go func() {
		<-ch
		fmt.Println("Unreachable")
	}()

	fmt.Println("Main exits, but goroutine is stuck")
}
