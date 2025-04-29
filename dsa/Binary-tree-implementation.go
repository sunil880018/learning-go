package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// func InsertNode(val int) *Node {

// }
// func DeleteNode() *Node {

// }
func LevelOrderTraversal(root *Node) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*Node{root}

	for len(queue) > 0 {
		size := len(queue)
		temp := []int{}
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			temp = append(temp, node.Val)
			size--
		}
		result = append(result, temp)
	}
	return result
}

func LevelOrderTraversalRecursion(root *Node, level int, result *[][]int) {
	if root == nil {
		return
	}

	if len(*result) <= level {
		*result = append(*result, []int{})
	}

	(*result)[level] = append((*result)[level], root.Val)

	LevelOrderTraversalRecursion(root.Left, level+1, result)
	LevelOrderTraversalRecursion(root.Right, level+1, result)
}
func PreOrderTraversal(root *Node, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	PreOrderTraversal(root.Left, result)
	PreOrderTraversal(root.Right, result)
}

func InorderTraversal(root *Node, result *[]int) {
	if root == nil {
		return
	}
	InorderTraversal(root.Left, result)
	*result = append(*result, root.Val)
	InorderTraversal(root.Right, result)
}
func PostOrderTraversal(root *Node, result *[]int) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.Left, result)
	PostOrderTraversal(root.Right, result)
	*result = append(*result, root.Val)
}

func main() {
	root := &Node{Val: 24}
	root.Left = &Node{Val: 34}
	root.Right = &Node{Val: 45}
	preOrderResult := []int{}
	inorderResult := []int{}
	postOrderResult := []int{}
	levelOrderRecursionResult := [][]int{}
	PreOrderTraversal(root, &preOrderResult)
	InorderTraversal(root, &inorderResult)
	PostOrderTraversal(root, &postOrderResult)
	LevelOrderTraversalRecursion(root, 0, &levelOrderRecursionResult)
	fmt.Println(preOrderResult, inorderResult, postOrderResult, levelOrderRecursionResult)
}
