// Key Characteristics of defer:

// 1 Execution Order: Deferred functions are executed in LIFO (Last-In-First-Out) order, meaning if you defer
// multiple function calls, they will execute in reverse order of how they were deferred.

// Common Use Cases:

// 2 Closing open resources (files, database connections).
// Unlocking a mutex after locking it.
// Printing final log messages or doing cleanup.

// 3 Execution Timing: Deferred functions are not executed immediately when encountered but will only execute
// after the function where they are defined returns.

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Start")

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	f, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("End")
	// Defer closing the file
	defer f.Close()
}

// output
// Start
// End
// Deferred 2
// Deferred 1
