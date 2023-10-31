package main

import "fmt"

func main() {
	fmt.Println(getAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3)) //[-1,-1,-1,5,4,4,-1,-1,-1]
}

// My first attempt
func getAverages(nums []int, k int) []int {
	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		res[i] = -1
	}

	var left, right, sum = 0, 0, 0
	slideWindow := 2*k + 1

	for right < len(nums) {
		sum += nums[right]
		if right-left+1 >= slideWindow {
			res[right-k] = sum / slideWindow
			sum -= nums[left]
			left++
		}
		right++
	}

	return res
}
