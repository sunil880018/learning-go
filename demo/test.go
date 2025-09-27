package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5} // len=5, cap=5
	fmt.Println("Length of x", len(x))
	fmt.Println("Capacity of x", cap(x))

	x = append(x, 6) //  len=6, cap=10
	fmt.Println("Length of x", len(x))
	fmt.Println("Capacity of x", cap(x))
	x = append(x, 7) //  len=7, cap=10
	fmt.Println("Length of x", len(x))
	fmt.Println("Capacity of x", cap(x))
	a := x[4:]
	fmt.Println(a) //a={5,6,7}, len=3, cap=6 because --> cap(x) - start_index = 10 - 4 = 6
	fmt.Println("Length of a", len(a))
	fmt.Println("Capacity of a", cap(a))

	y := alterSlice(a)

	fmt.Println(x) // x= {1,2,3,4,10,6,7} len=7, cap=10
	fmt.Println("Length of x", len(x))
	fmt.Println("Capacity of x", cap(x))
	fmt.Println(y)
	fmt.Println("Length of y", len(y)) // y= {10,6,7,11},len=4, cap=6
	fmt.Println("Capacity of y", cap(y))

}
func alterSlice(a []int) []int {
	a[0] = 10         // 10 modify in a also modify in x because both share the same backing array
	a = append(a, 11) // 11 append in a not x
	return a
}

// Key Learnings

// 1.Slices share the same backing array until a reallocation occurs.

// 2.Sub-slices keep capacity from the starting index.

// 3.append may or may not reallocate:

// 	If within capacity → modifies the same array.

// 	If exceeding capacity → allocates new memory.

// 4.Mutating one slice may affect others referencing the same array.
