package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from ch1"
	}()

	// Goroutine 2
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from ch2"
	}()

	// Multiplexing channels
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}

// output either print Message from ch2 or Message from ch1
