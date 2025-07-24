package main

import (
	"context"
	"fmt"
	"time"
)

func printNumbers(ctx context.Context) {
	for i := 1; i <= 5; i++ {
		if ctx.Err() != nil {
			fmt.Println("Stopped by context:", ctx.Err())
			return
		}
		fmt.Println(i)
		time.Sleep(1 * time.Second) // Simulate work
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go printNumbers(ctx)

	time.Sleep(6 * time.Second) // Wait long enough to see context timeout
	fmt.Println("Main done")
}

// output
// 1
// 2
// 3
// Stopped by context: context deadline exceeded
// Main done

// ✅ Why does context help prevent memory leaks?
// When you start a goroutine (background task), if you don’t have a way to stop it, it might:

// keep running forever,

// keep consuming CPU or memory,

// or block on something that will never finish.

// This causes a memory leak or resource leak.

// Using context allows you to cancel or timeout that goroutine.
