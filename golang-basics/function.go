package main

import "fmt"

// Defining a function that takes two integers and returns their sum
func add(x int, y int) int {
	return x + y
}

// Function that returns two values
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return x / y, nil
}

func main() {
	// Calling the function
	result := add(5, 3)
	divisionRessult, err := divide(10, 3)
	fmt.Println(result) // Output: 8

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", divisionRessult) // Output: divisionRessult: 3
	}

	// Defining and invoking an anonymous function
	func(message string) {
		fmt.Println(message)
	}("Hello, World!")

	// Assigning an anonymous function to a variable
	greet := func(name string) string {
		return "Hello, " + name
	}

	fmt.Println(greet("Go")) // Output: Hello, Go

	// A closure that captures the variable 'counter'
	counter := 0
	increment := func() int {
		counter++
		return counter
	}

	fmt.Println(increment()) // Output: 1
	fmt.Println(increment()) // Output: 2

}
