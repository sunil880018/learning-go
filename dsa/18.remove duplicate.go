package main

import (
	"fmt"
)

func removeDuplicates(nums []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}

func main() {
	input := []int{1, 2, 3, 2, 1, 4, 5, 3}
	unique := removeDuplicates(input)
	fmt.Println("Unique elements:", unique)
}
