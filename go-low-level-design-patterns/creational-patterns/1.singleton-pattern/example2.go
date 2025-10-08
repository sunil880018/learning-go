package singleton

import (
	"sync"
)

// Singleton is the struct representing the instance.
type Singleton struct{}

// singleInstance holds the single instance of Singleton.
var singleInstance *Singleton

// once ensures the instance is only created once.
var once sync.Once

// GetInstance provides access to the Singleton instance.
func GetInstance() *Singleton {
	// Use sync.Once to create the instance only once.
	once.Do(func() {
		singleInstance = &Singleton{}
	})
	return singleInstance
}

// Benefits of sync.Once
// Thread Safety: The sync.Once mechanism makes this implementation thread-safe without additional locks, ensuring only one instance is created even in concurrent scenarios.
// Lazy Initialization: The Singleton instance is created the first time GetInstance() is called, not at program startup.
