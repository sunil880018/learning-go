package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func(i int) {
			ch <- i
			fmt.Println("Sent:", i)
		}(i)
	}
	// time.Sleep(1 * time.Second)
	// close(ch)
	for i := 0; i < 3; i++ {
		v := <-ch
		fmt.Println("Received:", v)
	}
}
