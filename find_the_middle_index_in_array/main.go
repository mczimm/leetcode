package main

import "fmt"

func main() {
	fmt.Println(findMiddleIndex([]int{2, 3, -1, 8, 4})) // 3
}

// My first attempt 15.11.23
func findMiddleIndex(nums []int) int {
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

func findMiddleIndex2(nums []int) int {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}

	leftSum := 0
	for k, v := range nums {
		totalSum -= v
		if totalSum == leftSum {
			return k
		}
		leftSum += v
	}

	return -1
}
