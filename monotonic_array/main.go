package main

import "fmt"

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))                     //true
	fmt.Println(isMonotonic([]int{1, 3, 2}))                        //false
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}))                     //true
	fmt.Println(isMonotonic([]int{3, 1, -1, -1, 3, 3, 3, 5, 5, 5})) //false
	fmt.Println(isMonotonic([]int{-1, -1, -1, -1, 3, 4}))           //true
}

// My fisrt attempt
func isMonotonic(nums []int) bool {
	enc, dec := true, true
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dec = false
		}
		if nums[i] < nums[i-1] {
			enc = false
		}
	}

	return enc || dec
}
