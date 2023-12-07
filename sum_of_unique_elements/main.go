package main

import "fmt"

func main() {
	fmt.Println(sumOfUnique([]int{1, 2, 3, 2})) // 4
}

func sumOfUnique(nums []int) int {
	var numsM = make(map[int]int)
	var res int

	for i := 0; i < len(nums); i++ {
		numsM[nums[i]]++
	}

	for i, v := range numsM {
		if v <= 1 {
			res += i
		}
	}
	return res
}
