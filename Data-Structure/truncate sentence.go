package main

import (
	"fmt"
	"strings"
)

func truncateSentence(s string, k int) string {
	return strings.Join(strings.Split(s, " ")[:k], " ")
}

func main() {
	fmt.Println(truncateSentence("Hello world this is Go", 4)) // Output: "Hello world this is"
}
