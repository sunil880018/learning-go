package main

import "fmt"

func test() error {
	var err *MyError = nil
	return err
}

type MyError struct{}

func (e *MyError) Error() string {
	return "My error"
}

func main() {
	if test() == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

// Answer: not nil

// Explanation:

// The interface error holds a type and a value.

// In this case, it holds (*MyError)(nil), which is not nil as an interface.
