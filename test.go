package main

import (
	"fmt"
	"sync"
)

var mt sync.Mutex

func main() {
	couter := 0
	mt.Lock()
	for i := 0; i < 1000; i++ {
		couter++
	}
	mt.Unlock()
	fmt.Println("Counter : ", couter)
}
