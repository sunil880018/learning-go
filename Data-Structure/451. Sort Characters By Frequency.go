package main

import (
	"fmt"
	"sort"
)

func frequencySort(s string) string {
	// Step 1: Create a map to count frequencies using rune
	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}

	// Step 2: Convert the map to a slice of key-value pairs
	type pair struct {
		char  rune
		count int
	}
	var arr []pair
	for char, count := range freq {
		arr = append(arr, pair{char, count})
	}

	// Step 3: Sort the slice based on frequency in descending order
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].count > arr[j].count
	})

	// Step 4: Build the result string without using strings.Builder
	var result string
	for _, p := range arr {
		for i := 0; i < p.count; i++ {
			result += string(p.char)
		}
	}

	return result
}

func main() {
	fmt.Println(frequencySort("tree"))   // Output: "eert" or "eetr"
	fmt.Println(frequencySort("cccaaa")) // Output: "cccaaa" or "aaaccc"
	fmt.Println(frequencySort("Aabb"))   // Output: "bbAa" or "bbaA"
}
