package main

import "fmt"

type Person struct {
	Name string
}

// Method on Person
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s\n", p.Name)
}

type Student struct {
	Person
}

func main() {
	s := &Student{Person{Name: "Sunil"}}

	// Call method from embedded struct
	s.Greet()
}

