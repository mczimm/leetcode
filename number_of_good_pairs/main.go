package number_of_good_pairs

func numIdenticalPairs(nums []int) int {
	resM := make(map[int]int)
	res := 0
	for i := 0; i < len(nums); i++ {
		resM[nums[i]]++
	}

	for _, v := range resM {
		if v >= 2 {
			// this hint from leetcode
			res += v * (v - 1) / 2
		}
	}
	return res
}
