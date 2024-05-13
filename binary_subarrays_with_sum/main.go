package binary_subarrays_with_sum

//func main() {
//	fmt.Println(numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2)) //4
//	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0)) //15
//}

func sumAtMost(nums []int, goal int) int {
	if len(nums) == 1 && goal == 1 {
		return 1
	}
	left, tot, cur := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		cur += nums[right]

		for left <= right && cur > goal {
			cur -= nums[left]
			left++
		}
		tot += right - left + 1

	}
	return tot
}

func numSubarraysWithSum(nums []int, goal int) int {
	return sumAtMost(nums, goal) - sumAtMost(nums, goal-1)
}
