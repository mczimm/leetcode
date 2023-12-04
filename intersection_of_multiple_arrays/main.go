package main

import (
	"fmt"
	"sort"
)

func main() {
	// [3,4]
	fmt.Println(intersection([][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}}))
}

func intersection(nums [][]int) []int {
	var res []int
	numsM := make(map[int]int)
	numsL := len(nums)

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			numsM[nums[i][j]]++
		}
	}

	for i, v := range numsM {
		if v == numsL {
			res = append(res, i)
		}
	}

	sort.Ints(res)
	return res
}
