package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := a[:2]
	c := append(b, 10, 11)
	c[0] = 100
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}

// output
// a: [100 2 10 11]
// b: [100 2]
// c: [100 2 10 11]

// Explanation:

// append(b, 10, 11) modifies the original backing array since it has enough capacity.
// c[0] = 100 changes the shared memory.
