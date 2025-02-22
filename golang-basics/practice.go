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
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		temp := []int{}
		size := len(queue)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			size = size - 1
		}
		result = append(result, temp)
	}
	return result

}

func main() {
	// Example usage
	// root := &TreeNode{Val: 1}
	// root.Left = &TreeNode{Val: 2}
	// root.Right = &TreeNode{Val: 3}
	// root.Left.Left = &TreeNode{Val: 4}
	// root.Left.Right = &TreeNode{Val: 5}
	// root.Right.Left = &TreeNode{Val: 6}
	// root.Right.Right = &TreeNode{Val: 7}

	// result := levelOrderBottom(root)
	// fmt.Println(result) // Output: [[1] [2 3] [4 5 6 7]]

	slice := []int{2, 13, 22, 24, 8}
	fmt.Println(cap(slice))

}
