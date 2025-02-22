// sort.Slice provides a simple way to sort slices using a custom comparison function.
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"John", 25},
		{"Alice", 30},
		{"Bob", 20},
	}

	// Custom sort by Age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("Sorted by Age:", people)

	// Custom sort by Name (alphabetically)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})

	fmt.Println("Sorted by Name:", people)
}
