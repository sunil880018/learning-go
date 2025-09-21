// sync.RWMutex
//
//	A read-write lock:
//	Multiple readers can lock simultaneously (RLock).
//	Only one writer can lock (Lock).
//
// Use when: You have lots of reads but few writes.
// ✅ Use case: Cache reads, configuration data, shared map mostly read by goroutines.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var rw sync.RWMutex
	data := 0

	// Writer
	go func() {
		rw.Lock()
		data = 42
		rw.Unlock()
	}()

	// Reader
	rw.RLock()
	fmt.Println("Data:", data)
	rw.RUnlock()
}
