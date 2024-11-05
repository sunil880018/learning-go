package main

import (
	"fmt"
)

type ListNode struct {
	val  int
	next *ListNode
}

type Solution struct{}

func (s *Solution) detectCycle(head *ListNode) *ListNode {
	visited := make(map[*ListNode]struct{}) // Use a map to track visited nodes
	for head != nil {
		if _, found := visited[head]; found {
			return head // Cycle detected
		}
		visited[head] = struct{}{} // Mark the node as visited
		head = head.next
	}
	return nil // No cycle detected
}

func main() {
	// Example usage
	node1 := &ListNode{val: 3}
	node2 := &ListNode{val: 2}
	node3 := &ListNode{val: 0}
	node4 := &ListNode{val: -4}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node2 // Create a cycle

	solution := &Solution{}
	cycleNode := solution.detectCycle(node1)
	if cycleNode != nil {
		fmt.Println("Cycle detected at node with value:", cycleNode.val) // Output: Cycle detected at node with value: 2
	} else {
		fmt.Println("No cycle detected")
	}
}
