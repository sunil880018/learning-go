package main

import "fmt"

// Defining an Interface

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Implementing an Interface
// To implement an interface, a type must define all the methods listed in the interface.
// Unlike other languages, Go does not require explicit declarations that a type implements an interface.

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// Using an Interface
// You can use an interface to hold any value that implements it. For instance:
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %f, Perimeter: %f\n", s.Area(), s.Perimeter())
}

func main() {
	c := Circle{Radius: 5}
	PrintShapeInfo(c)
}

// Interfaces are satisfied implicitly: If a type implements all methods required by an interface,
// it satisfies that interface automatically.

// Interfaces in Go are powerful for enabling polymorphism, making it easier to write modular and testable code.
