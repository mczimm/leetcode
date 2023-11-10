package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))        //2
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))                 //1
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) //0
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))          //3
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))          //3
	fmt.Println()
	fmt.Println(minSubArrayLen2(7, []int{2, 3, 1, 2, 4, 3}))        //2
	fmt.Println(minSubArrayLen2(4, []int{1, 4, 4}))                 //1
	fmt.Println(minSubArrayLen2(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) //0
	fmt.Println(minSubArrayLen2(11, []int{1, 2, 3, 4, 5}))          //3
	fmt.Println(minSubArrayLen2(11, []int{1, 2, 3, 4, 5}))          //3
	fmt.Println()
	fmt.Println(minSubArrayLen3(7, []int{2, 3, 1, 2, 4, 3}))        //2
	fmt.Println(minSubArrayLen3(4, []int{1, 4, 4}))                 //1
	fmt.Println(minSubArrayLen3(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) //0
	fmt.Println(minSubArrayLen3(11, []int{1, 2, 3, 4, 5}))          //3
	fmt.Println(minSubArrayLen3(11, []int{1, 2, 3, 4, 5}))          //3
}

// My attemption 10.11.23
func minSubArrayLen(target int, nums []int) int {
	var sum, left int
	minLength := len(nums) //watched on leetcode

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			if right-left < minLength {
				minLength = right - left
			}
			sum -= nums[left]
			left++
		}
	}
	if minLength == len(nums) {
		return 0
	}
	return minLength + 1
}

// From the Editorial: https://leetcode.com/problems/minimum-size-subarray-sum/editorial/
func minSubArrayLen2(target int, nums []int) int {
	left, sum := 0, 0
	res := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			res = myMin(res, i+1-left)
			sum -= nums[left]
			left++
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func myMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Combined two above methods
func minSubArrayLen3(target int, nums []int) int {
	left, sum := 0, 0
	minLength := len(nums)
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			if i-left < minLength {
				minLength = i - left
			}
			sum -= nums[left]
			left++
		}
	}
	if minLength == len(nums) {
		return 0
	}
	return minLength + 1
}
