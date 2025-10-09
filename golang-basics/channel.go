// In Go, channels are a built-in feature used for goroutine communication.
// They enable safe data exchange between goroutines without using explicit locking (like sync.Mutex).

// A channel is a typed conduit that allows goroutines to send and receive values.
// Think of it like a pipe—one goroutine sends data into the pipe, and another goroutine receives it.

package main

import (
	"fmt"
)

func main() {
	ch := make(chan string) // Create a string channel

	// Goroutine to send data
	go func() {
		ch <- "Hello from Goroutine!" // Send data into the channel
	}()

	// Receiving data
	msg := <-ch // Receive data from the channel
	fmt.Println(msg)
}
