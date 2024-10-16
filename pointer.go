package main

import "fmt"

func increment(p *int) {
	*p += 1
}

// Pointers and Structs
// Pointers are especially useful when working with structs, where copying entire structs can be inefficient.

type Person struct {
	name string
	age  int
}

func birthday(p *Person) {
	p.age += 1
}

func main() {
	var p *int
	x := 5
	p = &x         // p is a pointer to the variable x
	fmt.Println(p) // Prints the memory address of x

	fmt.Println(*p) // Prints the value of x (5)
	*p = 10         // Changes the value of x to 10
	fmt.Println(x)  // Prints 10

	increment(&x)
	fmt.Println(x) // Prints 11

	person := Person{name: "John", age: 30}
	birthday(&person)
	fmt.Println(person.age) // Prints 31

	// The zero value of a pointer is nil. You should check if a pointer is nil before dereferencing it to avoid
	// runtime panics:

	if p != nil {
		fmt.Println(*p)
	}
}
