package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestSubarray([]int{0, 0, 1, 1}))                //2
	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))                //3
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1})) //5
	fmt.Println(longestSubarray([]int{1, 1, 1}))                   //2

	fmt.Println()
	fmt.Println(longestSubarray2([]int{0, 0, 1, 1}))                //2
	fmt.Println(longestSubarray2([]int{1, 1, 0, 1}))                //3
	fmt.Println(longestSubarray2([]int{0, 1, 1, 1, 0, 1, 1, 0, 1})) //5
	fmt.Println(longestSubarray2([]int{1, 1, 1}))                   //2

	fmt.Println()
	fmt.Println(longestSubarray3([]int{0, 0, 1, 1}))                //2
	fmt.Println(longestSubarray3([]int{1, 1, 0, 1}))                //3
	fmt.Println(longestSubarray3([]int{0, 1, 1, 1, 0, 1, 1, 0, 1})) //5
	fmt.Println(longestSubarray3([]int{1, 1, 1}))                   //2
}

// My attempt
func longestSubarray(nums []int) int {
	cnt, res, skip := 0, 0, 0

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if skip == 1 && nums[j] == 0 {
				if cnt > res {
					res = cnt
				}
				break
			}
			if nums[j] == 0 {
				skip = 1
			} else {
				cnt++
			}
		}
		if cnt == len(nums)-1 {
			return cnt
		}
		if cnt == len(nums) {
			return cnt - 1
		}
		if cnt > res {
			res = cnt
		}
		skip = 0
		cnt = 0
	}
	return res
}

// From the solutions: https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/solutions/3515286/concise-go-solution/
func longestSubarray2(nums []int) int {
	l, r := 0, 0
	zeroes := 0
	for ; r < len(nums); r++ {
		zeroes += (1 - nums[r])
		if zeroes > 1 {
			zeroes -= (1 - nums[l])
			l++
		}
	}
	return r - l - 1
}

// Editorial: https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/editorial/
func longestSubarray3(nums []int) int {
	zeroCount, longestWindow, start := 0, 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
		}
		for zeroCount > 1 {
			if nums[start] == 0 {
				zeroCount--
			}
			start++
		}
		longestWindow = max(longestWindow, i-start)
	}
	return longestWindow
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
