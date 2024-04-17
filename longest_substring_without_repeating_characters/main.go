package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
}

func lengthOfLongestSubstring(s string) int {
	result, left, right := 0, 0, 0
	tmp := make(map[byte]int)
	for right < len(s) {
		tmp[s[right]]++

		for tmp[s[right]] > 1 {
			tmp[s[left]]--
			left++
		}

		result = max(result, right-left+1)
		right++
	}
	return result
}
