package main

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)  // [5,6,7,1,2,3,4]
	rotate2([]int{1, 2, 3, 4, 5, 6, 7}, 3) // [5,6,7,1,2,3,4]
}

// extra space
func rotate(nums []int, k int) {
	tmp := make([]int, len(nums))
	n := len(nums)

	for i := 0; i < n; i++ {
		tmp[(k+i)%n] = nums[i]
	}

	for i := 0; i < n; i++ {
		nums[i] = tmp[i]
	}
}

// in place
func rotate2(nums []int, k int) {
	n := len(nums)
	reverse(&nums, 0, n-1)
	reverse(&nums, 0, k%n-1)
	reverse(&nums, k%n, n-1)
}

func reverse(nums *[]int, l, r int) {
	for l <= r {
		(*nums)[l], (*nums)[r] = (*nums)[r], (*nums)[l]
		l++
		r--
	}
}
