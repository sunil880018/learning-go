package main

import (
	"fmt"
	"math"
)

func lengthOfLongestSubstring(s string) int {
	freq := make(map[byte]int) // Frequency map to store the last index of each character
	maximumLength := 0
	left := 0

	for i := 0; i < len(s); i++ {
		// If the character is already in the map and its left index is within the current window
		if lastIndex, exists := freq[s[i]]; exists && left <= lastIndex {
			left = lastIndex + 1
		}
		freq[s[i]] = i // Update the last seen index of the character
		maximumLength = int(math.Max(float64(maximumLength), float64(i-left+1)))
	}

	return maximumLength
}

func main() {
	s := "abcabcbb"
	result := lengthOfLongestSubstring(s)
	fmt.Println("Length of Longest Substring Without Repeating Characters:", result) // Output: 3
}
