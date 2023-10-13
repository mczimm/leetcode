package main

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))  //[0,1,9,16,100]
	fmt.Println(sortedSquares2([]int{-4, -1, 0, 3, 10})) //[0,1,9,16,100]
}

func sortedSquares(nums []int) []int {
	var res []int
	left := 0
	right := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		sl := nums[left] * nums[left]
		sr := nums[right] * nums[right]
		if sr > sl {
			res = append([]int{sr}, res...)
			right--
		} else {
			res = append([]int{sl}, res...)
			left++
		}
	}

	return res
}

func sortedSquares2(nums []int) []int {
	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}
