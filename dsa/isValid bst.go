package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// Solution struct to encapsulate the methods.
type Solution struct{}

// Inorder traversal to fill the slice with node values.
func (s *Solution) inorder(root *TreeNode, nodes *[]int) {
	if root == nil {
		return
	}
	s.inorder(root.left, nodes)
	*nodes = append(*nodes, root.val)
	s.inorder(root.right, nodes)
}

// Method to validate if the binary tree is a valid BST.
func (s *Solution) isValidBST(root *TreeNode) bool {
	nodes := []int{}
	s.inorder(root, &nodes)

	for i := 0; i < len(nodes)-1; i++ {
		if nodes[i] >= nodes[i+1] {
			return false
		}
	}
	return true
}

// Example usage
func main() {
	// Constructing a binary tree
	root := &TreeNode{val: 2}
	root.left = &TreeNode{val: 1}
	root.right = &TreeNode{val: 3}

	solution := &Solution{}
	fmt.Println(solution.isValidBST(root)) // Output: true
}
