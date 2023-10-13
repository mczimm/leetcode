/*
Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

There is only one repeated number in nums, return this repeated number.

You must solve the problem without modifying the array nums and uses only constant extra space.

Follow up:

How can we prove that at least one duplicate number must exist in nums?
Can you solve the problem in linear runtime complexity?
*/

package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))  //2
	fmt.Println(findDuplicate2([]int{3, 1, 3, 4, 2})) //3
}

func findDuplicate(nums []int) int {
	res := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := res[v]; ok {
			return v
		}
		res[v] = struct{}{}
	}
	return 0
}

func findDuplicate2(nums []int) int {
	slow := nums[0]
	fast := nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
