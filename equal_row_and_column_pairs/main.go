package main

import (
	"fmt"
)

func main() {
	fmt.Println(equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}})) // 1
	//fmt.Println(equalPairs2([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}})) // 1
	fmt.Println(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}})) // 3
	//fmt.Println(equalPairs2([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}})) // 3
}

// My first attempt 29.11.23
func equalPairs(grid [][]int) int {
	var res, cnt int
	rows := make([][]int, len(grid), len(grid))
	cols := make([][]int, len(grid), len(grid))

	for i := 0; i < len(grid); i++ {
		rows[i] = grid[i]
		for j := 0; j < len(grid); j++ {
			cols[i] = append(cols[i], grid[j][cnt])
		}
		cnt++
	}

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(cols); j++ {
			if equal(rows[i], cols[j]) {
				res++
			}
		}
	}
	return res
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// From the solutions, but need replace struct with int
func equalPairs2(grid [][]int) int {
	gridTMP := make(map[[200]int]struct{})
	var res int
	tmp := [200]int{}

	for i := 0; i < len(grid); i++ {
		copy(tmp[:], grid[i])
		gridTMP[tmp] = struct{}{}
	}

	for i := 0; i < len(grid); i++ {
		t := [200]int{}
		for j := 0; j < len(grid); j++ {
			t[j] = grid[j][i]
			if _, ok := gridTMP[t]; ok {
				res++
			}
		}
	}
	return res
}
