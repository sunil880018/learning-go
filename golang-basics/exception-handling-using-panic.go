package main

// Go does not have traditional exception handling like try-catch in other languages
import "fmt"

func riskyOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Simulate a panic situation (e.g., divide by zero)
	n, d := 10, 0
	if d == 0 {
		panic("division by zero")
	}
	fmt.Println("Result:", n/d)
}

func main() {
	riskyOperation()
	fmt.Println("Program continues after recovery")
}

// Handling Unexpected Errors: panic and recover
// Go also has a panic-recover mechanism, similar to exceptions but used sparingly for critical or unexpected issues, not for regular error handling.

// panic: Causes the program to stop normal execution and unwinds the stack. Useful for unexpected situations like array index out-of-bounds.
// recover: Catches the panic and allows graceful recovery.
