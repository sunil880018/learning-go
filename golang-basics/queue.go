package main

import (
	"fmt"
)

// Queue structure
type Queue struct {
	items []int
}

// Enqueue adds an item to the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes the first item from the queue and returns it
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	dequeuedItem := q.items[0]
	q.items = q.items[1:]
	return dequeuedItem
}

func main() {
	queue := Queue{}

	// Enqueue items to the queue
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	// Dequeue items from the queue
	fmt.Println(queue.Dequeue()) // Output: 10
	fmt.Println(queue.Dequeue()) // Output: 20
	fmt.Println(queue.Dequeue()) // Output: 30
}
