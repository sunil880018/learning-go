package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root} // Initialize the queue with the root node
	checkReverse := 0          // Flag for zigzag traversal

	for len(queue) > 0 {
		queueSize := len(queue)
		temp := make([]int, 0, queueSize) // Slice to hold current level values

		for i := 0; i < queueSize; i++ {
			node := queue[0]              // Get the front node
			queue = queue[1:]             // Dequeue
			temp = append(temp, node.Val) // Add the node value to the current level slice

			// Enqueue left and right children if they exist
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Reverse the order of the temp slice if checkReverse is odd
		if checkReverse%2 == 1 {
			reverse(temp)
		}
		result = append(result, temp) // Append the current level to the result
		checkReverse++                // Increment the zigzag flag
	}

	return result
}

// Helper function to reverse a slice of integers
func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	result := zigzagLevelOrder(root)
	fmt.Println(result) // Output: [[1], [3, 2], [4, 5, 6, 7]]
}
