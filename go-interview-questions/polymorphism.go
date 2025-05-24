package main

import "fmt"

type Speaker interface {
	Speak()
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println(c.Name, "says meow")
}

type Cow struct {
	Name string
}

func (c Cow) Speak() {
	fmt.Println(c.Name, "says moo")
}

func makeItSpeak(s Speaker) {
	s.Speak()
}

func main() {
	cat := Cat{Name: "Whiskers"}
	cow := Cow{Name: "Bessie"}

	makeItSpeak(cat)
	makeItSpeak(cow)
}
