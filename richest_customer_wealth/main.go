package main

import "fmt"

func main() {
	fmt.Println(maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))            //6
	fmt.Println(maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))          //10
	fmt.Println(maximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}})) //17
}

func maximumWealth(accounts [][]int) int {
	var cur, res int

	for i := 0; i < len(accounts); i++ {
		for j := 0; j < len(accounts[i]); j++ {
			cur += accounts[i][j]
		}
		if cur > res {
			res = cur
		}
		cur = 0
	}

	return res
}
