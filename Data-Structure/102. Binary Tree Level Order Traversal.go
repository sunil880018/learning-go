package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		var temp []int
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:] // dequeue
			temp = append(temp, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// result = append([][]int{temp}, result...) // Prepend current level to the result
		result = append(result, temp) // Append the current level to the result
	}

	return result
}

func main() {
	// Example usage
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	result := levelOrderBottom(root)
	fmt.Println(result) // Output: [[4 5 6 7] [2 3] [1]]
}
