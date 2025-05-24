package main

import (
	"errors"
	"fmt"
)

// Creating Custom Errors

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("input cannot be negative")
	}
	return x * x, nil
}

func main() {
	r, err := sqrt(-34)
	fmt.Println(r, err)
}
