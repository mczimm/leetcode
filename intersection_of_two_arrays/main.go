package main

import "fmt"

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))       //2
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})) //4,9
}

// My first attempt
func intersection(nums1 []int, nums2 []int) []int {
	keepUnique := make(map[int]struct{})
	var res []int

	if len(nums1) < len(nums2) {
		intersection(nums2, nums1)
	}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				keepUnique[nums1[i]] = struct{}{}
			}
		}
	}
	for k := range keepUnique {
		res = append(res, k)
	}
	return res
}
