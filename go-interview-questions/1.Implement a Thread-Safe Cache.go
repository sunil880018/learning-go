// Question: Design a thread-safe in-memory cache that supports Get and Set operations.
// The cache should have a maximum size and use an eviction policy (like Least Recently Used - LRU) when
// it exceeds the maximum size.

package main

import (
	"container/list"
	"fmt"
	"sync"
)

// Cache structure
type Cache struct {
	mutex        sync.Mutex
	capacity     int
	data         map[string]*list.Element
	evictionList *list.List
}

// CacheEntry represents a single entry in the cache
type CacheEntry struct {
	key   string
	value string
}

// NewCache creates a new Cache instance
func NewCache(capacity int) *Cache {
	return &Cache{
		capacity:     capacity,
		data:         make(map[string]*list.Element),
		evictionList: list.New(),
	}
}

// Get retrieves the value from the cache
func (c *Cache) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, found := c.data[key]; found {
		c.evictionList.MoveToFront(element)
		return element.Value.(*CacheEntry).value, true
	}
	return "", false
}

// Set adds a new key-value pair to the cache
func (c *Cache) Set(key string, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, found := c.data[key]; found {
		element.Value.(*CacheEntry).value = value
		c.evictionList.MoveToFront(element)
		return
	}

	if c.evictionList.Len() >= c.capacity {
		c.evict()
	}

	entry := &CacheEntry{key, value}
	element := c.evictionList.PushFront(entry)
	c.data[key] = element
}

// evict removes the least recently used item
func (c *Cache) evict() {
	if element := c.evictionList.Back(); element != nil {
		c.evictionList.Remove(element)
		cacheEntry := element.Value.(*CacheEntry)
		delete(c.data, cacheEntry.key)
	}
}

func main() {
	cache := NewCache(3)

	cache.Set("1", "one")
	cache.Set("2", "two")
	cache.Set("3", "three")

	fmt.Println(cache.Get("1")) // Output: one true

	cache.Set("4", "four") // Evicts key "2"

	fmt.Println(cache.Get("2")) // Output:  false (not found)
	fmt.Println(cache.Get("3")) // Output: three true
	fmt.Println(cache.Get("4")) // Output: four true
}

// Explanation:
// Cache Struct: Contains a map for fast access, a linked list for maintaining the order of usage, and a mutex for thread safety.
// Get Method: Retrieves a value, moving the accessed entry to the front of the linked list to mark it as recently used.
// Set Method: Adds or updates an entry. If the cache exceeds capacity, it calls the evict method.
// Eviction Logic: Removes the least recently used item from both the linked list and the map.
