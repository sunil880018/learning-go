package main

import (
	"fmt"
	"sort"
)

// Definition for a binary tree node.
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// Function to perform vertical traversal of the binary tree.
func traversalStart(root *TreeNode, distance int, ans map[int][]int) {
	if root == nil {
		return
	}

	// Add the value to the corresponding distance in the map
	if _, exists := ans[distance]; !exists {
		ans[distance] = []int{}
	}
	ans[distance] = append(ans[distance], root.val)

	// Traverse left and right children with updated distance
	traversalStart(root.left, distance-1, ans)
	traversalStart(root.right, distance+1, ans)
}

// Function to get the vertical order traversal of the binary tree.
func verticalTraversal(root *TreeNode) [][]int {
	ans := make(map[int][]int)
	traversalStart(root, 0, ans)

	// Sort the keys of the ans map
	keys := make([]int, 0, len(ans))
	for key := range ans {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	finalAns := [][]int{}
	for _, key := range keys {
		// Sort the values in ascending order
		sort.Ints(ans[key])
		finalAns = append(finalAns, ans[key])
	}

	return finalAns
}

func main() {
	// Example binary tree creation
	root := &TreeNode{val: 1}
	root.left = &TreeNode{val: 2}
	root.right = &TreeNode{val: 3}
	root.left.left = &TreeNode{val: 4}
	root.left.right = &TreeNode{val: 5}
	root.right.right = &TreeNode{val: 6}

	result := verticalTraversal(root)
	fmt.Println(result) // Output: [[4],[2],[1],[5],[3],[6]]
}
