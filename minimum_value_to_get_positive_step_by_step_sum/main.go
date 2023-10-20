package main

import "fmt"

func main() {
	fmt.Println(minStartValue([]int{-3, 2, -3, 4, 2})) //5
	fmt.Println(minStartValue([]int{1, 2}))            //1
	fmt.Println(minStartValue([]int{1, -2, -3}))       //5
}

// My first attempt
func minStartValue(nums []int) int {
	startValue := 1

	for startValue+nums[0] <= 0 {
		startValue++
	}

	for {
		if check(nums, startValue) {
			return startValue
		} else {
			startValue++
		}
	}
}

func check(nums []int, currSum int) bool {
	for i := 0; i < len(nums); i++ {
		if currSum+nums[i] > 0 {
			currSum = currSum + nums[i]
			continue
		} else {
			return false
		}
	}
	return true
}

// From solutions:https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/solutions/2537047/go-o-n-solution/
func minStartValue2(nums []int) int {
	prefixSum, result := int(0), int(1)

	for _, num := range nums {
		prefixSum += num

		result = max(result, 1-prefixSum)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
