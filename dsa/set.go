package main

import "fmt"

func main() {
	// Creating a set of unique elements using struct{} as values
	set := map[string]struct{}{}

	// Add elements to the set
	set["apple"] = struct{}{}
	set["banana"] = struct{}{}
	set["cherry"] = struct{}{}

	// Check if an element exists in the set
	if _, exists := set["apple"]; exists {
		fmt.Println("apple is in the set")
	}

	// Iterating over the set
	for key := range set {
		fmt.Println(key)
		fmt.Println(set[key])
	}
}
