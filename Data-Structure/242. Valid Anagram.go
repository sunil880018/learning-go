package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ms := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		ms[rune(s[i])]++
	}
	mt := make(map[rune]int)

	for i := 0; i < len(t); i++ {
		mt[rune(t[i])]++
	}

	fmt.Println(mt, ms)
	for i := 0; i < len(t); i++ {
		if ms[rune(t[i])] != mt[rune(t[i])] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("listen", "silent")) // true
	fmt.Println(isAnagram("hello", "world"))   // false
}
