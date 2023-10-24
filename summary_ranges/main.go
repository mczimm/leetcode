package main

import (
	"fmt"
)

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))    //["0->2","4->5","7"]
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})) //["0","2->4","6","8->9"]
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var res []string
	left, prev := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]-prev > 1 {
			res = append(res, fmtString(left, prev))
			left, prev = nums[i], nums[i]
		} else {
			prev = nums[i]
		}
	}
	return append(res, fmtString(left, prev))
}

func fmtString(l, r int) string {
	if l == r {
		return fmt.Sprint(l)
	}

	return fmt.Sprint(l, "->", r)
}
