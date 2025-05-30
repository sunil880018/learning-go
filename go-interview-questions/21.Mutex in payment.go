// Why Mutex Is Important in Payment Systems:
// In a payment system, you often deal with shared resources like:

// 1.Wallet balances
// 2.Account transaction histories
// 3.Order statuses
// 4.Inventory counts
// 5. Mutex help in synchronisation

package main

import (
	"fmt"
	"sync"
	"time"
)

var balance = 2000
var mu sync.Mutex

func deduct(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	if balance >= amount {
		// Simulate some delay to see output
		time.Sleep(1000 * time.Millisecond)
		balance -= amount
		fmt.Println("✅ Payment successful, remaining:", balance)
	} else {
		fmt.Println("❌ Insufficient balance. Remaining:", balance)
	}
}

func main() {
	var wg sync.WaitGroup

	// Simulate 5 concurrent deductions
	amounts := []int{400, 400, 400, 600, 500}

	for _, amt := range amounts {
		wg.Add(1)
		go deduct(amt, &wg)
	}

	wg.Wait()

	fmt.Println("Final balance:", balance)
}

// output
// ✅ Payment successful, remaining: 1500
// ✅ Payment successful, remaining: 1100
// ✅ Payment successful, remaining: 700
// ✅ Payment successful, remaining: 300
// ❌ Insufficient balance. Remaining: 300
// Final balance: 300
