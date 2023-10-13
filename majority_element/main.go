/*
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

Follow-up: Could you solve the problem in linear time O(n) and in O(1) space?
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))             //3
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2})) //2
}

func majorityElement2(nums []int) int { // O(nlogN)
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElement(nums []int) int { // O(nlogN)
	var res int
	vote := 0

	for _, v := range nums {
		if vote == 0 {
			res = v
		}
		if v == res {
			vote++
		} else {
			vote--
		}
	}
	return res
}
