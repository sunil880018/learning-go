// Question: Create a worker pool that processes a stream of tasks concurrently, with a fixed number of workers.
// Each task can be represented as a function that takes an integer and returns an integer.

package main

import (
	"fmt"
	"sync"
)

// Task function type
type Task func(int) int

// WorkerPool structure
type WorkerPool struct {
	workerCount int
	tasks       chan Task
	wg          sync.WaitGroup
}

// NewWorkerPool creates a new WorkerPool
func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		workerCount: workerCount,
		tasks:       make(chan Task),
	}
}

// Start begins the worker pool
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerCount; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}
}

// worker processes tasks from the task channel
func (wp *WorkerPool) worker() {
	defer wp.wg.Done()
	for task := range wp.tasks {
		result := task(0) // In a real scenario, you'd pass an argument
		fmt.Println("Processed task result:", result)
	}
}

// Submit adds a new task to the worker pool
func (wp *WorkerPool) Submit(task Task) {
	wp.tasks <- task
}

// Wait waits for all workers to finish processing
func (wp *WorkerPool) Wait() {
	close(wp.tasks) // Closing the channel signals workers to stop
	wp.wg.Wait()
}

func main() {
	wp := NewWorkerPool(3)
	wp.Start()

	// Submit some tasks
	for i := 0; i < 10; i++ {
		n := i // capture the loop variable
		wp.Submit(func(int) int {
			return n * n // example task: square the number
		})
	}

	wp.Wait() // Wait for all tasks to complete
}

// Explanation:
// WorkerPool Struct: Contains the number of workers and a channel for tasks.
// Start Method: Initializes and starts a fixed number of workers that wait for tasks.
// Worker Method: Each worker continuously processes tasks from the channel.
// Submit Method: Sends a task to the workers through the channel.
// Wait Method: Closes the task channel and waits for all workers to finish processing.
