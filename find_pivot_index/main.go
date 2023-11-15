package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6})) // 3
}

func pivotIndex(nums []int) int {
	var left, total int

	for _, v := range nums {
		total += v
	}

	for i := 0; i < len(nums); i++ {
		if left == total-nums[i] {
			return i
		} else {
			left += nums[i]
			total -= nums[i]
		}
	}

	return -1
}
