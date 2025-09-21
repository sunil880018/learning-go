// A concurrent-safe map (no need for manual locks).
// Use when: You have a map shared by many goroutines and want thread-safe access without manually locking.
// ✅ Use case: Shared cache, tracking concurrent sessions, storing temporary data.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// Store values
	m.Store("foo", 1)
	m.Store("bar", 2)

	// Load value
	val, ok := m.Load("foo")
	if ok {
		fmt.Println("foo:", val)
	}

	// Range over map
	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
