package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) //49
}

func maxArea(height []int) int {
	mArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		width := right - left
		mArea = max(mArea, min(height[left], height[right])*width)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return mArea
}
