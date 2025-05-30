// What is the memory layout of structs in Go? How does field order affect memory usage?
// Answer: Go aligns struct fields to minimize padding. Field order affects size.

package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	a int8
	b int64
	c int8
}
type B struct {
	a int8
	c int8
	b int64
}

func main() {
	fmt.Println("Size of A:", unsafe.Sizeof(A{})) // Output: 24
	fmt.Println("Size of B:", unsafe.Sizeof(B{})) // Output: 16 (better)
}
