package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(smallestRangeI([]int{0, 10}, 2)) // 6
}

// From the discussion: https://leetcode.com/problems/smallest-range-i/
func smallestRangeI(nums []int, k int) int {
	sort.Ints(nums)
	low := nums[0] + k
	high := nums[len(nums)-1] - k
	res := high - low
	if res < 0 {
		return 0
	}
	return res
}
