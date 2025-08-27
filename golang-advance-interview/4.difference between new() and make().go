// Explain the difference between new() and make() in Go. When should you use each?
// new(Type) allocates zeroed memory and returns a pointer.

// make(Type, ...) initializes built-in types like slices, maps, and channels.
package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p) // 0

	s := make([]int, 3)
	fmt.Println(s) // [0 0 0]
}
