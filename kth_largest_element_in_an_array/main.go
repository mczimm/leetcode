package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findKthLargest2([]int{3, 2, 1, 5, 6, 4}, 2))          //5
	fmt.Println(findKthLargest2([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)) //4
	fmt.Println(findKthLargest2([]int{99, 99}, 1))                    //99
	fmt.Println(findKthLargest2([]int{-1, -1}, 2))                    //-1
	fmt.Println(findKthLargest2([]int{-1, 2, 0}, 2))                  //0
}

// works only with positive numbers
func findKthLargest(nums []int, k int) int {
	pos := make([][]int, 100000)
	var res int
	for _, num := range nums {
		pos[num] = append(pos[num], num)
	}
	for i := len(pos) - 1; i > 0; i-- {
		if k <= 0 {
			break
		}
		if pos[i] != nil {
			k = k - len(pos[i])
			if k <= 0 {
				res = pos[i][0]
			}
		}
	}
	return res
}

func findKthLargest2(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
