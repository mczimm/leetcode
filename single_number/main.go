/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.
*/

package main

import "fmt"

func main() {
	fmt.Println(singleNumberDict([]int{1, 0, 1}))       //0
	fmt.Println(singleNumberXOR([]int{2, 2, 1}))        //1
	fmt.Println(singleNumberDict([]int{4, 1, 2, 1, 2})) //4
	fmt.Println(singleNumberXOR([]int{-1, -1, -2}))     //-2
}

// Dict
func singleNumberDict(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	res := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		res[nums[i]]++
	}

	for i, v := range res {
		if v == 1 {
			return i
		}
	}
	return 0
}

// XOR
func singleNumberXOR(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}

	return result
}
