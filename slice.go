package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// slice := make([]int, 5)     // Slice with length 5, default initialized to 0
	// slice1 := make([]int, 3, 5) // Slice with length 3 and capacity 5
	// fmt.Println(slice, slice1)  // Output: [0 0 0 0 0]

	// fmt.Println(len(slice1)) // Output: 3 (length)
	// fmt.Println(cap(slice1)) // Output: 5 (capacity)

	// people := [3]Person{
	// 	{Name: "Alice", Age: 30},
	// 	{Name: "Bob", Age: 25},
	// }
	// fmt.Println(people) // Output: [{Alice 30} {Bob 25} { 0}]

	// Create a slice from an array
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4] // Elements from index 1 to 3
	fmt.Println("Slice from array:", slice)

	// Create a slice using make
	s := make([]int, 3, 5)
	fmt.Println("Slice with make:", s)
	fmt.Println("Length:", len(s), "Capacity:", cap(s))

	// Append to a slice
	s = append(s, 10, 20)
	fmt.Println("After appending:", s)

	// Slice an existing slice
	subslice := s[1:3]
	fmt.Println("Sub-slice:", subslice)
}

// A slice is a reference type, which means that it doesn't hold the actual data but rather points to an underlying array.
// The elements of a slice are stored in an array, and a slice itself contains:

// 1.A pointer to the array.
// 2.The length (the number of elements in the slice).
// 3.The capacity (the maximum number of elements the slice can grow to without reallocation).
