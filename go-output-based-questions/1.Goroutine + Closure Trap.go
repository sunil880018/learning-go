package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
// Answer:

// This prints unpredictable values like 3 3 3 (or fewer), not 0 1 2.


// Explanation:

// The goroutine closes over the loop variable i, and by the time it executes, i is already 3.

// Fix:

for i := 0; i < 3; i++ {
  val := i
  go func() {
      fmt.Println(val)
  }()
}
// output
// 0 1 2