// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. 
// The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	} else {
		left := 0
		for right := 1; right < len(nums); right++ {
			if nums[left] != nums[right] {
				left++
				nums[left] = nums[right]
			}
			nums[right] = -1
		}
		return left + 1
	}
}
