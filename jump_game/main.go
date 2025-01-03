package main

//True 1
//False 0

var ans []int

func main() {
	println(canJump([]int{2, 3, 1, 1, 4})) //true
	println(canJump([]int{3, 2, 1, 0, 4})) //false
}

//time limit exceeded
func canJump2(nums []int) bool {
	ans = make([]int, len(nums))

	for i := range nums {
		ans[i] = 0
	}
	ans[len(ans)-1] = 1
	return canJumpFromPosition(0, nums)
}

func canJumpFromPosition(position int, nums []int) bool {
	if position == len(nums)-1 {
		return ans[len(nums)-1] == 1
	}
	furthestJump := position + nums[position]
	if furthestJump > len(nums)-1 {
		furthestJump = len(nums) - 1
	}
	for nextPosition := position + 1; nextPosition <= furthestJump; nextPosition++ {
		if canJumpFromPosition(nextPosition, nums) {
			return true
		}
	}
	return false
}

func canJump(nums []int) bool {
	lastElement := nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastElement {
			lastElement = i
		}
	}
	return lastElement == 0
}
