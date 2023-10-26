package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))          //3
	fmt.Println(findLengthOfLCIS([]int{2, 2, 2, 2, 2}))          //1
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 7}))             //4
	fmt.Println(findLengthOfLCIS([]int{1}))                      //1
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 2, 3, 4}))    //3
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 2, 3, 4, 5})) //4
	fmt.Println()
	fmt.Println(findLengthOfLCIS2([]int{1, 3, 5, 4, 7}))          //3
	fmt.Println(findLengthOfLCIS2([]int{2, 2, 2, 2, 2}))          //1
	fmt.Println(findLengthOfLCIS2([]int{1, 3, 5, 7}))             //4
	fmt.Println(findLengthOfLCIS2([]int{1}))                      //1
	fmt.Println(findLengthOfLCIS2([]int{1, 3, 5, 4, 2, 3, 4}))    //3
	fmt.Println(findLengthOfLCIS2([]int{1, 3, 5, 4, 2, 3, 4, 5})) //4
}

// My first attempt
func findLengthOfLCIS(nums []int) int {
	var count int
	var res int

	if len(nums) <= 1 {
		return len(nums)
	}

	nums = append(nums, 0)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			count++
		} else {
			if res < count+1 {
				res = count + 1
			}

			count = 0
		}
	}
	if res < count {
		res = count
	}

	return res
}

// From the solutions
func findLengthOfLCIS2(nums []int) int {
	var ans, anchor = 0, 0

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] >= nums[i] {
			anchor = i
		}
		ans = int(math.Max(float64(ans), float64(i-anchor+1)))
	}
	return ans
}
