package main

func main() {
	println(xorAllNums2([]int{2, 1, 3}, []int{10, 2, 5, 0})) //13
}

// time limit exceeded
func xorAllNums(nums1 []int, nums2 []int) int {
	var res int

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			res ^= nums1[i] ^ nums2[j]
		}
	}
	return res
}

func xorAllNums2(nums1 []int, nums2 []int) int {
	var res int
	freq := make(map[int]int)
	len1 := len(nums1)
	len2 := len(nums2)

	for _, v := range nums1 {
		freq[v] += len2
	}
	for _, v := range nums2 {
		freq[v] += len1
	}
	for i, v := range freq {
		if v%2 != 0 {
			res ^= i
		}
	}
	return res
}
