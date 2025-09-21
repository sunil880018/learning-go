// Wait for multiple goroutines to finish.
// Use when: You want the main goroutine to wait until a group of goroutines finish.
// ✅ Use case: Parallel tasks like fetching multiple APIs concurrently, then processing results together.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Task 1 done")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Task 2 done")
	}()

	wg.Wait()
	fmt.Println("All tasks finished")
}
