package main

import "fmt"

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{1}, 1, []int{}, 0)
	merge([]int{0}, 0, []int{1}, 1)
}

// bubble sort
func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums2) == 0 {
		return
	}

	for i := n - 1; i+m >= m; i-- {
		nums1[m+i] = nums2[i]
	}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums1); j++ {
			if nums1[i] < nums1[j] {
				nums1[i], nums1[j] = nums1[j], nums1[i]
			}
		}
	}
	fmt.Println(nums1)
}

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	copy(nums1[m:], nums2)
// 	sort.Ints(nums1)
// }
