package main

import "fmt"

func main() {
	fmt.Println(fairCandySwap([]int{1, 1}, []int{2, 2})) //1,2
}

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	var a, b int
	bB := make(map[int]int)

	for i := 0; i < len(aliceSizes); i++ {
		a += aliceSizes[i]
	}
	for i := 0; i < len(bobSizes); i++ {
		b += bobSizes[i]
	}
	delta := (b - a) / 2

	for i := 0; i < len(bobSizes); i++ {
		bB[bobSizes[i]] = i
	}
	for i := 0; i < len(aliceSizes); i++ {
		if _, ok := bB[aliceSizes[i]+delta]; ok {
			return []int{aliceSizes[i], aliceSizes[i] + delta}
		}
	}
	return []int{}
}
