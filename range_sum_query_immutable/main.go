package main

import "fmt"

func main() {
	c := Constructor([]int{-2, 0, 3, -5, 2, -1}) //[0,2],[2,5],[0,5]
	fmt.Println(c.SumRange(0, 2))
	fmt.Println(c.SumRange(2, 5))
	fmt.Println(c.SumRange(0, 5))
	fmt.Println(c)
}

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	var sum int

	if left < 0 || left > right || right > len(this.nums) {
		return 0
	}

	for left <= right {
		sum += this.nums[left]
		left++
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
