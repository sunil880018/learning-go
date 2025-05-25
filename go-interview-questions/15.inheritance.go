package main

import "fmt"

// Base struct
type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println(a.Name, "makes a sound")
}

type Dog struct {
	Animal // embedded struct (inheritance-like)
	Breed  string
}

func (d Dog) Bark() {
	fmt.Println(d.Name, "barks! Woof!")
}

func main() {
	d := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}

	d.Speak() // inherited from Animal
	d.Bark()  // Dog's own method
}
