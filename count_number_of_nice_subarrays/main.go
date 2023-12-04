package main

import "fmt"

func main() {
	fmt.Println(numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3)) // 2
}

func numberOfSubarrays(nums []int, k int) int {
	var res, curr int
	var numsM = make(map[int]int)
	numsM[0]++
	for i := 0; i < len(nums); i++ {
		curr += nums[i] % 2
		//curr += nums[i]
		res += numsM[curr-k]
		numsM[curr]++
	}
	return res
}
