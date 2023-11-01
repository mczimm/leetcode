package main

import "fmt"

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))       //[2,2]
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2}))          //[2]
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})) //[4,9]
	fmt.Println(intersect([]int{3, 1, 2}, []int{1, 1}))          //[1]
}

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		tmp := nums1
		nums1 = nums2
		nums2 = tmp
		//intersect(nums2, nums1)
	}

	var res []int

	nums2D := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		nums2D[nums2[i]]++
	}

	for i := 0; i < len(nums1); i++ {
		if _, ok := nums2D[nums1[i]]; ok {
			res = append(res, nums1[i])
			if nums2D[nums1[i]] > 1 {
				nums2D[nums1[i]]--
			} else {
				delete(nums2D, nums1[i])
			}
		}
	}
	return res
}
