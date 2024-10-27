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