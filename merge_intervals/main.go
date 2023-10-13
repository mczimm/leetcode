package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))                    //[[1,5]]
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) //[[1,6],[8,10],[15,18]]
	fmt.Println(merge([][]int{{1, 3}}))
	fmt.Println(merge([][]int{{1, 4}, {5, 6}}))
	fmt.Println(merge([][]int{{1, 4}, {0, 4}}))
	fmt.Println(merge([][]int{{1, 4}, {2, 3}}))
	fmt.Println(merge([][]int{{1, 4}, {0, 2}, {3, 5}}))
}

func merge(intervals [][]int) [][]int {

	if len(intervals) < 1 {
		return [][]int{}
	}

	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{{intervals[0][0], intervals[0][1]}}

	for _, v := range intervals {
		last_end := result[len(result)-1][1]

		if v[0] <= last_end {
			result[len(result)-1][1] = max(last_end, v[1])
		} else {
			result = append(result, []int{v[0], v[1]})
		}
	}

	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
