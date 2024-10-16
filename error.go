package main

import (
	"errors"
	"fmt"
)

// Creating Custom Errors
var ErrNegativeInput = errors.New("input cannot be negative")

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeInput
	}
	return x * x, nil
}

func main() {
	r, err := sqrt(34)
	fmt.Println(r, err)
}
