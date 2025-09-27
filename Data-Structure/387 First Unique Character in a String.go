package main

import "fmt"

func firstUniqChar(s string) int {
	// Create a map to store the frequency of each character.
	mp := make(map[rune]int)

	// Loop through the string and count each character's occurrences.
	for _, char := range s {
		mp[char]++
	}

	// Loop through the string again and return the index of the first unique character.
	for i, char := range s {
		if mp[char] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(firstUniqChar("loveleetcode")) // Output: 2 (first unique character is 'v')
}
