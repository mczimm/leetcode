package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))                //[-1,3,-1]
	fmt.Println(nextGreaterElement([]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7})) //[7,7,7,7,7]
}

// brute force
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))

	for i := 0; i < len(nums1); i++ {
		ans[i] = -1
		found := false
		for j := 0; j < len(nums2); j++ {
			if found && nums2[j] > nums1[i] {
				ans[i] = nums2[j]
				break
			}
			if nums1[i] == nums2[j] {
				found = true
			}
		}
	}
	return ans
}
