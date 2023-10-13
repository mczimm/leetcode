package main

import "fmt"

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDuplicates(nums []int) []int {
	var res []int
	tmp := make(map[int]bool)

	for _, v := range nums {
		if _, ok := tmp[v]; ok {
			res = append(res, v)
			continue
		}
		tmp[v] = true
	}
	return res
}
