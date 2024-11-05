package main

import (
	"fmt"
)

func main() {
	alphabets := "abcdefgh"

	output := []byte{} // Initialize an empty slice of bytes
	vowels := []byte{} // Initialize another empty slice of bytes

	for i := 0; i < len(alphabets); i++ {
		letter := alphabets[i]
		if letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' {
			vowels = append(vowels, letter)
		} else {
			output = append(output, letter)
		}
	}

	output = append(output, vowels...) // Append vowels to the output

	fmt.Println(string(output)) // Output: bcdfghae
}
