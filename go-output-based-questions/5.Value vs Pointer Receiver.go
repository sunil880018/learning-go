package main

import "fmt"

type Counter struct {
	count int
}

func (c Counter) Increment() {
	c.count++
}

func (c *Counter) Decrement() {
	c.count--
}

func main() {
	c := Counter{}
	c.Increment()
	c.Decrement()
	fmt.Println(c.count)
}

// Answer: -1

// Explanation:
// Increment() uses a value receiver (works on a copy).
// Decrement() uses a pointer receiver (modifies original).
