package main

import "fmt"

func majorityElement(nums []int) int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	var ans int
	for i, val := range mp {
		fmt.Println(i, val)
		if (len(nums) / 2) < val {
			return i
		}
	}
	return ans
}
func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}
