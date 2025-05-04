package main

import "fmt"

func test() (result int) {
	c:= defer func() {
		result++
	}()
	fmt.Println(c)
	return 2
}

func main() {
	fmt.Println(test())
}

// Answer: 3

// Explanation:

// result is a named return variable. Defer modifies it before the function exits.
