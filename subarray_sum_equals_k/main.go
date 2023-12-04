package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2)) //2
	fmt.Println(subarraySum([]int{1, 2, 3}, 3)) //2
}

func subarraySum(nums []int, k int) int {
	var res, curr int
	var numsM = make(map[int]int)
	numsM[0]++
	for i := 0; i < len(nums); i++ {
		curr += nums[i]
		res += numsM[curr-k]
		numsM[curr]++
	}
	return res
}
