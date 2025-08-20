// Write a program that uses Goroutines to process multiple tasks concurrently.
// Consider you have 3 tasks A, B and C
// A -> Add two numbers
// B-> Sub two numbers
// C-> Mul two numbers
// ex: 10, 8 are input A-> Should return 18 B-> Should return 2 C-> should return 36
// A and B should run concurrently and C should start after A and B completed

package main

import (
	"fmt"
	"sync"
)

func Add(a, b int, wg *sync.WaitGroup, ch1 chan int) {
	defer wg.Done()
	sum := a + b
	fmt.Println("Add : ", sum)
	ch1 <- sum
}
func Sub(a, b int, wg *sync.WaitGroup, ch2 chan int) {
	defer wg.Done()
	sub := a - b
	fmt.Println("Sub : ", sub)
	ch2 <- sub
}
func Mul(a, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	mul := a * b
	fmt.Println("Mul : ", mul)
}
func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	a := make(chan int, 1)
	b := make(chan int, 1)
	go Add(10, 8, &wg, a)
	go Sub(10, 8, &wg, b)
	wg.Wait()
	wg.Add(1)
	x := <-a
	y := <-b
	go Mul(x, y, &wg)
	wg.Wait()
}
