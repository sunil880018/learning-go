package main

import (
	"fmt"
	"sync"
)

func Email(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("email", i)
}

func Sms(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("sms", i)
}

func main() {
	for i := 1; i <= 6; i++ {
		var wg sync.WaitGroup
		wg.Add(1)
		go Email(i, &wg)
		wg.Wait() // wait for Email to finish before Sms

		wg.Add(1)
		go Sms(i, &wg)
		wg.Wait() // wait for Sms to finish before next iteration
	}
}

// output
// email 1
// sms 1
// email 2
// sms 2
// email 3
// sms 3
// email 4
// sms 4
// email 5
// sms 5
// email 6
// sms 6
