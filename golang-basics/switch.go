package main

import "fmt"

func main() {
	x := 3
	switch x {
	case 1:
		fmt.Println("1....")
	case 2:
		fmt.Println("2...")
	default:
		fmt.Println("nothing")
	}
}
