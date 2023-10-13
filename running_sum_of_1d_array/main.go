package main

import "fmt"

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))     //[1,3,6,10]
	fmt.Println(runningSum([]int{1, 1, 1, 1, 1}))  //[1,2,3,4,5]
	fmt.Println(runningSum([]int{3, 1, 2, 10, 1})) //[3,4,6,16,17]
}

func runningSum(nums []int) []int {
	var res []int
	var cur int

	if len(nums) < 1 {
		return []int{}
	}
	if len(nums) == 1 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		cur += nums[i]
		res = append(res, cur)
	}

	return res
}
