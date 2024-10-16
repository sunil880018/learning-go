package main

import "fmt"

// Defining a struct type
type Rectangle struct {
	width  int
	height int
}

// Defining a method with a value receiver
func (r Rectangle) area() int { // Value Receiver
	return r.width * r.height
}

// Defining a method with a pointer receiver
func (r *Rectangle) scale(factor int) {
	r.width *= factor
	r.height *= factor
}

func main() {
	rect := Rectangle{width: 10, height: 5}
	fmt.Println("Area:", rect.area()) // Calls the area method
	rect.scale(2)
	fmt.Println("Area:", rect)
}

// When to Use Value vs Pointer Receiver

// Value Receiver: Use this when the method does not modify the object or when the object is small
// (like int, struct with few fields).

// Pointer Receiver: Use this when the method needs to modify the object or if the object is large and copying
// it would be inefficient.
