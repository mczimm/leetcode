package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))       //true
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))       //true
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2)) //false
	fmt.Println(containsNearbyDuplicate([]int{4, 1, 2, 3, 1, 5}, 3)) //true
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if k < 0 {
		return false
	}
	if len(nums) < 1 {
		return false
	}

	var d = make(map[int]int)

	for i, v := range nums {
		if _, ok := d[v]; ok {
			if int(math.Abs(float64(i-d[v]))) <= k {
				return true
			}
		}
		d[v] = i
	}

	return false
}
