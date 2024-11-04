package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	flip := true
	wg.Add(2)

	n := 1

	go func() {
		fmt.Println("Starting Odd Go Routine")
		for {
			if flip && n%2 == 1 {
				fmt.Println("Odd :", n)
				n++
				flip = false
			}
			if n == 10 {
				break
			}
		}
		wg.Done()
	}()

	go func() {
		fmt.Println("Starting even Go Routine")
		for {
			if flip == false && n%2 == 0 {
				fmt.Println("Even :", n)
				n++
				flip = true
			}
			if n == 11 {
				break
			}
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("shutdown")
}
