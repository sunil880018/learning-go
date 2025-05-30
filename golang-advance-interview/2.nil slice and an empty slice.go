// What is the difference between a nil slice and an empty slice?
package main

import "fmt"

func main() {
	var nilSlice []int    // nil
	emptySlice := []int{} // empty but not nil

	fmt.Println(nilSlice == nil)   // true
	fmt.Println(emptySlice == nil) // false
}

// nilSlice: Has no memory allocated.
// emptySlice: Allocated, but with zero length.
