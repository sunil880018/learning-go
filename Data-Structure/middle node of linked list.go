package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	first := head
	second := head
	// Move 'second' pointer twice as fast as 'first'
	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next.Next
	}
	return first
}

func main() {
	// Example usage: Creating a linked list 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	// Finding the middle node
	mid := middleNode(head)
	fmt.Println("Middle node value:", mid.Val) // Output: 3
}
