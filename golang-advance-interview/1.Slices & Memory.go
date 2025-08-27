// What happens when you assign a slice to another slice? Are they deep copied or shallow copied?
// Answer: Slices are shallow copied.

package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a // shallow copy: b points to the same underlying array
	b[0] = 99
	fmt.Println("a:", a) // Output: a: [99 2 3]
	fmt.Println("b:", b) // Output: b: [99 2 3]
	c := deepCopy(a)     // deep copy: c has its own underlying array
	c[0] = 100
	fmt.Println("a:", a) // Output: a: [99 2 3]
	fmt.Println("c:", c) // Output: c: [100 2 3]
}

// dep copy
func deepCopy(src []int) []int {
	dst := append([]int{}, src...)
	return dst
}
