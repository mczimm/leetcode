package main

import "fmt"

func main() {
	fmt.Println(findLucky([]int{2, 2, 3, 4}))       //2
	fmt.Println(findLucky([]int{1, 2, 2, 3, 3, 3})) //3
	fmt.Println(findLucky([]int{2, 2, 2, 3, 3}))    //-1
}

func findLucky(arr []int) int {
	var arrM = make(map[int]int)
	var res = -1

	for i := 0; i < len(arr); i++ {
		arrM[arr[i]]++
	}

	for i, v := range arrM {
		if i == v {
			if v > res {
				res = v
			}
		}
	}
	return res
}
