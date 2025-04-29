package main

import "fmt"

func main() {
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
}

// output
// Third
// Second
// First

// Explanation:

// Deferred functions execute in LIFO (Last-In, First-Out) order.
