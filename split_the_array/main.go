package main

import "fmt"

func main() {
	fmt.Println(isPossibleToSplit([]int{1, 1, 2, 2, 3, 4})) //true
	fmt.Println(isPossibleToSplit([]int{1, 1, 1, 1}))       //false
}

func isPossibleToSplit(nums []int) bool {
	size := len(nums)
	idxArr := make([]int, 101)
	for i := 0; i < size; i++ {
		idxArr[(nums[i])]++
		if idxArr[nums[i]] > 2 {
			return false
		}
	}
	return true
}
