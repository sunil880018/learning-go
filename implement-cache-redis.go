package main

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	Name  string
	Email string
}

// wrap data with expiry
type cacheItem struct {
	data   Data
	expiry time.Time
}

type Memory struct {
	store map[string]cacheItem
	mu    sync.RWMutex
}

type Cache interface {
	Get(id string) (Data, bool)
	Set(id string, data Data, ttl time.Duration)
}

func (m *Memory) Get(id string) (Data, bool) {
	m.mu.RLock()
	item, ok := m.store[id]
	m.mu.RUnlock()

	if !ok {
		return Data{}, false
	}

	// check expiry
	if time.Now().After(item.expiry) {
		m.mu.Lock()
		delete(m.store, id)
		m.mu.Unlock()
		return Data{}, false
	}

	return item.data, true
}

func (m *Memory) Set(id string, data Data, ttl time.Duration) {
	m.mu.Lock()
	m.store[id] = cacheItem{
		data:   data,
		expiry: time.Now().Add(ttl),
	}
	m.mu.Unlock()
}

func CacheHolder(cache Cache) {
	data := Data{Name: "abc", Email: "abc@example.com"}
	cache.Set("123", data, 3*time.Second) // TTL = 3 sec

	if val, ok := cache.Get("123"); ok {
		fmt.Println("Got from cache:", val)
	} else {
		fmt.Println("Not found in cache")
	}

	time.Sleep(4 * time.Second) // wait past TTL

	if val, ok := cache.Get("123"); ok {
		fmt.Println("Got from cache:", val)
	} else {
		fmt.Println("Expired / Not found in cache")
	}
}

func main() {
	memCache := &Memory{store: make(map[string]cacheItem)}
	CacheHolder(memCache)
}
