package main

import "fmt"

func main() {
	var p *int
	var x int
	x = 14
	p = &x
	fmt.Println(x, *p)
}
