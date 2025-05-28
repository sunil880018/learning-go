package main

import "fmt"

func main() {
	ch := make(chan int, 5) // buffered channel with capacity 5

	// Sending 5 values — won't block
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Println("Sent:", i)
	}

	// Next send would block if uncommented
	// ch <- 6

	// Receiving values
	for i := 1; i <= 5; i++ {
		val := <-ch
		fmt.Println("Received:", val)
	}
}
