package main

import (
	"fmt"
)

func main() {
	fmt.Println(firstUniqChar("leetcode"))     //0
	fmt.Println(firstUniqChar("loveleetcode")) //2
	fmt.Println(firstUniqChar("aabb"))         //-1
}

// My first attempt
func firstUniqChar(s string) int {
	sD := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		sD[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if sD[s[i]] == 1 {
			return i
		}
	}
	return -1
}
