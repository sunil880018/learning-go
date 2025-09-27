package main

import (
	"fmt"
	"sort"
)

// Input: words = ["the","day","is","sunny","the","the","the","sunny","is","is"], k = 4
// Output: ["the","is","sunny","day"]
// Explanation: "the", "is", "sunny" and "day" are the four most frequent words, with the number of occurrence
// being 4, 3, 2 and 1 respectively.

func topKFrequent(words []string, k int) []string {
	mp := make(map[string]int)
	for i := 0; i < len(words); i++ {
		mp[words[i]]++
	}

	type pair struct {
		word  string
		count int
	}

	var arr []pair
	for word, value := range mp {
		arr = append(arr, pair{word, value})
	}
	sort.Slice(arr, func(i int, j int) bool {
		if arr[i].count == arr[j].count {
			return arr[i].word < arr[j].word // Compare words lexicographically
		}
		return arr[i].count > arr[j].count
	})

	var ans []string

	for _, value := range arr {
		if len(ans) == k {
			break
		}
		ans = append(ans, value.word)
	}
	return ans
}

func main() {
	words := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	fmt.Println(topKFrequent(words, 4))
}
