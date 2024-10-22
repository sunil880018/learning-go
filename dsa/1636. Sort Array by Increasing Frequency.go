package main

import (
	"fmt"
	"sort"
)

// Input: nums = [1,1,2,2,2,3]
// Output: [3,1,1,2,2,2]
// Explanation: '3' has a frequency of 1, '1' has a frequency of 2, and '2' has a frequency of 3.

func frequencySort(nums []int) []int {
	ans := []int{}

	mp := make(map[int]int)
	for _, value := range nums {
		mp[value]++
	}

	type pair struct {
		value int
		count int
	}

	arr := []pair{}
	for value, count := range mp {
		arr = append(arr, pair{value, count})
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i].count == arr[j].count {
			return arr[i].value > arr[j].value
		}
		return arr[i].count < arr[j].count
	})
	fmt.Println(arr)

	for _, value := range arr {
		for i := 0; i < value.count; i++ {
			ans = append(ans, value.value)
		}
	}

	return ans
}

func main() {
	arr := []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}
	fmt.Println(frequencySort(arr))
}
