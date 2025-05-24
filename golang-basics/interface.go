package main

import "fmt"

// Define an interface
type Speaker interface {
	Speak()
}

// Implement the interface on type Dog
type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

// Implement the interface on type Human
type Human struct{}

func (h Human) Speak() {
	fmt.Println("Hello!")
}

func makeItSpeak(s Speaker) {
	s.Speak() // Polymorphic call
}

func main() {
	makeItSpeak(Dog{})   // Output: Woof!
	makeItSpeak(Human{}) // Output: Hello!
}

// Polymorphism allows different types to be treated as the same interface type,
// as long as they implement the required methods.
