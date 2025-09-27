package main

import (
	"fmt"
)

func main() {
	const connectionEdges = 4
	connectionFrom := []int{1, 2, 3, 5}
	connectionTo := []int{2, 3, 4, 6}
	graph := make(map[int][]int)

	for i := 0; i < connectionEdges; i++ {
		from := connectionFrom[i]
		to := connectionTo[i]

		// Initialize the slice if the key does not exist
		if _, exists := graph[from]; !exists {
			graph[from] = []int{}
		}
		if _, exists := graph[to]; !exists {
			graph[to] = []int{}
		}

		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	// Print the graph
	for key, value := range graph {
		fmt.Printf("%d: %v\n", key, value)
	}
}
