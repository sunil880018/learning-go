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
	for i := 0; i < len(t); i++ {
		if ms[rune(t[i])] != mt[rune(t[i])] {
			return false
		}
	}
	return true
}