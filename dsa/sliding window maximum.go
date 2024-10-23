package main

import (
	"container/heap"
	"fmt"
)

// Define a structure for the max-heap that stores pairs (value, index)
type Pair struct {
	value int
	index int
}

// Define a priority queue (max-heap) that implements heap.Interface
type MaxHeap []Pair

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool {
	// Compare values to maintain the max-heap property
	return h[i].value > h[j].value
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	// Create a max-heap
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	var result []int

	// Initialize the heap with the first window
	for i := 0; i < k; i++ {
		heap.Push(maxHeap, Pair{nums[i], i})
	}
	result = append(result, (*maxHeap)[0].value)

	// Slide the window across the array
	for j := k; j < len(nums); j++ {
		// Add the current element to the max-heap
		heap.Push(maxHeap, Pair{nums[j], j})

		// Remove elements from the heap that are out of the window range
		for (*maxHeap)[0].index <= j-k {
			heap.Pop(maxHeap)
		}

		// Record the maximum of the current window
		result = append(result, (*maxHeap)[0].value)
	}

	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	result := maxSlidingWindow(nums, k)
	fmt.Println("Max sliding window:", result) // Output: [3, 3, 5, 5, 6, 7]
}
