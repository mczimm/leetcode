package length_of_longest_subarray_with_at_most_k_frequency

//package main

//func main() {
//	fmt.Println(maxSubarrayLength([]int{1, 2, 3, 1, 2, 3, 1, 2}, 2)) //6
//	fmt.Println(maxSubarrayLength([]int{1, 2, 1, 2, 1, 2, 1, 2}, 1)) //2
//	fmt.Println(maxSubarrayLength([]int{5, 5, 5, 5, 5, 5, 5}, 4))    //4
//	fmt.Println(maxSubarrayLength([]int{1, 1, 1, 3}, 2))             //3
//}

func maxSubarrayLength(nums []int, k int) int {
	res, start := 0, -1
	numsM := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		numsM[nums[i]]++
		for numsM[nums[i]] > k {
			start++
			numsM[nums[start]]--
		}
		res = max(res, i-start)
	}

	return res
}
