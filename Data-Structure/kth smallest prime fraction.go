package main

import (
	"fmt"
	"sort"
)

func sortFractionValue(p1, p2 []int) int {
	return p1[0]*p2[1] - p1[1]*p2[0]
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	pairArray := [][]int{}

	// Generate all pairs of fractions
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			pairArray = append(pairArray, []int{arr[i], arr[j]})
		}
	}

	// Sort the pairs based on the fraction value
	sort.Slice(pairArray, func(i, j int) bool {
		return sortFractionValue(pairArray[i], pairArray[j]) < 0
	})

	// Return the k-th smallest fraction
	return pairArray[k-1]
}

func main() {
	arr := []int{1, 2, 3, 5}
	k := 3
	result := kthSmallestPrimeFraction(arr, k)
	fmt.Println(result) // Output: [2, 5]
}
