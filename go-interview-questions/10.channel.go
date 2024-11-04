package main

import (
	"fmt"
)

func main() {
	numChannel := make(chan int, 10) // Create a buffered channel with capacity 10

	// Sending data to the channel
	for i := 1; i <= 10; i++ {
		numChannel <- i // This won't block until the buffer is full
		fmt.Println("Sent:", i)
	}

	close(numChannel) // Close the channel to indicate no more data will be sent

	// Receiving data from the channel
	for num := range numChannel {
		fmt.Println("Received:", num)
	}
}
