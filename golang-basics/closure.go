package main

import "fmt"

func main() {
	// A function that returns a closure (anonymous function)
	counter := func() func() int {
		count := 0
		// This returned function closes over the variable 'count'
		return func() int {
			count++ // Increment the outer variable
			return count
		}
	}

	// Create a new counter closure
	increment := counter()

	fmt.Println(increment()) // Output: 1
	fmt.Println(increment()) // Output: 2
	fmt.Println(increment()) // Output: 3
}
