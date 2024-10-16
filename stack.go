package main

import (
	"fmt"
)

// Stack structure
type Stack struct {
	items []int
}

// Push adds an item to the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes the last item from the stack and returns it
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	lastIndex := len(s.items) - 1
	poppedItem := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return poppedItem
}

func main() {
	stack := Stack{}

	// Push items to the stack
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	// Pop items from the stack
	fmt.Println(stack.Pop()) // Output: 30
	fmt.Println(stack.Pop()) // Output: 20
	fmt.Println(stack.Pop()) // Output: 10
}
