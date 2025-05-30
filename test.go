package main

import "fmt"

func main() {
	ch := make(chan int)
	// This goroutine waits forever, leaking memory
	defer close(ch)
	go func() {
		ch <- 1
		fmt.Println("Unreachable")
		close(ch)
	}()

	fmt.Println("Main exits, but goroutine is stuck")
}
