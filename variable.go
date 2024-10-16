package main

import (
	"fmt"
)

func main() {
	// Declaring a variable with 'var'
	var x int = 10
	fmt.Println("Value of x:", x)

	// Declaring and initializing a variable (type inferred)
	var y = "Hello, Go!"
	fmt.Println("Value of y:", y)

	// Using shorthand syntax (:=) to declare and initialize
	z := 3.14
	fmt.Println("Value of z:", z)

	// Declaring multiple variables
	a, b := 100, "Multiple variables"
	fmt.Println("Value of a:", a, "and b:", b)

	// Block declaration
	var (
		num   int    = 42
		lang  string = "Golang"
		check bool   = true
	)
	fmt.Println("num:", num, "lang:", lang, "check:", check)
}
