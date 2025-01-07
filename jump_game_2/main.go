package main

func main() {
	println(jump([]int{2, 3, 1, 1, 4})) //2
}

func jump(nums []int) int {
	//lastElement := nums[len(nums)-1]
	var cur1, cur2, ans int
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] >= cur1 {
			cur1 = nums[i] + i
		}
		if i == cur2 {
			ans++
			cur2 = cur1
		}
	}
	return ans
}
