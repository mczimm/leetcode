package searchinsertposition

func searchInsert(nums []int, target int) int {
	//Input: nums = [1,3,5,6], target = 7
	//Output: 4
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}
	return left
}
