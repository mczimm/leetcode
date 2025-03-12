package main

import "sort"

func main() {
	println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1})) //1
	println(lastStoneWeight([]int{2, 2}))             //1
}

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		x := stones[len(stones)-1]
		y := stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		if x >= y {
			z := x - y
			if z > 0 {
				stones = append(stones, z)
			}
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}
