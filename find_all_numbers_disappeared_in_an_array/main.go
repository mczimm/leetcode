package main

import "fmt"

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})) //[5,6]
	fmt.Println(findDisappearedNumbers([]int{1, 1}))                   //[2]
}

// My first attempt
func findDisappearedNumbers(nums []int) []int {
	d := make(map[int]struct{})
	var res []int
	nums = append([]int{0}, nums...)

	for i := 1; i < len(nums); i++ {
		d[nums[i]] = struct{}{}
	}

	for i := 1; i < len(nums); i++ {
		if _, ok := d[i]; !ok {
			res = append(res, i)
		}
	}
	return res
}
