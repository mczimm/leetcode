package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))                 //true
	fmt.Println(uniqueOccurrences([]int{1, 2}))                             //false
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0})) //true
	fmt.Println(uniqueOccurrences([]int{-5, -6, 2, 6, -5, -7, 5}))          //false
}

func uniqueOccurrences(arr []int) bool {
	freqMap := make(map[int]int)
	resMap := make(map[int]struct{})
	for i := 0; i < len(arr); i++ {
		freqMap[arr[i]]++
	}
	for _, v := range freqMap {
		if _, ok := resMap[v]; !ok {
			resMap[v] = struct{}{}
		} else {
			return false
		}
	}
	return true
}
