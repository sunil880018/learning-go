package main

import "fmt"

func main() {
	// Create a map with string keys and int values
	myMap := make(map[string]int)

	// Add key-value pairs to the map
	myMap["apple"] = 10
	myMap["banana"] = 20

	// Accessing a value by key
	value := myMap["apple"]

	fmt.Println(value)

	// Check if a key exists in the map
	if value, exists := myMap["apple"]; exists {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key not found")
	}

	// Delete a key from the map
	delete(myMap, "banana")

}
