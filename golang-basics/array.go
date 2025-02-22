package main

import "fmt"

func main() {
	var arr [3]int   // Declare an array of 3 integers
	fmt.Println(arr) // Output: [0 0 0] (default zero value for int)

	var arr1 [4]bool
	fmt.Println(arr1) // default [false false false false]
}
