package main

import (
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

type Solution struct{}

func (s *Solution) search(root *TreeNode, n *TreeNode, temp []*TreeNode) bool {
	if root == nil {
		return false
	}

	temp = append(temp, root) // Add current node to the path
	if root.val == n.val {
		return true
	}

	// Search in the left and right subtrees
	if s.search(root.left, n, temp) || s.search(root.right, n, temp) {
		return true
	}

	// Backtrack if not found in this subtree
	temp = temp[:len(temp)-1]
	return false
}

func (s *Solution) lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	vp := []*TreeNode{}
	vq := []*TreeNode{}

	s.search(root, p, vp)
	s.search(root, q, vq)

	var ans *TreeNode
	i := 0
	for i < len(vp) && i < len(vq) {
		if vp[i].val == vq[i].val {
			ans = vp[i]
		}
		i++
	}

	return ans
}

func main() {
	// Example usage:
	tree := &TreeNode{val: 3}
	tree.left = &TreeNode{val: 5}
	tree.right = &TreeNode{val: 1}
	tree.left.left = &TreeNode{val: 6}
	tree.left.right = &TreeNode{val: 2}
	tree.left.right.left = &TreeNode{val: 7}
	tree.left.right.right = &TreeNode{val: 4}
	tree.right.left = &TreeNode{val: 0}
	tree.right.right = &TreeNode{val: 8}

	solution := &Solution{}
	p := tree.left             // Node with value 5
	q := tree.left.right.right // Node with value 4
	lca := solution.lowestCommonAncestor(tree, p, q)
	if lca != nil {
		fmt.Println("Lowest Common Ancestor:", lca.val) // Output: 5
	} else {
		fmt.Println("Lowest Common Ancestor not found")
	}
}
