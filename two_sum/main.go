package main

import "fmt"

func main() {
	nums := []int{2, 5, 5, 11}
	target := int(13)
	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSum2(nums, target))
}

func twoSum2(nums []int, target int) []int {
	mmap := make(map[int]int)
	res := make([]int, 2, 2)

	//for i := 0; i < len(nums); i++ {
	//	mmap[nums[i]] = i
	//}

	for i := 0; i < len(nums); i++ {
		tt := target - nums[i]
		if _, ok := mmap[tt]; ok && mmap[tt] != i {
			res[0] = i
			res[1] = mmap[tt]
			break
		}
		mmap[nums[i]] = i
	}
	return res
}

func twoSum(nums []int, target int) []int {
	var res []int
	cnt := 1
	for i := 0; i < len(nums); i++ {
		if len(res) == 2 {
			break
		}
		tmp := target - nums[i]
		for j := cnt; j < len(nums); j++ {
			if tmp == nums[j] {
				res = append(res, i, j)
				break
			}
		}
		cnt++
	}
	return res
}
