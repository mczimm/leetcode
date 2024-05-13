package main

import "fmt"

func main() {
	fmt.Println(maximumUniqueSubarray([]int{4, 2, 4, 5, 6}))             //17
	fmt.Println(maximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5})) //8
}

func maximumUniqueSubarray(nums []int) int {
	numsUniq := make(map[int]int)
	left, cur, res := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		numsUniq[nums[right]]++
		cur += nums[right]
		for numsUniq[nums[right]] > 1 {
			numsUniq[nums[left]]--
			cur -= nums[left]
			left++
		}
		res = max(res, cur)
	}
	return res
}
