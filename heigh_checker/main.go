package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3})) // 3
}

// My first attempt 16.11.23
func heightChecker(heights []int) int {
	var res int
	var heightsSorted = make([]int, len(heights))
	copy(heightsSorted, heights)
	sort.Ints(heightsSorted)

	for i := 0; i < len(heights); i++ {
		if heights[i] != heightsSorted[i] {
			res++
		}
	}

	return res
}
