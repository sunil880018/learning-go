// A goroutine in Go is a lightweight thread managed by the Go runtime. It enables concurrent execution of functions and
// is a core feature of Go's concurrency model. Goroutines are much more efficient than traditional threads,
// allowing developers to run thousands or even millions of them simultaneously without consuming significant system resources.

package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go printMessage("Hello, Goroutine!") // Starts a new goroutine, this is separate goroutine
	printMessage("Main Function")        // run in main goroutine means main thread
	// The Go runtime does not wait for goroutines to finish unless explicitly told to.
	// The main goroutine (main function) terminates, causing all other child goroutines to be stopped immediately.
}

// goroutines often communicate with each other using channels.
// channels provide a way to safely pass data between goroutines and are integral to Go's concurrency model.

// package main

// import "fmt"

// func sayHello(c chan string) {
// 	c <- "Hello, Goroutines!"
// }

// func main() {
// 	ch := make(chan string)
// 	go sayHello(ch)       // Start the goroutine
// 	fmt.Println(<-ch)      // Receive the message from the channel
// }
