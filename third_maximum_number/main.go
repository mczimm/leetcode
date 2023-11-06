package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(thirdMax([]int{3, 2, 1}))          //1
	fmt.Println(thirdMax([]int{1, 1, 2}))          //2
	fmt.Println(thirdMax([]int{14}))               //14
	fmt.Println(thirdMax([]int{1, 2, 2, 5, 3, 5})) //2
}

func thirdMax(nums []int) int {
	cnt, maxx := 1, 0

	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > maxx {
			maxx = nums[i]
		}

		if nums[i] > nums[i-1] {
			cnt++
		}
		if cnt >= 3 {
			return nums[i-1]
		}
	}
	return maxx
}
