package main

import "fmt"

func test() (result int) {
	defer func() {
		result++
	}()
	return 2
}

func main() {
	fmt.Println(test())
}

// Answer: 3

// Explanation:

// result is a named return variable. Defer modifies it before the function exits.
