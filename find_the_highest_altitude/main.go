package main

import "fmt"

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7})) // 1
}

func largestAltitude(gain []int) int {
	res := 0
	tmp := 0

	for i := 0; i < len(gain); i++ {
		tmp += gain[i]
		if tmp > res {
			res = tmp
		}
	}
	return res
}
