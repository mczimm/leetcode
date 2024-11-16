package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad")) //bab
	fmt.Println(longestPalindrome("cbbd"))  //bb
	fmt.Println(longestPalindrome("c"))     //c
	fmt.Println(longestPalindrome("ac"))    //a
	fmt.Println(longestPalindrome("acc"))   //cc
}

// brute-force
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	palidrome := func(i, j int) bool {
		left := i
		right := j - 1

		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	for i := len(s); i > 0; i-- {
		for j := 0; j <= len(s)-i; j++ {
			if palidrome(j, i+j) {
				return s[j : i+j]
			}
		}
	}
	return ""
}
