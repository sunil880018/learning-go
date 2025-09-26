package main

import "fmt"

func main() {
	// input - {0, 7, 0, 3, 12}
	// output -{0,0,7,3,12}
	arr := []int{0, 7, 0, 3, 12}
	mp := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		mp[arr[i]]++
	}
	fmt.Println(mp)
	output := make([]int, len(arr))

	k := 0
	for i := 0; i < len(arr); i++ {
		val := mp[arr[i]]
		for j := 1; j <= val; j++ {
			output[k] = arr[i]
			k = k + 1
		}
		mp[arr[i]] = 0
	}
	fmt.Println(output)

}
