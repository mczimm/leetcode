package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1} //true
	//nums := []int{1, 2, 3, 4} //false
	//nums := []int{3, 3} //true

	res := containsDuplicate2(nums)
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

// 24.04.2024
func containsDuplicate2(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	checkMap := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		if _, ok := checkMap[nums[i]]; ok {
			return true
		}
		checkMap[nums[i]] = struct{}{}
	}
	return false
}
