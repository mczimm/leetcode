package main

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})  //[1,3,12,0,0]
	moveZeroes2([]int{0, 1, 0, 3, 12}) //[1,3,12,0,0]
}

// My first attempt
func moveZeroes(nums []int) {
	l := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[l] = nums[i]
			l++
		}
	}
	for l < len(nums) {
		nums[l] = 0
		l++
	}
}

func moveZeroes2(nums []int) {
	l := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] != 0 {
			nums[r], nums[l] = nums[l], nums[r]
			l++
		}
	}
}
