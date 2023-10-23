package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))           //1
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{3}))              //1
	fmt.Println(findContentChildren([]int{10, 9, 8, 7}, []int{5, 6, 7, 8})) //2
}

// My first attempt
func findContentChildren(g []int, s []int) int {
	//g - child
	//s - cookies

	if len(s) == 0 {
		return 0
	}

	var res int
	var cookie int
	sort.Ints(g)
	sort.Ints(s)

	for child := 0; child < len(s); child++ {
		if s[child] >= g[cookie] {
			res++
			cookie++
		}
		if cookie == len(g) {
			break
		}
	}
	return res
}
