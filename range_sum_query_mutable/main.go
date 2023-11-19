package main

import "fmt"

func main() {
	c := Constructor([]int{1, 3, 5})
	fmt.Println(c.SumRange(0, 2))
	c.Update(1, 2)
	fmt.Println(c.SumRange(0, 2))
}

type NumArray struct {
	record []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		record: nums,
	}
}

func (this *NumArray) Update(index int, val int) {
	this.record[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	var res int
	for left <= right {
		res += this.record[left]
		left++
	}
	return res
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
