package main

import (
	"fmt"
)

func main() {
	alphabets := "abcdefgh"

	output := []rune{} // Initialize an empty slice of runes
	vowels := []rune{} // Initialize another empty slice of runes

	for i := 0; i < len(alphabets); i++ {
		letter := rune(alphabets[i]) // Convert byte to rune
		if letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' {
			vowels = append(vowels, letter)
		} else {
			output = append(output, letter)
		}
	}

	output = append(output, vowels...) // Append vowels to the output

	fmt.Println(string(output)) // Output: bcdfghae
}
