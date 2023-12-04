package main

import (
	"fmt"
	"sort"
)

func main() {
	// [[1,2,10],[4,5,7,8]]
	fmt.Println(findWinners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}))
	//[[1,2,5,6],[]]
	fmt.Println(findWinners([][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}))
}

func findWinners(matches [][]int) [][]int {
	var res [][]int
	winM := make(map[int]int)
	loseM := make(map[int]int)
	var win, lose []int

	for i := 0; i < len(matches); i++ {
		winM[matches[i][0]]++
		loseM[matches[i][1]]++
	}

	for i := range winM {
		if _, ok := loseM[i]; !ok {
			win = append(win, i)
		}
	}

	for i := range loseM {
		if loseM[i] == 1 {
			lose = append(lose, i)
		}
	}

	sort.Ints(win)
	sort.Ints(lose)
	res = append(res, win, lose)

	return res
}
