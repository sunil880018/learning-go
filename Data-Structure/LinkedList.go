package main

import "fmt"

type LinkedList struct {
	Head *Node
}
type Node struct {
	Val  int
	Next *Node
}

func (ll *LinkedList) display() {
	head := ll.Head
	for head != nil {
		fmt.Printf(" -> %d", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func (ll *LinkedList) insertNode(Val int) {
	newNode := &Node{Val: Val}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	head := ll.Head
	for head.Next != nil {
		head = head.Next
	}
	head.Next = newNode

}

func (ll *LinkedList) deleteNode() {
	if ll.Head == nil {
		return
	}
	if ll.Head.Next == nil {
		ll.Head = nil
		return
	}
	head := ll.Head
	for head.Next.Next != nil {
		head = head.Next
	}
	head.Next = nil
}

func main() {
	ll := &LinkedList{}
	ll.insertNode(10)
	ll.insertNode(30)
	ll.insertNode(20)
	ll.insertNode(50)
	ll.insertNode(60)
	ll.insertNode(40)
	ll.display()
	ll.deleteNode()
	ll.display()
}
