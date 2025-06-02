package main

import (
	"fmt"
)

// Producer function
func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // send to channel
	}
	close(ch) // close channel when done
}

// Consumer function
func consumer(ch chan int, done chan bool) {
	for val := range ch {
		fmt.Println("Consumed:", val)
	}
	done <- true
}

func main() {
	ch := make(chan int)
	done := make(chan bool)

	// Start producer and consumer goroutines
	go producer(ch)
	go consumer(ch, done)

	// Wait for consumer to finish
	<-done
}

// output
// Consumed: 1
// Consumed: 2
// Consumed: 3
// Consumed: 4
// Consumed: 5
