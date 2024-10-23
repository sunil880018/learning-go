package main

import (
	"fmt"
)

// findPairWithSum finds a pair of numbers in arr that add up to the target sum.
func findPairWithSum(arr []int, target int) []int {
	// Create a map to store the numbers we have seen so far.
	numMap := make(map[int]int)

	for i, num := range arr {
		complement := target - num

		// Check if the complement exists in the map.
		if _, exists := numMap[complement]; exists {
			return []int{complement, num} // Return the pair.
		}

		// Store the current number in the map.
		numMap[num] = i
	}

	return nil // Return nil if no pair is found.
}

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	result := findPairWithSum(arr, target)

	if result != nil {
		fmt.Printf("Pair found: %d and %d\n", result[0], result[1])
	} else {
		fmt.Println("No pair found")
	}
}
