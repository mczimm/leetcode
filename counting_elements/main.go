package main

import "fmt"

func main() {
	fmt.Println(countElements([]int{1, 2, 3}))          // 2
	fmt.Println(countElements([]int{1, 3, 2, 3, 5, 0})) // 3
	fmt.Println(countElements([]int{1, 1, 2, 2}))       // 2
}

func countElements(arr []int) int {
	arrMap := make(map[int]bool)
	var res int

	for i := 0; i < len(arr); i++ {
		arrMap[arr[i]] = true
	}

	for _, v := range arr {
		if ok, _ := arrMap[v+1]; ok {
			res++
		}
	}
	return res
}
