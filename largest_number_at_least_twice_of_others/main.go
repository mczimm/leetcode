package main

func main() {
	println(dominantIndex([]int{3, 6, 1, 0})) //1
	println(dominantIndex([]int{1, 2, 3, 4})) //-1
}

func dominantIndex(nums []int) int {
	var idx int
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[idx] {
			idx = i
		}
	}

	for i := 0; i < len(nums); i++ {
		if idx != i && nums[idx] < nums[i]*2 {
			return -1
		}
	}
	return idx
}
