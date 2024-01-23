package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170})) //["Mary","Emma","John"]
}

func sortPeople(names []string, heights []int) []string {
	var res []string
	var tmpHM = make(map[int]string)
	for i := 0; i < len(heights); i++ {
		tmpHM[heights[i]] = names[i]
	}

	sort.Ints(heights)
	for i := len(heights) - 1; i >= 0; i-- {
		res = append(res, tmpHM[heights[i]])
	}
	return res
}
