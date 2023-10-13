package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMaxAverage([]int{0}, 1)) //12.75000
}

func findMaxAverage(nums []int, k int) float64 {
	res := float64(-1 << 63)
	cnt, cur := float64(0), float64(0)

	for i := 0; i < len(nums); i++ {
		cur += float64(nums[i])
		cnt++

		if cnt == float64(k) {
			res = math.Max(res, cur/float64(k))
		}
		if cnt > float64(k) {
			cur -= float64(nums[i-k])
			res = math.Max(res, cur/float64(k))
		}
	}

	return res
}
