package main

import (
	"fmt"
)

func Add(ch1 chan int, a, b int) {
	ch1 <- a + b
}

func Sub(ch2 chan int, a, b int) {
	ch2 <- a - b
}

func Mul(ch3 chan int, a, b int) {
	ch3 <- a * b
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Run Add and Sub
	go Add(ch1, 4, 7)
	go Sub(ch2, 6, 2)

	a := <-ch1
	b := <-ch2

	// Run Mul
	go Mul(ch3, a, b)

	mul := <-ch3
	fmt.Println("Mul:", mul)
}

// another approach using waitgroup
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func Add(a, b int, res *int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	*res = a + b
// }

// func Sub(a, b int, res *int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	*res = a - b
// }

// func Mul(a, b int, res *int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	*res = a * b
// }

// func main() {
// 	var wg sync.WaitGroup

// 	var addRes, subRes, mulRes int

// 	wg.Add(2)
// 	go Add(4, 7, &addRes, &wg)
// 	go Sub(6, 2, &subRes, &wg)
// 	wg.Wait() // wait for Add and Sub

// 	wg.Add(1)
// 	go Mul(addRes, subRes, &mulRes, &wg)
// 	wg.Wait() // wait for Mul

// 	fmt.Println("Mul:", mulRes)
// }
