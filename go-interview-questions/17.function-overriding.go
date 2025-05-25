package main

import "fmt"

type Animal struct{}

func (a Animal) Speak() {
	fmt.Println("Animal speaks")
}

type Dog struct {
	Animal
}

func (d Dog) Speak() { // This "overrides" Animal.Speak
	fmt.Println("Dog barks")
}

func main() {
	d := Dog{}
	d.Speak()        // Dog's Speak is called
	d.Animal.Speak() // Call base version if needed
}
