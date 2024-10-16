package main

import "fmt"

func main() {
	// Traditional for loop
	for i := 0; i < 3; i++ {
		fmt.Println("Traditional loop:", i)
	}

	// While-like for loop
	i := 0
	for i < 3 {
		fmt.Println("While-like loop:", i)
		i++
	}

	// Infinite loop
	j := 0
	for {
		fmt.Println("Infinite loop:", j)
		j++
		if j == 3 {
			break
		}
	}

	// Range-based loop over slice
	arr := []string{"apple", "banana", "cherry"}
	for index, value := range arr {
		fmt.Printf("Index %d: %s\n", index, value)
	}
}
