/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.
*/

package main

import "fmt"

func main() {
	//fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))  //[24,12,8,6]
	fmt.Println(productExceptSelf2([]int{1, 2, 3, 4})) //[24,12,8,6]
}

func productExceptSelf(nums []int) []int { // time limit exceeded
	var res []int

	for i := 0; i < len(nums); i++ { //1
		r := 1
		for j := 0; j < len(nums); j++ { //1 2
			if i == j {
				continue
			}
			r *= nums[j]
		}
		res = append(res, r)
	}
	return res
}

func productExceptSelf2(nums []int) []int {
	var res []int
	one := 1
	//{1, 2, 3, 4}
	for i := 0; i < len(nums); i++ {
		res = append(res, one) // 1 1 2 6
		one *= nums[i]         // 1 2 6 24
	}

	one = 1

	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= one  // 6 8 12 24
		one *= nums[i] // 4 12 24
	}

	return res
}
