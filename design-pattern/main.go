package main

import (
	"design-pattern/singleton" // Import the package where the Singleton is defined
	"fmt"
)

func main() {
	instance1 := singleton.GetInstance()
	instance2 := singleton.GetInstance()

	if instance1 == instance2 {
		fmt.Println("Both instances are the same!")
	} else {
		fmt.Println("Different instances exist, which is incorrect for a singleton.")
	}
}
