package main

import (
	"fmt"
	"sync"
)

func Email(i int, wg *sync.WaitGroup, ch1 chan bool) {
	defer wg.Done()
	ch1 <- true
	fmt.Println("email ", i)
}
func Sms(i int, wg *sync.WaitGroup, ch2 chan bool) {
	defer wg.Done()
	ch2 <- true
	fmt.Println("sms ", i)
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 6; i++ {
		wg.Add(2)
		ch1 := make(chan bool)
		ch2 := make(chan bool)
		go Email(i, &wg, ch1)
		<-ch1
		go Sms(i, &wg, ch2)
		<-ch2
	}
	wg.Wait()
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

//  Let's break it down for one iteration of the loop (i = 1):

//  go Email(i, &wg, ch1)
// <-ch1         // block until Email sends on ch1

// You start Email in a goroutine.

// But then you immediately do <-ch1, which blocks the main goroutine until Email sends to ch1.

// That means Sms won't start until after Email has sent to ch1 (i.e., has already run or at least started).

// Same goes for Sms:

// go Sms(i, &wg, ch2)
// <-ch2         // block until Sms sends on ch2
