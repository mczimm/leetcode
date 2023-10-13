package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 1} //true
	// nums := []int{1, 2, 3, 4} //false
	nums := []int{3, 3} //true

	res := containsDuplicate(nums)
	fmt.Printf("%v\n", res)
}

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	countMap := make(map[int]int)

	for _, v := range nums {
		countMap[v]++
		if countMap[v] > 1 {
			return true
		}
	}

	return false
}
