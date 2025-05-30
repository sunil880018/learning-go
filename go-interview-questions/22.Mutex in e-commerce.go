package main

import (
	"fmt"
	"sync"
)

type Inventory struct {
	stock map[string]int
	mu    sync.RWMutex
}

func NewInventory() *Inventory {
	return &Inventory{
		stock: make(map[string]int),
	}
}

// Add stock to an item
func (inv *Inventory) AddStock(item string, quantity int) {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	inv.stock[item] += quantity
	fmt.Printf("📦 Added %d units of %s. Total stock: %d\n", quantity, item, inv.stock[item])
}

// Purchase an item
func (inv *Inventory) Buy(item string) {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	if inv.stock[item] > 0 {
		inv.stock[item]--
		fmt.Printf("✅ %s purchased. Remaining stock: %d\n", item, inv.stock[item])
	} else {
		fmt.Printf("❌ %s out of stock.\n", item)
	}
}

// Check stock level
func (inv *Inventory) GetStock(item string) int {
	inv.mu.RLock()
	defer inv.mu.RUnlock()
	return inv.stock[item]
}

func main() {
	inv := NewInventory()

	inv.AddStock("iPhone", 10)
	inv.AddStock("MacBook", 5)

	var wg sync.WaitGroup

	items := []string{"iPhone", "MacBook", "iPhone", "iPhone", "MacBook", "MacBook"}

	for _, item := range items {
		wg.Add(1)
		go func(item string) {
			defer wg.Done()
			inv.Buy(item)
		}(item)
	}

	wg.Wait()
	fmt.Println("Remaining stock for iPhone:", inv.GetStock("iPhone"))
	fmt.Println("Remaining stock for MacBook:", inv.GetStock("MacBook"))
}

// output
// 📦 Added 10 units of iPhone. Total stock: 10
// 📦 Added 5 units of MacBook. Total stock: 5
// ✅ MacBook purchased. Remaining stock: 4
// ✅ iPhone purchased. Remaining stock: 9
// ✅ MacBook purchased. Remaining stock: 3
// ✅ iPhone purchased. Remaining stock: 8
// ✅ iPhone purchased. Remaining stock: 7
// ✅ MacBook purchased. Remaining stock: 2
// Remaining stock for iPhone: 7
// Remaining stock for MacBook: 2
