package main

import "fmt"

func main() {
	//3,4
	fmt.Println(findIntersectionValues([]int{4, 3, 2, 3, 1}, []int{2, 2, 5, 2, 3, 6}))
	//0,0
	fmt.Println(findIntersectionValues([]int{3, 4, 2, 3}, []int{1, 5}))
}

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	var res []int
	nums1M := make(map[int]int)
	nums2M := make(map[int]int)
	var n1, n2 int

	for i := 0; i < len(nums1); i++ {
		nums1M[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		nums2M[nums2[i]]++
	}

	for i := 0; i < len(nums1); i++ {
		if _, ok := nums2M[nums1[i]]; ok {
			n1++
		}
	}
	for i := 0; i < len(nums2); i++ {
		if _, ok := nums1M[nums2[i]]; ok {
			n2++
		}
	}
	res = append(res, n1, n2)

	return res
}
