🌀 What is a Goroutine?
A goroutine is a lightweight, concurrent function execution managed by the Go runtime.

You create a goroutine simply by prefixing a function call with the go keyword:

✅ Key Characteristics of Goroutines

| Feature                        | Description                                                                                                                                          |
| ------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Lightweight**                | Goroutines use a very small initial stack (about 2 KB), and this grows/shrinks as needed. You can easily create thousands or millions of goroutines. |
| **Managed by Go runtime**      | Go schedules and manages goroutines independently of the OS. They are multiplexed onto a smaller number of system threads.                           |
| **Non-blocking**               | Go's runtime includes a **scheduler** that handles blocking operations (like I/O or channel communication) intelligently.                            |
| **Communication via Channels** | Goroutines communicate and synchronize via **channels**, not shared memory.                                                                          |


🧠 How Goroutines Differ from Traditional Threads

| Feature              | Goroutines                               | Traditional Threads (e.g., Java, C++)     |
| -------------------- | ---------------------------------------- | ----------------------------------------- |
| **Creation Cost**    | Very low (\~2 KB stack)                  | High (\~1 MB stack)                       |
| **Scheduling**       | Managed by Go runtime (user-space)       | Managed by OS (kernel-space)              |
| **Scalability**      | Can run millions of goroutines           | Threads are limited (typically thousands) |
| **Communication**    | Use channels (CSP model)                 | Usually use shared memory, mutexes        |
| **Preemption**       | Cooperative + preemptive (since Go 1.14) | Fully preemptive (by OS)                  |
| **Stack Management** | Dynamically growing stack                | Fixed or manually managed                 |

🔄 Goroutines vs Threads Analogy
Think of goroutines as green threads (user-level threads).

The Go runtime has a scheduler that multiplexes goroutines over a pool of system threads.

A goroutine will block only that goroutine, not the entire thread — thanks to Go's I/O and scheduling model.


⚠️ Gotchas
Goroutines do not automatically synchronize — you must use channels or sync primitives (sync.WaitGroup, sync.Mutex).

If your main function exits, all running goroutines are killed — so you must wait for them if needed.

Goroutines are not parallel unless multiple CPUs are used and GOMAXPROCS > 1.

package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	// Simulate work
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait() // Wait for all workers to finish
}


✅ Summary
Goroutines are a lightweight, efficient abstraction for concurrency in Go.

They differ from traditional threads in that they are scheduled by the Go runtime and are extremely scalable.

Channels enable a safe and elegant way for goroutines to communicate without explicit locks.