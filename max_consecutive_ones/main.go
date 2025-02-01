package main

import "fmt"

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})) //3
}

func findMaxConsecutiveOnes(nums []int) int {
	curr, res := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			curr++
		} else {
			if curr > res {
				res = curr
			}
			curr = nums[i]
		}
	}
	if curr > res {
		res = curr
	}
	return res
}

func findMaxConsecutiveOnes2(nums []int) int {
	var ones, res int

	for _, v := range nums {
		if v == 1 {
			ones++
		} else {
			res = max(res, ones)
			ones = 0
		}
	}
	return max(res, ones)
}
