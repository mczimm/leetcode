package main

import "fmt"

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)) //6
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 1, 0, 0}, 0))             //3
	fmt.Println(longestOnes([]int{0, 0, 0, 0}, 0))                      //0
	fmt.Println(longestOnes([]int{0}, 1))                               //1
	fmt.Println(longestOnes([]int{1, 0, 1, 0}, 0))                      //1
}

func longestOnes(nums []int, k int) int {
	zeroCount, longestWindow, start := 0, 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
		}
		for zeroCount > k {
			if nums[start] == 0 {
				zeroCount--
			}
			start++
		}
		longestWindow = max(longestWindow, i-start+1)
	}
	return longestWindow
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
