package main

func main() {
	println(isArraySpecial([]int{2, 1, 4})) //true
	println(isArraySpecial([]int{2, 10}))   //false
}

func isArraySpecial(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]%2 == 0 {
			if nums[i+1]%2 == 0 {
				return false
			}
		}
		if nums[i]%2 != 0 {
			if nums[i+1]%2 != 0 {
				return false
			}
		}
	}
	return true
}
