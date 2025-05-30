// What is escape analysis in Go and how does it affect performance?
// Answer: Escape analysis determines whether a variable should be allocated on the stack or heap.

// Stack is faster.

// Heap has GC overhead.

func alloc() *int {
	a := 5
	return &a // escapes to heap
}
