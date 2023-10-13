package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3})) //[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
	fmt.Println(subsets([]int{0}))       //[[],[0]]
}

func subsets(nums []int) [][]int {
	var res [][]int
	res = append(res, []int{})

	for _, val := range nums {
		for _, sub := range res {
			subset := append(append([]int{}, sub...), val)
			res = append(res, subset)
		}
	}

	return res
}
