package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))                   //2
	fmt.Println(missingNumber([]int{0, 1}))                      //2
	fmt.Println(missingNumber([]int{1, 0}))                      //2
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) //8
}

// From the solutions:https://leetcode.com/problems/missing-number/solutions/4149107/go-o-1-space/
func missingNumber(nums []int) int {
	res := len(nums)

	for i := 0; i < len(nums); i++ {
		res += i - nums[i]
	}
	return res
}
