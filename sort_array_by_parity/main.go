package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4})) //[2,4,3,1]
	fmt.Println(sortArrayByParity([]int{0}))          //[0]
}

func sortArrayByParity(nums []int) []int {
	var res []int

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			res = append([]int{nums[i]}, res...)
		} else {
			res = append(res, nums[i])
		}
	}

	return res
}
