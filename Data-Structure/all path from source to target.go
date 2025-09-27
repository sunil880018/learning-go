package main

import (
	"fmt"
)

type Solution struct{}

// dfs performs a depth-first search to find all paths from the current node to the target node.
func (s *Solution) dfs(graph [][]int, node int, path []int, results *[][]int) {
	path = append(path, node) // Add the current node to the path
	if node == len(graph)-1 { // Check if we reached the target node
		// Append a copy of the current path to the results
		*results = append(*results, append([]int{}, path...))
	} else {
		// Explore all neighbors of the current node
		for _, next := range graph[node] {
			s.dfs(graph, next, path, results)
		}
	}
	path = path[:len(path)-1] // Backtrack: remove the current node from the path
}

// allPathsSourceTarget initializes the DFS and returns all paths from source to target.
func (s *Solution) allPathsSourceTarget(graph [][]int) [][]int {
	results := [][]int{}
	path := []int{}
	s.dfs(graph, 0, path, &results)
	return results
}

// Test case
func main() {
	solution := &Solution{}
	graph := [][]int{{1, 2}, {3}, {3}, {}} // Example graph
	result := solution.allPathsSourceTarget(graph)
	fmt.Println(result) // Output: [[0, 1, 3], [0, 2, 3]]
}
