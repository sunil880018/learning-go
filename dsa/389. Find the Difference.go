package main

func findTheDifference(s string, t string) byte {
	result := byte(0)

	for i := 0; i < len(s); i++ {
		result ^= s[i]
	}
	for i := 0; i < len(t); i++ {
		result ^= t[i]
	}

	return result
}

func main() {
	s := "abcd"
	t := "abcde"
	println(findTheDifference(s, t))
}
