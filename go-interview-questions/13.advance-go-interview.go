// 1. How does Go achieve concurrency, and how is it different from parallelism?
// Answer:

// Concurrency (assign Single CPU) is about dealing with multiple tasks at once, while parallelism (assign multiple CPUs) is about running multiple tasks simultaneously.

// Go achieves concurrency using goroutines (lightweight threads) and channels (communication mechanism).

// Goroutines are multiplexed onto OS threads by the Go runtime scheduler, enabling efficient concurrency even with thousands of goroutines.

// 2. What happens under the hood when you use go func() to start a goroutine?
// Answer:

// The Go runtime allocates a small stack for the new goroutine (~2KB, dynamically grows).

// A new entry is added to the goroutine scheduler's queue.

// The Go scheduler maps goroutines to OS threads using a work-stealing and M:N scheduling model (M goroutines on N OS threads).

// Context switches are handled in user space, which makes them cheaper than OS threads.

// 3. What are data races in Go? How do you detect and avoid them?
// Answer:

// A data race occurs when:

// Two or more goroutines access the same variable concurrently.

// At least one of them writes to the variable.

// There’s no synchronization.

// Use -race flag in go run or go test to detect data races.

// Avoid them by using channels, mutexes, or sync/atomic operations.

// 4. What is the memory model of Go?
// Answer:

// Go’s memory model defines when reads and writes in different goroutines are guaranteed to be visible to each other.

// Synchronization via channels, mutexes, or atomic operations provides memory visibility guarantees.

// Without proper synchronization, there is no guarantee that changes made by one goroutine will be visible to another.

// 5. How does garbage collection work in Go?
// Answer:

// Go uses a concurrent, tri-color mark-and-sweep garbage collector.

// It runs concurrently with the application (non-stop-the-world).

// Memory is categorized into black (live), gray (reachable but unscanned), and white (garbage).

// GC aims for low pause times and is optimized for latency.

// 6. Explain the select statement with a real-world scenario.
// Answer:

// select lets a goroutine wait on multiple channel operations.

// Real-world use: timeout pattern.

// select {
// case res := <-resultChan:
//     fmt.Println(res)
// case <-time.After(2 * time.Second):
//     fmt.Println("timeout")
// }
// 7. How do you handle context cancellation in Go?
// Answer:

// Use the context package (context.WithCancel, context.WithTimeout, etc.) to cancel goroutines.

// Example:
// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// defer cancel()

// select {
// case <-ctx.Done():
//     fmt.Println("Context cancelled:", ctx.Err())
// }
// 8. What’s the difference between buffered and unbuffered channels in terms of synchronization?
// Answer:

// Unbuffered channels block sender and receiver until both are ready (synchronous communication).

// Buffered channels allow sending without a receiver up to the buffer limit (asynchronous until full).

// Buffered channels introduce complexity — risk of goroutine leaks if not drained properly.

// 9. How does Go handle dependency management (Go modules)?
// Answer:

// With Go Modules (go.mod, go.sum), Go handles versions, reproducible builds, and vendoring.

// go mod tidy, go mod download, go mod vendor are key commands.

// Modules support versioning via Semantic Import Versioning (v2+ versions are in path).

// 10. What are the downsides of goroutines and how do you mitigate them?
// Answer:

// Downsides:

// Goroutine leaks (e.g., blocked on channel forever).

// Uncontrolled spawning → memory pressure.

// Mitigation:

// Use context to cancel.

// Limit goroutines using worker pools.

// Monitor using tools like pprof.