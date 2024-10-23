package main

import (
	"fmt"
	"sync"
	"time"
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

	// Start timing
	startTime := time.Now()

	// Submit some tasks
	for i := 0; i < 10; i++ {
		n := i // capture the loop variable
		wp.Submit(func(int) int {
			time.Sleep(100 * time.Millisecond) // simulate work with sleep
			return n * n                       // example task: square the number
		})
	}

	wp.Wait() // Wait for all tasks to complete

	// Stop timing
	elapsedTime := time.Since(startTime)
	fmt.Printf("Total time taken: %s\n", elapsedTime)
}
