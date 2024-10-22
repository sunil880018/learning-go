package main

import "fmt"

// Definition for singly-linked list node.
type ListNode struct {
	val  int
	next *ListNode
}

// Solution struct to hold the hasCycle method.
type Solution struct{}

func (s *Solution) hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for slow != nil && fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true // Cycle detected
		}
	}

	return false // No cycle
}

func main() {
	// Create a linked list with a cycle: 1 -> 2 -> 3 -> 4 -> 2 (cycle)
	head := &ListNode{val: 1}
	head.next = &ListNode{val: 2}
	head.next.next = &ListNode{val: 3}
	head.next.next.next = &ListNode{val: 4}
	head.next.next.next.next = head.next // Create a cycle

	// Instantiate the solution and call the hasCycle method
	solution := &Solution{}
	fmt.Println(solution.hasCycle(head)) // Output: true (the linked list has a cycle)
}
