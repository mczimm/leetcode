package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(equalSubstring("abcd", "bcdf", 3))      //3
	fmt.Println(equalSubstring("abcd", "bcdf", 3))      //3
	fmt.Println(equalSubstring("abcd", "cdef", 3))      //1
	fmt.Println(equalSubstring("abcd", "acde", 0))      //1
	fmt.Println(equalSubstring("krrgw", "zjxss", 19))   //2
	fmt.Println(equalSubstring("pxezla", "loewbi", 25)) //4
}

// Firts attempt 13.11.2023
func equalSubstring(s string, t string, maxCost int) int {
	var sum, left int
	res := math.MinInt
	for i := 0; i < len(s); i++ {
		sum += abs(int(s[i]) - int(t[i]))
		//fmt.Printf("%d-%d = %d", s[i], t[i], sum)
		//fmt.Println()
		if sum > maxCost {
			sum -= abs(int(s[left]) - int(t[left]))
			left++
		}
		res = max(res, i-left+1)
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
