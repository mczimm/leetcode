package main

import "fmt"

func main() {
	fmt.Println(findMissingRanges([]int{0, 1, 3, 50, 75}, 0, 99)) //[[2,2],[4,49],[51,74],[76,99]]
	fmt.Println(findMissingRanges([]int{-1}, -1, -1))             //[]
	fmt.Println(findMissingRanges([]int{}, 1, 1))                 //[[1,1]]
}

func findMissingRanges(nums []int, lower int, upper int) [][]int {
	var res [][]int
	l := len(nums)

	if l == 0 {
		res = append(res, []int{lower, upper})
		return res
	}

	if lower < nums[0] {
		res = append(res, []int{lower, nums[0] - 1})
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] <= 1 {
			continue
		}
		res = append(res, []int{nums[i] + 1, nums[i+1] - 1})
	}
	if upper > nums[l-1] {
		res = append(res, []int{nums[l-1] + 1, upper})
	}
	return res
}
