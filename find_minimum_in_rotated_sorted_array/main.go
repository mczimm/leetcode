package main

func main() {
	println(findMin([]int{3, 4, 5, 1, 2})) // 1
	println(findMin([]int{1}))             // 1
}

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	if len(nums) == 1 {
		return nums[0]
	}

	if nums[right] > nums[0] {
		return nums[0]
	}

	for right >= left {
		mid := left + (right-left)/2

		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		if nums[mid] > nums[0] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
