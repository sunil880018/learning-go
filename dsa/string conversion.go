package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Integer to string
	num := 42
	strNum := strconv.Itoa(num)
	fmt.Println("Integer to string:", strNum)

	// String to integer
	str := "123"
	intNum, err := strconv.Atoi(str)
	if err == nil {
		fmt.Println("String to integer:", intNum)
	}

	// Float to string
	f := 3.14159
	strFloat := strconv.FormatFloat(f, 'f', 2, 64)
	fmt.Println("Float to string:", strFloat)

	// String to float
	floatStr := "2.718"
	floatNum, err := strconv.ParseFloat(floatStr, 64)
	if err == nil {
		fmt.Println("String to float:", floatNum)
	}
}
